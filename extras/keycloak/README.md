# Access a broker with Keycloak

This documentation explains how to access an Orion Context Broker that the endpoints of NGSI API
are protected by Keycloak.

## Start up

Start up an Orion context broker and Keycloak.

```
docker-compose up -d
```

You can see five containers.

```
docker-compose ps
```

```
       Name                      Command               State                                     Ports
-----------------------------------------------------------------------------------------------------------------------------------------
keycloak_keycloak_1   /opt/jboss/tools/docker-en ...   Up      8080/tcp, 8443/tcp
keycloak_mongo_1      docker-entrypoint.sh --noj ...   Up      27017/tcp
keycloak_nginx_1      /docker-entrypoint.sh ngin ...   Up      0.0.0.0:1026->1026/tcp,:::1026->1026/tcp, 0.0.0.0:80->80/tcp,:::80->80/tcp
keycloak_orion_1      /usr/bin/contextBroker -fg ...   Up      1026/tcp
keycloak_postgres_1   docker-entrypoint.sh postgres    Up      5432/tcp
```

## Add the broker

Add the broker to NGSI Go configuration.

```
ngsi broker add \
  --host orion-with-keycloak \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType keycloak \
  --idmHost http://localhost/auth/realms/fiware_service \
  --username fiware \
  --password 1234 \
  --clientId ngsi_api \
  --clientSecret 8eb5d01d-d155-4b73-9414-a3c28ee4aba6
```

## Access the broker

The following command allows you to access the broker with Basic authentication.

```
ngsi version --host orion-with-keycloak
```

## Additional information

| Items         | Contents                             |
| ------------- | ------------------------------------ |
| Keycloak host | http://localhost/                    |
| Orion host    | http://localhost:1026/               |
| Admin user    | name: admin, password: 1234          |
| User          | name: fiware, password: 1234         |
| Realm         | fiware_service                       |
| client id     | ngsi_api                             |
| client secret | 8eb5d01d-d155-4b73-9414-a3c28ee4aba6 |

### Authorization Basic

```
echo "ngsi_api:8eb5d01d-d155-4b73-9414-a3c28ee4aba6" | base64
```

```
bmdzaV9hcGk6OGViNWQwMWQtZDE1NS00YjczLTk0MTQtYTNjMjhlZTRhYmE2Cg==
```

### curl

```
curl localhost:1026/version
```

```
<html>
<head><title>403 Forbidden</title></head>
<body>
<center><h1>403 Forbidden</h1></center>
<hr><center>nginx/1.19.10</center>
</body>
</html>
```

```
export TOKEN=`ngsi token --host orion-with-keycloak`;echo $TOKEN
```

```
curl localhost:1026/version -H "Authorization: Bearer $TOKEN"
```
