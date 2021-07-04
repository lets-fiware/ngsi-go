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
  --brokerHost https://remote-orion.letsfiware.jp/ \
  --idmType keyrock \
  --idmHost https://keyrock/oauth2/token \
  --username keyrock001@letsfiware.jp \
  --password 0123456789 \
  --clientId 00000000-1111-2222-3333-444444444444 \
  --clientSecret 55555555-6666-7777-8888-999999999999
```

```
ngsi --config ngsi-go-config.json version --host remote-orion
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
registration_proxy_mongo_1      docker-entrypoint.sh --noj ...   Up      27017/tcp
registration_proxy_orion_1      /usr/bin/contextBroker -fg ...   Up      0.0.0.0:1026->1026/tcp,:::1026->1026/tcp
registration_proxy_regproxy_1   /usr/local/bin/ngsi --conf ...   Up      0.0.0.0:1028->1028/tcp,:::1028->1028/tcp
```

This following command is a regproxy sanity check.

```
curl -sS http://localhost:1028/health
```

```
{
  "ngsi-go": {
    "version": "0.8.4 (git_hash:35ee7560911b403029880242789039f2532c817b)",
    "csource": "https://remote-orion.letsfiware.jp/"
  }
}
```

## Add registration

Create a registration in a local broker.

```
curl -sS http://localhost:1026/v2/registrations \
-H 'Content-Type: application/json' \
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
curl -sS http://localhost:1026/v2/entities
```
