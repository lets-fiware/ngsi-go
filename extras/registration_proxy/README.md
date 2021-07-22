# Registration proxy

This documentation explains how to get remote entities from a remote broker that is protected
by a security policy enforce point by accessing a local broker with registration.

## Prepair

Clone the NGSI Go repository and move to `ngsi-go/extras/registration_proxy` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/registration_proxy
```

## Add a remote broker

Add a remote broker to NGSI Go configuration. The host name should be `remote-orion:`.
The `ngsi-go-config.json` will be created in current directory.

```
ngsi --config ngsi-go-config.json \
  broker add \
  --host remote-orion \
  --ngsiType v2 \
  --brokerHost http://remote-orion:1026 \
  --idmType keyrock \
  --idmHost http://keyrock:3000/oauth2/token \
  --username admin@test.com \
  --password 1234 \
  --clientId a1a6048b-df1d-4d4f-9a08-5cf836041d14 \
  --clientSecret e4cc0147-e38f-4211-b8ad-8ae5e6a107f9
```

## Start up local broker

Start up a local broker and a regproxy.

```
docker-compose up --build -d
```

You can see three containers.

```
docker-compose ps
```

```
            Name                           Command               State                    Ports
-----------------------------------------------------------------------------------------------------------------
registration_proxy_keyrock_1        /nodejs/bin/node ./bin/www       Up (healthy)     0.0.0.0:3000->3000/tcp,:::3000->3000/tcp
registration_proxy_local-mongo_1    docker-entrypoint.sh --noj ...   Up               27017/tcp
registration_proxy_local-orion_1    /usr/bin/contextBroker -fg ...   Up               0.0.0.0:1026->1026/tcp,:::1026->1026/tcp
registration_proxy_mysql_1          docker-entrypoint.sh mysqld      Up               3306/tcp, 33060/tcp
registration_proxy_orion-proxy_1    /nodejs/bin/node ./bin/www       Up (unhealthy)   0.0.0.0:1027->1026/tcp,:::1027->1026/tcp, 1027/tcp
registration_proxy_regproxy_1       /usr/local/bin/ngsi --stde ...   Up               0.0.0.0:1028->1028/tcp,:::1028->1028/tcp
registration_proxy_remote-mongo_1   docker-entrypoint.sh --noj ...   Up               27017/tcp
registration_proxy_remote-orion_1   /usr/bin/contextBroker -fg ...   Up               1026/tcp
```

## Add regproxy server

```
ngsi server add --host regproxy \
  --serverType regproxy \
  --serverHost http://localhost:1028
```

The following command is a regproxy sanity check.

```
ngsi regproxy health --host regproxy --pretty
```

```
{
  "ngsi-go": "regproxy",
  "version": "0.8.4-next (git_hash:a86aaf88f0f9d08f8d5388f9e5e83b6ebb90e7f7)",
  "health": "OK",
  "csource": "http://remote-orion:1026",
  "verbose": true,
  "uptime": "0 d, 0 h, 0 m, 1 s",
  "timesent": 0,
  "success": 0,
  "failure": 0
}
```

## Add a remote orion

```
ngsi broker add --host remote-orion \
  --ngsiType v2 \
  --brokerHost http://localhost:1027/ \
  --idmType keyrock \
  --idmHost http://localhost:3000/oauth2/token \
  --username admin@test.com \
  --password 1234 \
  --clientId a1a6048b-df1d-4d4f-9a08-5cf836041d14 \
  --clientSecret e4cc0147-e38f-4211-b8ad-8ae5e6a107f9
```

## Create a remote entity in the remote orion

```
ngsi create --host remote-orion \
  --service federation \
  --path /iot entity \
  --keyValues \
   --data '{"id":"urn:ngsi-ld:Device:device001","type":"Device","temperature":30}'
```

## Add registration in a loca broker

```
curl -sS http://localhost:1026/v2/registrations \
-H 'Content-Type: application/json' \
-H 'Fiware-Service: federation' \
-H 'Fiware-Servicepath: /iot' \
-H 'Accept: application/json' -d @-  <<EOF
{
  "dataProvided": {
    "entities": [
      {
        "id": "urn:ngsi-ld:Device:device001",
        "type": "Device"
      }
    ],
    "attrs": [
      "temperature"
    ]
  },
  "provider": {
    "http": {
      "url": "http://regproxy:1028/v2"
    }
  }
}
EOF
```

## Get remote entities from remote orion

You can get remote entities from remote orion by accessing a local broker as shown:

```
curl -sS http://localhost:1026/v2/entities?type=Device -H 'Fiware-Service: federation' -H 'Fiware-Servicepath: /iot'
```

```
[{"id":"urn:ngsi-ld:Device:device001","type":"Device","temperature":{"type":"Number","value":30,"metadata":{}}}]
```

## Check regproxy status

```
ngsi regproxy health --host regproxy --pretty
```

```
{
  "ngsi-go": "regproxy",
  "version": "0.8.4-next (git_hash:a86aaf88f0f9d08f8d5388f9e5e83b6ebb90e7f7)",
  "health": "OK",
  "csource": "http://remote-orion:1026",
  "verbose": true,
  "uptime": "0 d, 0 h, 9 m, 0 s",
  "timesent": 1,
  "success": 1,
  "failure": 0
}
```
