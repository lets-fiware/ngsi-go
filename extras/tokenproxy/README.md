# Tokenproxy

This documentation explains tokenproxy that provides auth token from Keyrock.

## Prepair

Clone the NGSI Go repository and move to `ngsi-go/extras/tokenproxy` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/tokenproxy
```

## Start up

Start up Tokenproxy, Keyrock and Orion context broker.

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
tokenproxy_keyrock_1       /nodejs/bin/node ./bin/www       Up (healthy)     0.0.0.0:3000->3000/tcp,:::3000->3000/tcp
tokenproxy_mongo_1         docker-entrypoint.sh --noj ...   Up               27017/tcp
tokenproxy_mysql_1         docker-entrypoint.sh mysqld      Up               3306/tcp, 33060/tcp
tokenproxy_orion-proxy_1   /nodejs/bin/node ./bin/www       Up (unhealthy)   0.0.0.0:1026->1026/tcp,:::1026->1026/tcp, 1027/tcp
tokenproxy_orion_1         /usr/bin/contextBroker -fg ...   Up               1026/tcp
tokenproxy_tokenproxy_1    /usr/local/bin/ngsi --stde ...   Up               0.0.0.0:1029->1029/tcp,:::1029->1029/tcp
```

### Sanity check

Check Tokenproxy is ready by executing the following command.

```
ngsi server add --host tokenproxy --serverType tokenproxy --serverHost http://0.0.0.0:1029
```

```
ngsi tokenproxy health --host tokenproxy --pretty
```

```
{
  "ngsi-go": "tokenproxy",
  "version": "0.8.4-next (git_hash:74e11de6ba1883ba35fad2187ce8c9ebce4dd8cc)",
  "health": "OK",
  "idm": "http://keyrock:3000/oauth2/token",
  "clientId": "a1a6048b-df1d-4d4f-9a08-5cf836041d14",
  "clientSecret": "e4cc0147-e38f-4211-b8ad-8ae5e6a107f9",
  "verbose": true,
  "uptime": "0 d, 0 h, 0 m, 1 s",
  "timesent": 0,
  "success": 0,
  "revoke": 0,
  "failure": 0
}
```

Check Orion is ready by executing the following command.

```
curl localhost:1026/version
```

```
Auth-token not found in request header
```

## Add a broker with token proxy.

Add a broker to NGSI Go configuration.

```
ngsi broker add --host orion-with-tokenproxy \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType tokenproxy \
  --idmHost http://localhost:1029 \
  --username admin@test.com\
  --password 1234
```

The following command allows you to access the broker with token proxy.

```
ngsi version --host orion-with-tokenproxy
```

## How to get a token

```
curl http://localhost:1029/token \
  --data "username=admin@test.com" \
  --data "password=1234"
```

```
{
  "access_token": "41cc7caa5f2ba2da9b250273b2445c0c5c3cb3d6",
  "token_type": "bearer",
  "expires_in": 3599,
  "refresh_token": "6fd8edc885580f50a9c37a1e43f505a68f40b6a2",
  "scope": [
    "bearer"
  ]
}
```

## How to revoke a token

```
curl http://localhost:1029/revoke \
  --data "token=6fd8edc885580f50a9c37a1e43f505a68f40b6a2"
```
