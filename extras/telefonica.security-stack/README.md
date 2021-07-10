# Access a broker with Telefónica security stack

This documentation explains how to access an Orion Context Broker that the endpoints of NGSI API
are protected by Telefónica security stack (Keystone).

## Prepair

Clone the NGSI Go repository and move to `ngsi-go/extras/telefonica.security-stack` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/telefonica.security-stack
```

## Start up

Start up an Orion context broker and Keystone.

```
docker-compose up -d
```

You can see seven containers.

```
docker-compose ps
```

```
                 Name                                Command               State                                     Ports
-------------------------------------------------------------------------------------------------------------------------------------------------------------
telefonicasecurity-stack_keypass_1        /opt/keypass/keypass-entry ...   Up      127.0.0.1:17070->7070/tcp, 127.0.0.1:7071->7071/tcp
telefonicasecurity-stack_keystone_1       /opt/keystone/keystone-ent ...   Up      35357/tcp, 0.0.0.0:5001->5001/tcp,:::5001->5001/tcp
telefonicasecurity-stack_mongo_1          docker-entrypoint.sh mongod      Up      27017/tcp
telefonicasecurity-stack_mysql_1          docker-entrypoint.sh mysqld      Up      3306/tcp, 33060/tcp
telefonicasecurity-stack_orchestrator_1   /opt/orchestrator/bin/orch ...   Up      0.0.0.0:8084->8084/tcp,:::8084->8084/tcp
telefonicasecurity-stack_orion_1          /usr/bin/contextBroker -fg ...   Up      1026/tcp
telefonicasecurity-stack_pep-orion_1      docker/entrypoint.sh --  c ...   Up      0.0.0.0:1026->1026/tcp,:::1026->1026/tcp,
                                                                                   0.0.0.0:11211->11211/tcp,:::11211->11211/tcp
```

### Sanity check

Check the service is ready by executing the following command.

```
curl localhost:1026/version
```

```
{"name":"MISSING_HEADERS","message":"Some headers were missing from the request: [\"fiware-service\",\"fiware-servicepath\",\"x-auth-token\"]"}
```

## Add a broker

Add a broker to NGSI Go configuration.

```
ngsi broker add \
  --host orion-with-keystone \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --service fiware \
  --idmType thinkingcities \
  --idmHost http://localhost:5001/v3/auth/tokens \
  --username orion \
  --password 2HNzujGZ60NIyGs9 \
```
## Access the broker

The following commands allow you to create and get a entity in an Orion context broker

```
ngsi create --host orion-with-keystone entity --keyValues --data '{"id":"id001"}'
```

```
ngsi list --host orion-with-keystone entities --pretty
```

```
[
  {
    "id": "id001",
    "type": "Thing"
  }
]
```

## Additional information

## Users

| User             | Username              | Password         | Description                                                                         |
| ---------------- | --------------------- | ---------------- | ----------------------------------------------------------------------------------- |
| Cloud admin      | cloud_admin           | proxypwd1234     | It is the user with more privileges. It is able to create and delete services.      |
| Service admin    | orion                 | 2HNzujGZ60NIyGs9 | It is able to administrate a service, creating subservices inside a service.        |
| Subservice admin | subservice_admin_user | fiwarepass12345  | It is able to create and modify entities and other operation within the subservice. |

### Reference

-   [Telefónica Thinking Cities - Read the Docs](https://thinking-cities.readthedocs.io/en/latest/)
