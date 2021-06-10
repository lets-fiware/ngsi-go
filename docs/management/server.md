# server - Management command

-   [List servers](#list-servers)
-   [Get server](#get-server)
-   [Add server](#add-server)
-   [Update server](#update-server)
-   [Delete server](#delete-server)

<a name="list-servers"></a>

## List servers

```console
ngsi server list [options]
```

### Options

| Options                | Description                                          |
| ---------------------- | ---------------------------------------------------- |
| --host value, -h value | specify host or alias                                |
| --json, -j             | print JSON format (default: false)                   |
| --pretty, -P           | pretty format (default: false)                       |
| --clearText            | show obfuscated items as clear text (default: false) |
| --help                 | show help (default: false)                           |

#### Example 1

```console
ngsi server list
```

```text
comet 
```

<a name="get-server"></a>

## Get server

```console
ngsi server get [options]
```

### Options

| Options                | Description                                          |
| ---------------------- | ---------------------------------------------------- |
| --host value, -h value | specify host or alias                                |
| --json, -j             | print JSON format (default: false)                   |
| --pretty, -P           | pretty format (default: false)                       |
| --clearText            | show obfuscated items as clear text (default: false) |
| --help                 | show help (default: false)                           |

#### Example 1

```console
ngsi server list --host comet
```

```text
serverType comet
serverHost http://localhost:8666
FIWARE-Service openiot
FIWARE-ServicePath /
```

#### Example 2

```console
ngsi server get --host comet--json
```

```text
{"serverType":"comet","serverHost":"http://localhost:8666","tenant":"openiot","scope":"/"}
```

<a name="add-server"></a>

## Add server

```console
ngsi server add [options]
```

### Options

| Options                        | Description                    |
| ------------------------------ | ------------------------------ |
| --host value, -h value         | specify host or alias          |
| --serverHost value, -b value   | specify context server host    |
| --serverType value             | specify FIWARE GE Type         |
| --idmType value, -t value      | specify token type             |
| --idmHost value, -m value      | specify identity manager host  |
| --apiPath value, -a value      | specify API path               |
| --username value, -U value     | specify username               |
| --password value, -P value     | specify password               |
| --clientId value, -I value     | specify client id              |
| --clientSecret value, -S value | specify client secret          |
| --token value                  | specify oauth token            |
| --service value, -s value      | specify FIWARE Service         |
| --path value, -p value         | specify FIWARE ServicePath     |
| --safeString value             | Use safe string: `off` or `on` |
| --help                         | show help (default: false)     |

> **Note:** Orion interprets the FIWARE Service name (tenant name) in lowercase. To use a coherent FIWARE Service name,
> NGSI Go allows only lowercase letters in FIWARE Service name. Please have a look at
> [MULTI TENANCY section in Orion documentation](https://fiware-orion.readthedocs.io/en/master/user/multitenancy/index.html).

#### Example 1

Add QuantumLeap as an alias.

```console
ngsi server add --host ql \
  --serverType quantumleap \
  --serverHost http://quantumleap:8668
```

#### Example 2

Add STH-Comet with Keyrock as an alias.

```console
ngsi server add \
  --host comet \
  --serverType comet \
  --serverHost http://comet:8666 \
  --idmType keyrock \
  --idmHost https://keyrock \
  --username keyrock001@letsfiware.jp \
  --password 0123456789 \
  --clientId 00000000-1111-2222-3333-444444444444 \
  --clientSecret 55555555-6666-7777-8888-999999999999
```

### Server type

Specify the following value to `--serverType` option when you add an alias for FIWARE GE.

| FIWARE GE   | serverType  |
| ----------- | ----------- |
| STH-Comet   | comet       |
| QuantumLeap | quantumleap |
| IoT Agent   | iota        |
| Perseo FE   | perseo      |
| Perseo Core | perseo-core |
| Keyrock     | keyrock     |

#### Example 3

The following example is how to add Keyrock as a server type. Specify an admin user name of Keyrock to
`--username` option.

```console
ngsi server add \
  --host fiware-idm \
  --serverType keyrock \
  --serverHost https://idm.letsfiware.jp \
  --username admin@letsfiware.jp \
  --password 1234567
```

### Parameters for Identity Managers

| idmType              | Required parameters                                 | Description                                                                  |
| -------------------- | --------------------------------------------------- | ---------------------------------------------------------------------------- |
| password             | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keycloak, WSO2, Apinf, and Stellio) |
| keyrock              | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keyrock                             |
| KeyrockTokenProvider | idmHost, username, password                         | It provides auth token from Keyrock                                          |
| tokenproxy           | idmHost, username, password                         | It provides auth token from Keyrock                                          |

### FIWARE Service and FIWARE ServicePath

Specify the `--service` and/or `--path` parameter when adding a new alias.

#### Example 4

```console
ngsi server add \
  --host sth \
  --serverHost http://comet:8666 \
  --serverType comet \
  --service open \
  --path /iot
```

You can add a new alias using an exising alias.
Specify an existing alias to the `--serverHost` parameter when adding a new alias.

#### Example 5

```console
ngsi server add \
  --host sht \
  --serverHost comet \
  --service open \
  --path /iot
```

### API Path

The NGSI Go assumes that the root of FIWARE Open APIs is a root of URL.

```console
https://quantumleap.letsfiware.jp/v2/entities
```

You should use the `--apiPath` parameter if your server has a special URL.

If the root of the NGSI API is in a sub-directory:

```console
https://fiware-server/quantumleap/v2/entities
```

You should set the `--apiPath` parameter to as shown:

```text
--apiPath "/,/quantumleap"
```

If the path of NGSI API is changed to a special path:

```text
https://fiware-server/quantumleap/v2.0/entities
```

You should set the `--apiPath` parameter to as shown:

```text
--apiPath "/v2,/quantumleap/v2.0"
```

<a name="update-server"></a>

## Update server

```console
ngsi server upadte [options]
```

### Options

| Options                        | Description                      |
| ------------------------------ | -------------------------------- |
| --host value, -h value         | specify host or alias (Required) |
| --serverHost value, -b value   | specify context server host      |
| --idmType value, -t value      | specify token type               |
| --idmHost value, -m value      | specify identity manager host    |
| --apiPath value, -a value      | specify API path                 |
| --username value, -U value     | specify username                 |
| --password value, -P value     | specify password                 |
| --clientId value, -I value     | specify client id                |
| --clientSecret value, -S value | specify client secret            |
| --token value                  | specify oauth token              |
| --service value, -s value      | specify FIWARE Service           |
| --path value, -p value         | specify FIWARE ServicePath       |
| --safeString value             | Use safe string: `off` or `on`   |
| --help                         | show help (default: false)       |

#### Example 1

```console
ngsi server update --host comet --serverHost http://sth-comet:8666
```

<a name="delete-server"></a>

## Delete server

```console
ngsi server upadte [options]
```

### Options

| Options                | Description                      |
| ---------------------- | -------------------------------- |
| --host value, -h value | specify host or alias (Required) |
| --help                 | show help (default: false)       |

#### Example 1

```console
ngsi server delete --host comet
```
