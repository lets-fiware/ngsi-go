# Usage

## Syntax

ngsi [global options] command [common options] sub-command [options]

### NGSI Command

| command | sub-command  | Description         |
| ------- | ------------ | ------------------- |
| append  | attrs        | append attributes   |
| create  | entity       | create entity       |
|         | entities     | create entities     |
|         | subscription | create subscription |
|         | registration | create registration |
| delete  | entity       | delete entity       |
|         | entities     | delete entities     |
|         | attr         | delete attribute    |
|         | subscription | delete subscription |
|         | registration | delete registration |
| get     | entity       | get entity          |
|         | entities     | get entities        |
|         | attr         | get attribute       |
|         | attrs        | get attributes      |
|         | types        | get types           |
|         | subscription | get subscription    |
|         | registration | get registration    |
| list    | entities     | list entties        |
|         | types        | list types          |
|         | subscription | list subscription   |
|         | registration | list registration   |
| replace | entities     | replace entities    |
|         | attrs        | replace attrs       |
| update  | entities     | update entities     |
|         | attr         | update attribute    |
|         | attrs        | update attributes   |
|         | subscription | update subscription |
| upsert  | entities     | upsert entities     |

### Convenience command

| command  | sub-command  | Description                                                      |
| -------- | ------------ | ---------------------------------------------------------------- |
| cp       | -            | copy entities                                                    |
| wc       | -            | print number of entities, subscriptions, registrations, or types |
| man      | -            | print urls of document                                           |
| ls       | -            | list entities                                                    |
| rm       | -            | remove entities                                                  |
| template | subscription | create template of subscription                                  |
|          | registration | create template of registration                                  |
| version  | -            | print the version of Context Broker                              |

### Management commnad

| command  | sub-command  | Description     |
| -------- | ------------ | --------------- |
| broker   | list         | list brokers    |
|          | get          | get brokes      |
|          | add          | add brokes      |
|          | update       | update brokes   |
|          | delete       | delete brokes   |
| context  | list         | list @context   |
|          | add          | add @context    |
|          | update       | udpate @context |
|          | delete       | delete @context |
| settings | list         | list settings   |
|          | delete       | delete settings |
|          | clear        | clear settings  |
| token    | -            | manage token    |

## Global Options

| Options	     | Description                                      |
| -------------- | ------------------------------------------------ |
| --syslog LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --stderr LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --config FILE  | specify configuration FILE                       |
| --cache FILE   | specify cache FILE                               |
| --batch, -B    | don't use previous args (batch) (default: false) |
| --help         | show help (default: false)                       |
| --version, -v  | print the version (default: false)               |

## Common options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

## Safe string

```
$ ngsi broker get -host orion
{"brokerHost":"http://localhost:1026","ngsiType":"v2","safeString":"off"}
```

The value of the `name` attribute has a forbidden characters.

```
$ ngsi create entity --keyValues \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "<Lemonade>",
      "size": "S",
      "price": 99
}'
entityCreate006 400 Bad Request {"error":"BadRequest","description":"Invalid characters in attribute value"}
```

Create entity with `--safeString on`

```
$ ngsi create entity --keyValues --safeString on \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "<Lemonade>",
      "size": "S",
      "price": 99
}'
```

```
$ ngsi get attr --id urn:ngsi-ld:Product:110 --attrName name
"%3CLemonade%3E"
```

```
$ ngsi get attr --id urn:ngsi-ld:Product:110 --attrName name --safeString on
"<Lemonade>"
```

## Error message

An error message consists of a prefix and a body.

E.g. entityCreate006 400 Bad Request {"error":"BadRequest","description":"Invalid characters in attribute value"}

The error message has `entityCreate006` as a prefix. A prefix consists of a Go function name and a position in the funciton.
The function name is `entityCreate`. The position is 6th.

### Detailed error information

You can get a detailed error information by running a command with the `--stderr info` option.

```
$ ngsi --stderr info version --host http://192.168.11.0
version
version003 Get "http://192.168.11.0/version": dial tcp 192.168.11.0:80: connect: no route to host: no route to host
httpRequest003 Get "http://192.168.11.0/version": dial tcp 192.168.11.0:80: connect: no route to host
Get "http://192.168.11.0/version": dial tcp 192.168.11.0:80: connect: no route to host
dial tcp 192.168.11.0:80: connect: no route to host
connect: no route to host
no route to host
abnormal termination
```

-   The first line shows that the version command was run.
-   The last line shows that the command terminated abnormally.
-   The lines between the first line and the last one shows a stack that Go functions were called.
-   The second line shows that a Go function that returned an error to a user.
-   The line before the last one shows the Go function where the error occurred. In the case, the function is not
    a function of the NGSI Go so that it doesn't have a prefix.
