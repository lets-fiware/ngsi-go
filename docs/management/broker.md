# broker - Management command

-   [List brokers](#list-brokers)
-   [Get broker](#get-broker)
-   [Add broker](#add-broker)
-   [Update broker](#update-broker)
-   [Delete broker](#delete-broker)

<a name="list-brokers"></a>

## List brokers

```console
ngsi broker list [options]
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
ngsi broker list
```

```text
orion orion-ld mylab
```

#### Example 2

```console
ngsi broker list --host orion
```

```text
brokerHost http://localhost:1026
ngsiType v2
safeString on
```

#### Example 3

```console
ngsi broker list --json
```

```json
{"brokerHost":"http://localhost:1026","ngsiType":"v2","safeString":"on"}
```

<a name="get-broker"></a>

## Get broker

```console
ngsi broker get [options]
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
ngsi broker get --host orion
```

```text
brokerHost http://orion:1026
ngsiType v2
```

#### Example 2

```console
ngsi broker get --host orion --json
```

```text
{"brokerHost":"http://localhost:1026","ngsiType":"v2","safeString":"off"}
```

<a name="add-broker"></a>

## Add broker

```console
ngsi broker add [options]
```

### Options

| Options                        | Description                                               |
| ------------------------------ | --------------------------------------------------------- |
| --host value, -h value         | specify host or alias                                     |
| --brokerHost value, -b value   | specify context broker host                               |
| --ngsiType value               | specify NGSI type: v2 or ld                               |
| --brokerType value             | specify NGSI-LD broker type: orion-ld, scorpio or stellio |
| --idmType value, -t value      | specify token type                                        |
| --idmHost value, -m value      | specify identity manager host                             |
| --apiPath value, -a value      | specify API path                                          |
| --username value, -U value     | specify username                                          |
| --password value, -P value     | specify password                                          |
| --clientId value, -I value     | specify client id                                         |
| --clientSecret value, -S value | specify client secret                                     |
| --token value                  | specify oauth token                                       |
| --service value, -s value      | specify FIWARE Service                                    |
| --path value, -p value         | specify FIWARE ServicePath                                |
| --safeString value             | Use safe string: `off` or `on`                            |
| --help                         | show help (default: false)                                |

-   The `--brokerType` option is used with `--ngsiType ld` option when adding a NGSI-LD broker. It is not needed for
    NGSIv2.

> **Note:** Orion interprets the FIWARE Service name (tenant name) in lowercase. To use a coherent FIWARE Service name,
> NGSI Go allows only lowercase letters in FIWARE Service name. Please have a look at
> [MULTI TENANCY section in Orion documentation](https://fiware-orion.readthedocs.io/en/master/user/multitenancy/index.html).

#### Example 1

NGSIv2 (for FIWARE Orion)

```console
ngsi broker add \
  --host orion \
  --brokerHost http://localhost:1026 \
  --ngsiType v2
```

#### Example 2

NGSI-LD Broker

```console
ngsi broker add \
  --host orion-ld \
  --brokerHost http://localhost:1026 \
  --ngsiType ld
```

#### Example 3

NGSI-LD Broker

```console
ngsi broker add \
  --host scorpio \
  --brokerHost http://localhost:9090 \
  --ngsiType ld \
  --brokerType scorpio
```

#### Example 4

Orion-LD with Keyrock

```console
ngsi broker add \
  --host orion-ld \
  --ngsiType ld \
  --brokerHost https://orion-ld \
  --idmType keyrock \
  --idmHost https://keyrock \
  --username keyrock001@letsfiware.jp \
  --password 0123456789 \
  --clientId 00000000-1111-2222-3333-444444444444 \
  --clientSecret 55555555-6666-7777-8888-999999999999
```

#### Example 5

Telefónica Thinking Cities (Keystone)

```console
ngsi broker add \
  --host myinstance \
  --ngsiType v2 \
  --brokerHost http://localhost:1026 \
  --idmType ThinkingCities \
  --idmHost http://localhost:5001/v3/auth/tokens \
  --username usertest \
  --password '<ofuscated>' \
  --service smartcity
```

### NGSI type

Specify `v2` to `--ngsiType` when you add an alias for FIWARE Orion Context Broker.

### Parameters for Identity Managers

| idmType                                                                    | Required parameters                                 | Description                                                                  |
| -------------------------------------------------------------------------- | --------------------------------------------------- | ---------------------------------------------------------------------------- |
| password                                                                   | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keycloak, WSO2, Apinf, and Stellio) |
| [keyrock](https://fiware-idm.readthedocs.io/)                              | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keyrock                             |
| [KeyrockTokenProvider](https://github.com/FIWARE-Ops/KeyrockTokenProvider) | idmHost, username, password                         | It provides auth token from Keyrock                                          |
| tokenproxy                                                                 | idmHost, username, password                         | It provides auth token from Keyrock                                          |
| [ThinkingCities](https://thinking-cities.readthedocs.io/)                  | idmHost, username, password                         | It provides auth token from Keystone                                         |

### FIWARE Service and FIWARE ServicePath

Specify the `--service` and/or `--path` parameter when adding a new alias.

#### Example 1

```console
ngsi broker add \
  --host myinstance \
  --brokerHost http://localhost:1026 \
  --ngsiType v2 \
  --service open \
  --path /iot
```

You can add a new alias using an exising alias.
Specify an existing alias to the `--brokerHost` parameter when adding a new alias.

#### Example 2

```console
ngsi broker add \
  --host myinstance \
  --brokerHost orion \
  --service open \
  --path /iot
```

### API Path

The NGSI Go assumes that the root of NGSI API is a root of URL.

```console
https://orion.letsfiware.jp/v2/entities
```

You should use the `--apiPath` parameter if your broker has a special URL.

If the root of the NGSI API is in a sub-directory:

```console
https://fiware-server/orion/v2/entities
```

You should set the `--apiPath` parameter to as shown:

```text
--apiPath "/,/orion"
```

If the path of NGSI API is changed to a special path:

```text
https://fiware-server/orion/v2.0/entities
```

You should set the `--apiPath` parameter to as shown:

```text
--apiPath "/v2,/orion/v2.0"
```

<a name="update-broker"></a>

## Update broker

```console
ngsi broker upadte [options]
```

### Options

| Options                        | Description                                               |
| ------------------------------ | --------------------------------------------------------- |
| --host value, -h value         | specify host or alias (Required)                          |
| --brokerHost value, -b value   | specify context broker host                               |
| --ngsiType value               | specify NGSI type: v2 or ld                               |
| --brokerType value             | specify NGSI-LD broker type: orion-ld, scorpio or stellio |
| --idmType value, -t value      | specify token type                                        |
| --idmHost value, -m value      | specify identity manager host                             |
| --apiPath value, -a value      | specify API path                                          |
| --username value, -U value     | specify username                                          |
| --password value, -P value     | specify password                                          |
| --clientId value, -I value     | specify client id                                         |
| --clientSecret value, -S value | specify client secret                                     |
| --token value                  | specify oauth token                                       |
| --service value, -s value      | specify FIWARE Service                                    |
| --path value, -p value         | specify FIWARE ServicePath                                |
| --safeString value             | Use safe string: `off` or `on`                            |
| --help                         | show help (default: false)                                |

#### Example 1

```console
ngsi broker update --host orion --ngsiType v2
```

<a name="delete-broker"></a>

## Delete broker

```console
ngsi broker upadte [options]
```

### Options

| Options                | Description                      |
| ---------------------- | -------------------------------- |
| --host value, -h value | specify host or alias (Required) |
| --help                 | show help (default: false)       |

#### Example 1

```console
ngsi broker delete --host orion
```
