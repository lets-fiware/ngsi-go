# broker - Management command

-   [List brokers](#list-brokers)
-   [Add broker](#add-broker)
-   [Update broker](#update-broker)
-   [Delete broker](#delete-broker)

## List brokers

```
ngsi broker list [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --host value, -h value          | specify host or alias                        |
| --json, -j                      | print JSON format (default: false)           |
| --help                          | show help (default: false)                   |

#### Example 1

```
$ ngsi broker list
orion orion-ld mylab
```

#### Example 2

```
$ ngsi broker list --host orion
brokerHost http://localhost:1026
ngsiType v2
safeString on
```

#### Example 3

```
$ ngsi broker list --host orion --json
{"brokerHost":"http://localhost:1026","ngsiType":"v2","safeString":"on"}
```

## Add broker

```
ngsi broker list [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --host value, -h value          | specify host or alias                        |
| --brokerHost value, -b value    | specify context broker host                  |
| --ngsiType value                | specify NGSI type: v2 or ld (default: ld)    |
| --idmType value, -t value       | specify token type                           |
| --idmHost value, -m value       | specify identity manager host                |
| --apiPath value, -a value       | specify API path                             |
| --username value, -U value      | specify username                             |
| --password value, -P value      | specify password                             |
| --clientId value, -I value      | specify client id                            |
| --clientSecret value, -S value  | specify client secret                        |
| --token value                   | specify oauth token                          |
| --service value, -s value       | specify FIWARE Service                       |
| --path value, -p value          | specify FIWARE ServicePath                   |
| --safeString value              | Use safe string: `off` or `on` (default: on) |
| --help                          | show help (default: false)                   |

#### Example 1

NGSIv2 (for FIWARE Orion)

```
$ ngsi broker add \
  --host orion \
  --brokerHost http://localhost:1026 \
  --ngsiType v2
```

#### Example 2

NGSI-LD Broker

```
$ ngsi broker add \
  --host orion-ld \
  --brokerHost http://localhost:1026
```

#### Example 3

Orion-LD with Keyrock

```
$ ngsi broker add \
  --host orion-ld \
  --brokerHost https://orion-ld \
  --idmType keyrock \
  --idmHost https://keyrock \
  --username keyrock001@letsfiware.jp \
  --password 0123456789 \
  --clientId 00000000-1111-2222-3333-444444444444 \
  --clientSecret 55555555-6666-7777-8888-999999999999
```

### NGSI type

Specify `v2` to `--ngsiType` when you add an alias for FIWARE Orion Context Broker.

### Parameters for Identity Managers

| idmType              | Required parameters                                 | Description                                                                  |
| -------------------- | --------------------------------------------------- | ---------------------------------------------------------------------------- |
| password             | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keycloak, WSO2, Apinf, and Stellio) |
| keyrock              | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keyrock                             |
| KeyrockTokenProvider | idmHost, username, password                         | It provides auth token from Keyrock                                          |
| tokenProxy           | idmHost, username, password                         | It provides auth token from Keyrock                                          |

### FIWARE Serivce and FIWARE ServicePath

Specify the `--service` and/or `--path` parameter when adding a new alias.

#### Example 4

```
$ ngsi broker add \
  --host myinstance \
  --brokerHost http://localhost:1026 \
  --ngsiType v2
  --service open
  --path /iot
```

You can add a new alias using an exising alias.
Specify an existing alias to the `--brokerHost` parameter when adding a new alias.

#### Example 5

```
$ ngsi broker add \
  --host myinstance \
  --brokerHost orion \
  --service open
  --path /iot
```
### API Path

The NGSI Go assumes that the root of NGSI API is a root of URL.

```
https://orion.letsfiware.jp/v2/entities
```

You should use the `--apiPath` parameter if your broker has a special URL.

If the root of the NGSI API is in a sub-directory:

```
https://fiware-server/orion/v2/entities
```

You should set the `--apiPath` parameter to as shown:

```
--apiPath "/,/orion"
```

If the path of NGSI API is changed to a special path:

```
https://fiware-server/orion/v2.0/entities
```

You should set the `--apiPath` parameter to as shown:

```
--apiPath "/v2,/orion/v2.0"
```

## Update broker

```
ngsi broker upadte [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --host value, -h value          | specify host or alias (Required)             |
| --brokerHost value, -b value    | specify context broker host                  |
| --ngsiType value                | specify NGSI type: v2 or ld (default: ld)    |
| --idmType value, -t value       | specify token type                           |
| --idmHost value, -m value       | specify identity manager host                |
| --apiPath value, -a value       | specify API path                             |
| --username value, -U value      | specify username                             |
| --password value, -P value      | specify password                             |
| --clientId value, -I value      | specify client id                            |
| --clientSecret value, -S value  | specify client secret                        |
| --token value                   | specify oauth token                          |
| --service value, -s value       | specify FIWARE Service                       |
| --path value, -p value          | specify FIWARE ServicePath                   |
| --safeString value              | Use safe string: `off` or `on` (default: on) |
| --help                          | show help (default: false)                   |

#### Example 1

```
$ ngsi broker update --host orion --ngsiType v2
```

## Delete broker

```
$ ngsi broker upadte [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --host value, -h value          | specify host or alias (Required)             |
| --help                          | show help (default: false)                   |

#### Example 1

```
$ ngsi broker delete --host orion
```
