# applications / pep - Keyrock command

This command allows you to manage PEP Proxies in an application for Keyrock.

-   [list pep proxies](#list-pep-proxies)
-   [create pep proxy](#create-a-pep-proxy)
-   [reset pep proxy](#reset-a-pep-proxy)
-   [delete pep proxy](#delete-a-pep-proxy)


## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-pep-proxies"></a>

## List PEP Proxies

This command lists PEP Proxies.

```console
ngsi applications [command options] pep --aid {id} list [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications pep --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "pep_proxy": {
    "id": "pep_proxy_d2d3c969-703d-4193-8278-fa0fb491dd82",
    "oauth_client_id": "8b58ecff-fb45-4811-945c-6f42339db06b"
  }
}
```

<a name="create-a-pep-proxy"></a>

## Create PEP Proxy

This command creates a PEP Proxy.

```console
ngsi application [command options] pep --aid {id} create [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --run                 | run command (default: false)   |
| --verbose, -v         | verbose (default: false)       |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications pep --aid 8b58ecff-fb45-4811-945c-6f42339db06b create --run --pretty
```

```console
{
  "pep_proxy": {
    "id": "pep_proxy_d2d3c969-703d-4193-8278-fa0fb491dd82",
    "password": "pep_proxy_ac5f951d-ae96-4cf8-95cb-92b476c32d27"
  }
}
```

<a name="reset-a-pep-proxy"></a>

## Reset a PEP Proxy

This command resets a PEP Proxy.

```console
ngsi application [command options] pep --aid {id} reset [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --run                 | run command (default: false)   |
| --verbose, -v         | verbose (default: false)       |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications pep --aid 8b58ecff-fb45-4811-945c-6f42339db06b reset --run --pretty
```

```console
{
  "new_password": "pep_proxy_7e6a8364-f129-4043-be4f-77887076d3a3"
}
```

<a name="delete-a-pep-proxy"></a>

## Delete a PEP Proxy

This command deletes a PEP Proxy.

```console
ngsi application [command options] pep --aid {id} delete [options]
```

### Options

| Options               | Description                |
| --------------------- | -------------------------- |
| --aid value, -i value | application id             |
| --run                 | run command (default: false)   |
| --help                | show help (default: false) |

### Examples

#### Request:

```console
ngsi applications pep --aid 8b58ecff-fb45-4811-945c-6f42339db06b delete --run
```
