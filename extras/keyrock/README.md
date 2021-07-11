# Access a broker with Keyrock

This documentation explains how to access an Orion Context Broker that the endpoints of NGSI API
are protected by Keyrock.

## Prepair

Clone the NGSI Go repository and move to `ngsi-go/extras/keyrock` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/keyrock
```

## Start up

Start up an Orion context broker and Keyrock.

```
docker-compose up -d
```

You can see seven containers.

```
docker-compose ps
```

```
        Name                       Command                       State                                 Ports
-----------------------------------------------------------------------------------------------------------------------------------
keyrock_keyrock_1       /nodejs/bin/node ./bin/www       Up (healthy)            0.0.0.0:3000->3000/tcp,:::3000->3000/tcp
keyrock_mongo_1         docker-entrypoint.sh --noj ...   Up                      27017/tcp
keyrock_mysql_1         docker-entrypoint.sh mysqld      Up                      3306/tcp, 33060/tcp
keyrock_orion-proxy_1   /nodejs/bin/node ./bin/www       Up (health: starting)   0.0.0.0:1026->1026/tcp,:::1026->1026/tcp, 1027/tcp
keyrock_orion_1         /usr/bin/contextBroker -fg ...   Up                      1026/tcp
```

### Sanity check

Check the service is ready by executing the following command.

```
curl localhost:1026/version
```

```
Auth-token not found in request header
```

## Add a broker

Add a broker to NGSI Go configuration.

```
ngsi broker add --host orion-with-keyrock \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType keyrock \
  --idmHost http://localhost:3000/oauth2/token \
  --username admin@test.com \
  --password 1234 \
  --clientId a1a6048b-df1d-4d4f-9a08-5cf836041d14 \
  --clientSecret e4cc0147-e38f-4211-b8ad-8ae5e6a107f9
```
## Access the broker

The following command allows you to access the broker with Keyrock.

```
ngsi version --host orion-with-keyrock
```

## Additional information

| Key                | Value                                          |
| ------------------ | ---------------------------------------------- |
| IDM_DB_NAME        | idm                                            |
| IDM_DB_USER        | idm                                            |
| IDM_DB_PASS        | keyrock2020                                    |
| IDM_ADMIN_USER     | admin                                          |
| IDM_ADMIN_EMAIL    | admin@test.com                                 |
| IDM_ADMIN_PASS     | 1234                                           |
| Client ID          | a1a6048b-df1d-4d4f-9a08-5cf836041d14           |
| Client secret      | e4cc0147-e38f-4211-b8ad-8ae5e6a107f9           |
| PEP_PROXY_APP_ID   | a1a6048b-df1d-4d4f-9a08-5cf836041d14           |
| PEP_PROXY_USERNAME | pep_proxy_58b1a6db-1bc0-4323-837e-f100511af19c |
| PEP_PASSWORD       | pep_proxy_1bad5dbf-7ae9-49a8-b0f6-c66e4570357a |
