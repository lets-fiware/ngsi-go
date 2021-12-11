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
| --host VALUE, -h VALUE | context broker host alias                            |
| --json, -j             | JSON format (default: false)                         |
| --pretty, -P           | pretty format (default: false)                       |
| --clearText            | show obfuscated items as clear text (default: false) |
| --singleLine, -1       | list one file per line (default: false)              |
| --help                 | show help (default: true)                            |

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
| --host VALUE, -h VALUE | context broker host alias (required)                 |
| --json, -j             | JSON format (default: false)                         |
| --pretty, -P           | pretty format (default: false)                       |
| --clearText            | show obfuscated items as clear text (default: false) |
| --help                 | show help (default: true)                            |

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

| Options                        | Description                                       |
| ------------------------------ | ------------------------------------------------- |
| --host VALUE, -h VALUE         | context broker host alias (required)              |
| --brokerHost VALUE, -b VALUE   | context broker host address or alias              |
| --ngsiType VALUE               | NGSI type: v2 or ld                               |
| --brokerType VALUE             | NGSI-LD broker type: orion-ld, scorpio or stellio |
| --idmType VALUE, -t VALUE      | token type                                        |
| --idmHost VALUE, -m VALUE      | identity manager host                             |
| --apiPath VALUE, -a VALUE      | API path                                          |
| --username VALUE, -U VALUE     | username                                          |
| --password VALUE, -P VALUE     | password                                          |
| --clientId VALUE, -I VALUE     | client id                                         |
| --clientSecret VALUE, -S VALUE | client secret                                     |
| --headerName VALUE             | header name for apikey                            |
| --headerValue VALUE            | header value for apikey                           |
| --headerEnvValue VALUE         | name of environment variable for apikey           |
| --tokenScope VALUE             | scope for token                                   |
| --token VALUE                  | token VALUE                                       |
| --service VALUE, -s VALUE      | FIWARE Service VALUE                              |
| --path VALUE, -p VALUE         | FIWARE ServicePath VALUE                          |
| --safeString VALUE             | use safe string (VALUE: on/off)                   |
| --overWrite, -O                | overwrite broker alias (default: false)           |
| --help                         | show help (default: true)                         |

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
  --idmHost https://keyrock/oauth2/token \
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

#### Example 6

Orion with Keycloak

```
ngsi broker add \
  --host orion-with-keycloak \
  --ngsiType v2 \
  --brokerHost http://localhost:1026 \
  --idmType keycloak \
  --idmHost http://keycloak:8080/auth/realms/fiware_service \
  --username fiware \
  --password 1234 \
  --clientId ngsi_api \
  --clientSecret 8eb5d01d-d155-4b73-9414-a3c28ee4aba6
```

#### Example 7

Orion with WSO2

```
ngsi broker add \
  --host orion-with-wso2 \
  --ngsiType v2 \
  --brokerHost http://localhost:1026 \
  --idmType wso2 \
  --idmHost http://wso2am:8243/token \
  --username fiware \
  --password 1234 \
  --clientId 0000000000000000000000_A_ZZZ \
  --clientSecret 00000000-1111-2222-3333-444444444444
```

#### Example 8

Orion with Kong (client credentials)

```
ngsi broker add \
  --host kong \
  --ngsiType v2 \
  --brokerHost http://localhost:8443/ngsi \
  --idmType kong \
  --idmHost "https://localhost:8443/ngsi/oauth2/token,http://localhost:8001/" \
  --clientId orion \
  --clientSecret 1234
```

#### Example 9

Orion with Basic authentication

```console
ngsi broker add \
  --host orion-with-basic-auth \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType basic \
  --username fiware \
  --password 1234
```

#### Example 11

Orion with APIKEY

```console
ngsi broker add --host orion-with-apikey \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType apikey \
  --headerName Authorization \
  --headerValue "Bearer magic1234"
```

#### Example 12

Orion with APIKEY (environment value)

```console
export TOKEN="Bearer magic1234"
ngsi broker add --host orion-with-apikey-env \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType apikey \
  --headerName Authorization \
  --headerEnvValue TOKEN
```

### NGSI type

Specify `v2` to `--ngsiType` when you add an alias for FIWARE Orion Context Broker.

### Parameters for Identity Managers

| idmType                                                                    | Required parameters                                 | Description                                            |
| -------------------------------------------------------------------------- | --------------------------------------------------- | ------------------------------------------------------ |
| basic                                                                      | username, password                                  | Basic authentication                                   |
| password                                                                   | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials.                 |
| [keyrock](https://fiware-idm.readthedocs.io/)                              | idmHost, username, password, clientId, clientSecret | This type is for Password Credentials of Keyrock.      |
| [KeyrockTokenProvider](https://github.com/FIWARE-Ops/KeyrockTokenProvider) | idmHost, username, password                         | It provides auth token from Keyrock.                   |
| tokenproxy                                                                 | idmHost, username, password                         | It provides auth token from Keyrock.                   |
| [ThinkingCities](https://thinking-cities.readthedocs.io/)                  | idmHost, username, password                         | It provides auth token from Keystone.                  |
| Keycloak                                                                   | idmHost, username, password, clientId, clientSecret | It provides auth token from Keycloak.                  |
| WSO2                                                                       | idmHost, username, password, clientId, clientSecret | It provides auth token from WSO2.                      |
| Kong (client credentials)                                                  | idmHost, clientId, clientSecret                     | It provides auth token from Kong.                      |
| apikey                                                                     | headerName, either headerValue or headerEnvValue    | It allows you to set a header name and a header value. |

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

| Options                        | Description                                       |
| ------------------------------ | ------------------------------------------------- |
| --host VALUE, -h VALUE         | context broker host alias (required)              |
| --brokerHost VALUE, -b VALUE   | context broker host address or alias              |
| --ngsiType VALUE               | NGSI type: v2 or ld                               |
| --brokerType VALUE             | NGSI-LD broker type: orion-ld, scorpio or stellio |
| --idmType VALUE, -t VALUE      | token type                                        |
| --idmHost VALUE, -m VALUE      | identity manager host                             |
| --apiPath VALUE, -a VALUE      | API path                                          |
| --username VALUE, -U VALUE     | username                                          |
| --password VALUE, -P VALUE     | password                                          |
| --clientId VALUE, -I VALUE     | client id                                         |
| --clientSecret VALUE, -S VALUE | client secret                                     |
| --headerName VALUE             | header name for apikey                            |
| --headerValue VALUE            | header value for apikey                           |
| --headerEnvValue VALUE         | name of environment variable for apikey           |
| --tokenScope VALUE             | scope for token                                   |
| --token VALUE                  | token VALUE                                       |
| --service VALUE, -s VALUE      | FIWARE Service VALUE                              |
| --path VALUE, -p VALUE         | FIWARE ServicePath VALUE                          |
| --safeString VALUE             | use safe string (VALUE: on/off)                   |
| --help                         | show help (default: true)                         |

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

| Options                 | Description                           |
| ----------------------- | ------------------------------------- |
| --host VALUE, -h VALUE  | context broker host alias (required)  |
| --help                  | show help (default: true)             |

#### Example 1

```console
ngsi broker delete --host orion
```
