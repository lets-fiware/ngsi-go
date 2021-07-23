# Geoproxy server

This documentation describes the geoproxy server that provides POST-based Query API Endpoint.

If you request too long URL, then it may give “414 Request URI too large” error message. The reason for those
"really long URLs" are the URI parameters such as `coords`, `q`, `attrs` for GET /v2/entities. The command
solves this problem by POST-based Query `POST /v2/ex/entities`.

## Prepair

Clone the NGSI Go repository and move to `ngsi-go/extras/geoproxy` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/geoproxy
```

## Start up

Start up a Geoproxy server and an Orion context broker.

```
docker-compose up -d
```

You can see four containers.

```
            Name                          Command               State                Ports
--------------------------------------------------------------------------------------------------------
geoproxy_geoproxy_1   /usr/local/bin/ngsi --stde ...   Up
geoproxy_mongo_1      docker-entrypoint.sh --noj ...   Up      27017/tcp
geoproxy_nginx_1      /docker-entrypoint.sh ngin ...   Up      0.0.0.0:1026->1026/tcp,:::1026->1026/tcp, 0.0.0.0:1030->1030/tcp,:::1030->1030/tcp, 80/tcp
geoproxy_orion_1      /usr/bin/contextBroker -fg ...   Up      1026/tcp
```

### Sanity check

Check the geoproxy is ready by executing the following command.

```
ngsi server add --host geoproxy --serverType geoproxy --serverHost http://localhost:1026
```

```
ngsi geoproxy health --host geoproxy
```

```
{
  "ngsi-go": "geoproxy",
  "version": "0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)",
  "health": "OK",
  "orion": "http://orion:1026/v2/entities",
  "verbose": true,
  "uptime": "0 d, 0 h, 0 m, 1 s",
  "timesent": 0,
  "success": 0,
  "failure": 0
}
```

Check the Orion is ready by executing the following command.

```
curl localhost:1026/version
```

## Create an entity

```
curl http://localhost:1026/v2/entities?options=keyValues --data '{"id":"device001","type":"Device","temperature":30}' -H "Content-type: application/json"
```

## Get a entity

## GET /v2/entities?type=Device

```
curl http://localhost:1026/v2/entities -G --data "type=Device"
```

```
[{"id":"device001","type":"Device","temperature":{"type":"Number","value":30,"metadata":{}}}]
```

## POST /v2/entities

```
curl http://localhost:1026/v2/entities --data "type=Device"
```

```
{
  "error": "UnsupportedMediaType",
  "description": "not supported content type: application/x-www-form-urlencoded"
}
```

## POST /v2/ex/entities

```
curl http://localhost:1026/v2/ex/entities --data "type=Device"
```

```
[{"id":"device001","type":"Device","temperature":{"type":"Number","value":30,"metadata":{}}}]
```
