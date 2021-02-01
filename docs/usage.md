# Usage

-   [Syntax](#syntax)
    -   [NGSI Command](#ngsi-command)
    -   [Time series command](#time-series-command)
    -   [Convenience command](#convenience-command)
    -   [Management commnad](#management-commnad)
    -   [Global Options](#global-options)
    -   [Common options](#common-options)
-   [DateTime options](#datetime-options)
-   [--data option](#data-option)
-   [Safe string](#safe-string)
-   [Error message](#error-message)
    -   [Detailed error information](#detailed-error-information)

<a name="syntax"/>

## Syntax

```console
ngsi [global options] command [common options] sub-command [options]
```

<a name="ngsi-command"/>

### NGSI command

| command                      | sub-command  | Description         |
| ---------------------------- | ------------ | ------------------- |
| [append](./ngsi/append.md)   | attrs        | append attributes   |
| [create](./ngsi/create.md)   | entity       | create entity       |
|                              | entities     | create entities     |
|                              | subscription | create subscription |
|                              | registration | create registration |
| [delete](./ngsi/delete.md)   | entity       | delete entity       |
|                              | entities     | delete entities     |
|                              | attr         | delete attribute    |
|                              | subscription | delete subscription |
|                              | registration | delete registration |
| [get](./ngsi/get.md)         | entity       | get entity          |
|                              | entities     | get entities        |
|                              | attr         | get attribute       |
|                              | attrs        | get attributes      |
|                              | types        | get types           |
|                              | subscription | get subscription    |
|                              | registration | get registration    |
| [list](./ngsi/list.md)       | entities     | list entties        |
|                              | types        | list types          |
|                              | subscription | list subscription   |
|                              | registration | list registration   |
| [replace](./ngsi/replace.md) | entities     | replace entities    |
|                              | attrs        | replace attrs       |
| [update](./ngsi/update.md)   | entities     | update entities     |
|                              | attr         | update attribute    |
|                              | attrs        | update attributes   |
|                              | subscription | update subscription |
| [upsert](./ngsi/upsert.md)   | entities     | upsert entities     |

<a name="time-series-command"/>

### Time series command

| command                             | sub-command | Description                                                           |
| ----------------------------------- | ----------- | --------------------------------------------------------------------- |
| [hdelete](./time_series/hdelete.md) | attr        | delete all the data associated to certain attribute of certain entity |
|                                     | entity      | delete historical data of a certain entity                            |
|                                     | entities    | delete historical data of all entities of a certain type              |
| [hget](./time_series/hget.md)       | attr        | get hstory of an attribute                                            |
|                                     | attrs       | get history of attributes                                             |
|                                     | entities    | list of all the entity id                                             |

<a name="convenience-command"/>

### Convenience command

| command                               | sub-command     | Description                                                      |
| ------------------------------------- | --------------- | ---------------------------------------------------------------- |
| [admin](./convenience/admin.md)       | log             | print or set logging level for FIWARE Orion                      |
|                                       | trace           | print, set or delete trace level for FIWARE Orion                | 
|                                       | semaphore       | print semaphore for FIWARE Orion                                 |
|                                       | metrics         | print, reset or delete metrics for FIWARE Orion                  |
|                                       | statistics      | print or delete statistics for FIWARE Orion                      |
|                                       | cacheStatistics | print or delete cache statistics for FIWARE Orion                |
| [apis](.convenience/apis.md)          | -               | print endpoints of FWARE Open APIs                               |
| [cp](./convenience/cp.md)             | -               | copy entities                                                    |
| [wc](./convenience/wc.md)             | -               | print number of entities, subscriptions, registrations, or types |
| [man](./convenience/man.md)           | -               | print urls of document                                           |
| [health](./convenience/health.md)     | -               | print health status of FIWARE GEs                                |
| [ls](./convenience/ls.md)             | -               | list entities                                                    |
| [rm](./convenience/rm.md)             | -               | remove entities                                                  |
| [receiver](./convenience/receiver.md) | -               | notification receiver                                            |
| [template](./convenience/template.md) | subscription    | create template of subscription                                  |
|                                       | registration    | create template of registration                                  |
| [version](./convenience/version.md)   | -               | print the version of Context Broker                              |

<a name="management-commnad"/>

### Management commnad

| command                              | sub-command  | Description     |
| ------------------------------------ | ------------ | --------------- |
| [broker](./management/broker.md)     | list         | list brokers    |
|                                      | get          | get broker      |
|                                      | add          | add broker      |
|                                      | update       | update broker   |
|                                      | delete       | delete broker   |
| [context](./management/context.md)   | list         | list @context   |
|                                      | add          | add @context    |
|                                      | update       | udpate @context |
|                                      | delete       | delete @context |
| [settings](./management/settings.md) | list         | list settings   |
|                                      | delete       | delete settings |
|                                      | clear        | clear settings  |
| [server](./management/server.md)     | list         | list servers    |
|                                      | get          | get server      |
|                                      | add          | add server      |
|                                      | update       | update server   |
|                                      | delete       | delete server   |
| [token](./management/token.md)       | -            | manage token    |

<a name="global-options"/>

### Global Options

| Options        | Description                                      |
| -------------- | ------------------------------------------------ |
| --syslog LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --stderr LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --config FILE  | specify configuration FILE                       |
| --cache FILE   | specify cache FILE                               |
| --batch, -B    | don't use previous args (batch) (default: false) |
| --help         | show help (default: false)                       |
| --version, -v  | print the version (default: false)               |

<a name="common-options"/>

### Common options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="datetime-options"/>

## DateTime options

Some commands have the following options for specifying the date and time:

- expires {value}
- fromDate {value}
- toDate {value}

These options can have a value as shown:

| Values          | Examples                                          |
| --------------- | ------------------------------------------------- |
| ISO8601         | YYYY-MM-DDThh:mm:ss.ssZ, YYYY-MM-DDThh:mm:ss.sssZ |
| year, years     | 1year, 3years, -5years                            |
| month, months   | 1month, 11months, -3months                        |
| day, days       | 1day, 3days, -10days                              |
| hour, hours     | 1hour, 5hours, -2hours                            |
| minute, minutes | 1minute, 7minutes, -1minute                       |

You can specify a negative value for a date and time in the past.

### Examples

Specify a future date for an expiration date.

```console
ngsi create --host orion subscription --idPattern ".*" --type Sensor \
--wAttrs temperature --nAttrs temperature --url http://orion:1026/ \
--expires 1day
```

To get historical data, specify a past date.

```console
ngsi hget --host quantumleap attrs --id device001 --attrs A1,A2 --hLimit 3 \
--fromDate -5years --toDate -3years
```

<a name="data-option"/>

## --data option

### argument

```console
ngsi create entity --keyValues \
--data ' {
      "id":"urn:ngsi-ld:Product:001",
      "type":"Product",
      "name": "Lemonade",
      "size": "S",
      "price": 99
}'
```

### pipe

```console
echo "{ \"id\":\"urn:ngsi-ld:Product:003\", \"type\":\"Product\", \"name\": \"Lemonade\", \"size\": \"S\", \"price\": 99 }" | ngsi create entity --keyValues --data @-
```

```console
echo "{ \"id\":\"urn:ngsi-ld:Product:003\", \"type\":\"Product\", \"name\": \"Lemonade\", \"size\": \"S\", \"price\": 99 }" | ngsi create entity --keyValues --data stdin
```

```console
echo '{ "id":"urn:ngsi-ld:Product:002", "type":"Product", "name": "Lemonade", "size": "S", "price": 99 }' | ngsi create entity --keyValues --data @-
```

### file

```console
ngsi create entity --keyValues --data @data.json
```

data.json:

```json
{
  "id":"urn:ngsi-ld:Product:001",
  "type":"Product",
  "name": "Lemonade",
  "size": "S",
  "price": 99
}
```

<a name="safe-string"/>

## Safe string

```console
ngsi broker get -host orion
```

```json
{"brokerHost":"http://localhost:1026","ngsiType":"v2","safeString":"off"}
```

The value of the `name` attribute has a forbidden characters.

```console
ngsi create entity --keyValues \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "<Lemonade>",
      "size": "S",
      "price": 99
}'
```

```text
entityCreate006 400 Bad Request {"error":"BadRequest","description":"Invalid characters in attribute value"}
```

Create entity with `--safeString on`

```console
ngsi create entity --keyValues --safeString on \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "<Lemonade>",
      "size": "S",
      "price": 99
}'
```

Get an attribute value with `--safeString off`

```console
ngsi get attr --id urn:ngsi-ld:Product:110 --attrName name --safeString off
```

```json
"%3CLemonade%3E"
```

Get an attribute value with `--safeString on`

```console
ngsi get attr --id urn:ngsi-ld:Product:110 --attrName name --safeString on
```

```json
"<Lemonade>"
```

<a name="error-message"/>

## Error message

An error message consists of a prefix and a body. E.g.

```text
entityCreate006 400 Bad Request {"error":"BadReqest","description":"Invalid characters in attribute value"}
```

The error message has `entityCreate006` as a prefix. A prefix consists of a Go function name and a position in the funciton.
The function name is `entityCreate`. The position is 6th.

<a name="detailed-error-information"/>

### Detailed error information

You can get a detailed error information by running a command with the `--stderr info` option.

```console
ngsi --stderr info version --host http://192.168.11.0
```

```text
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
