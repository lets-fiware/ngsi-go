# Access a broker with APIKEY

This documentation explains how to access an Orion Context Broker that the endpoints of NGSI API
are protected by APIKEY.

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
        Name                      Command                   State                              Ports
---------------------------------------------------------------------------------------------------------------------------
apikey_keyrock_1       /nodejs/bin/node ./bin/www       Up (healthy)     0.0.0.0:3000->3000/tcp,:::3000->3000/tcp
apikey_mongo_1         docker-entrypoint.sh --noj ...   Up               27017/tcp
apikey_mysql_1         docker-entrypoint.sh mysqld      Up               3306/tcp, 33060/tcp
apikey_orion-proxy_1   /nodejs/bin/node ./bin/www       Up (unhealthy)   0.0.0.0:1026->1026/tcp,:::1026->1026/tcp, 1027/tcp
apikey_orion_1         /usr/bin/contextBroker -fg ...   Up               1026/tcp
```

### Sanity check

Check the service is ready by executing the following command.

```
curl localhost:1026/version
```

```
Auth-token not found in request header
```

## Add a broker with an immediate value.

Add a broker to NGSI Go configuration.

```
ngsi broker add --host orion-with-apikey \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType apikey \
  --headerName Authorization \
  --headerValue "Bearer magic1234"
```

The following command allows you to access the broker with Keyrock.

```
ngsi version --host orion-with-apikey
```

## Add a broker with an environment value.

Add a broker to NGSI Go configuration.

```
export TOKEN="Bearer magic1234"
ngsi broker add --host orion-with-apikey-env \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType apikey \
  --headerName Authorization \
  --headerEnvValue TOKEN
```

The following command allows you to access the broker with Keyrock.

```
ngsi version --host orion-with-apikey-env
```
