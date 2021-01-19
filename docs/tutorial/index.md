# NGSI Go tutorial

## Getting Started with NGSI Go

### Start

```console
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go
script/tutorial.sh start
```
### Shell

```console
script/tutorial.sh shell
```

```console
docker-compose exec ngsi-test ash
~/e2e # 
```

```console
ngsi broker list --pretty
```

```json
{
  "orion": {
    "serverType": "broker",
    "serverHost": "http://orion:1026",
    "ngsiType": "v2"
  },
  "orion-ld": {
    "serverType": "broker",
    "serverHost": "http://orion-ld:1026",
    "ngsiType": "ld"
  }
}
```

```console
ngsi server list --pretty
```

```json
{
  "comet": {
    "serverType": "comet",
    "serverHost": "http://comet:8666",
    "tenant": "openiot",
    "scope": "/"
  },
  "quantumleap": {
    "serverType": "quantumleap",
    "serverHost": "http://quantumleap:8668",
    "tenant": "openiot",
    "scope": "/"
  }
}
```

-   [NGSI-LD CRUD](ngsi-ld-crud.md)
-   [NGSIv2 CRUD](ngsi-v2-crud.md)
-   [STH-Comet](comet.md)
-   [QuantumLeap](quantumleap.md)

### Stop

```console
script/tutorial.sh stop
```
