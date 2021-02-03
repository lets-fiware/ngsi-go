# NGSI Go tutorial

## Getting Started with NGSI Go

### Start

```console
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go
script/tutorial.sh start
```

> If CrateDB exits immediately with a
> `max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]` error, this can be fixed
> by running the `sudo sysctl -w vm.max_map_count=262144` command on the host machine. For further information look within
> the CrateDB [documentation](https://crate.io/docs/crate/howtos/en/latest/admin/bootstrap-checks.html#bootstrap-checks)
> and Docker
> [troubleshooting guide](https://crate.io/docs/crate/howtos/en/latest/deployment/containers/docker.html#troubleshooting)

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
