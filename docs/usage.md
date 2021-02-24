# Usage

-   [Syntax](#syntax)
    -   [Convenience command](#convenience-command)
    -   [NGSI Command](#ngsi-command)
    -   [Time series command](#time-series-command)
    -   [Cygnus command](#cygnus-command)
    -   [IoT Agent command](#iot-agent-command)
    -   [Perseo command](#context-aware-cep-command)
    -   [Keyrock](#keyrock-command)
    -   [Management commnad](#management-commnad)
    -   [Global Options](#global-options)
    -   [Common options](#common-options)
-   [DateTime options](#datetime-options)
-   [--data option](#data-option)
-   [Safe string](#safe-string)
-   [Error message](#error-message)
    -   [Detailed error information](#detailed-error-information)

<a name="syntax"></a>

## Syntax

```console
ngsi [global options] command [common options] sub-command [options]
```

<a name="convenience-command"></a>

### Convenience command

| command                               | sub-command                                                | sub-sub-commnand                                         | Description                                                      |
| ------------------------------------- | ---------------------------------------------------------- | -------------------------------------------------------- | ---------------------------------------------------------------- |
| [admin](./convenience/admin.md)       | [log](./convenience/admin.md#log)                          | -                                                        | print or set logging level for FIWARE Orion                      |
|                                       | [trace](./convenience/admin.md#trace)                      | -                                                        | print, set or delete trace level for FIWARE Orion                | 
|                                       | [semaphore](./convenience/admin.md#semaphore)              | -                                                        | print semaphore for FIWARE Orion                                 |
|                                       | [metrics](./convenience/admin.md#metrics)                  | -                                                        | print, reset or delete metrics for FIWARE Orion, Cygnus          |
|                                       | [statistics](./convenience/admin.md#statistics)            | -                                                        | print or delete statistics for FIWARE Orion, Cygnus              |
|                                       | [cacheStatistics](./convenience/admin.md#cache-statistics) | -                                                        | print or delete cache statistics for FIWARE Orion                |
|                                       | [appenders](./convenience/appenders.md)                    | [list](./convenience/appenders.md#list-appenders)        | list appenders                                                   |
|                                       |                                                            | [get](./convenience/appenders.md#get-a-appender)         | get a appender                                                   |
|                                       |                                                            | [create](./convenience/appenders.md#create-a-appender)   | create a appender                                                |
|                                       |                                                            | [upadte](./convenience/appenders.md#update-a-appender)   | update a appender                                                |
|                                       |                                                            | [delete](./convenience/appenders.md#delete-a-appender)   | delete a appender                                                |
|                                       | [loggers](./convenience/loggers.md)                        | [list](./convenience/loggers.md#list-loggers)            | List loggers                                                     |
|                                       |                                                            | [get](./convenience/loggers.md#get-a-logger)             | get a logger                                                     |
|                                       |                                                            | [create](./convenience/loggers.md#create-a-logger)       | create a logger                                                  |
|                                       |                                                            | [update](./convenience/loggers.md#update-a-logger)       | updata a logger                                                  |
|                                       |                                                            | [delete](./convenience/loggers.md#delete-a-logger)       | delete a logger                                                  |
|                                       | [scorpio](./convenience/scorpio.md)                        | [list](./convenience/scorpio.md#list-information-paths)  | List information paths                                           |
|                                       |                                                            | [types](./convenience/scorpio.md#print-types)            | Print types                                                      |
|                                       |                                                            | [localtypes](./convenience/scorpio.md#print-local-types) | Print local types                                                |
|                                       |                                                            | [stats](./convenience/scorpio.md#print-stats)            | Print stats                                                      |
|                                       |                                                            | [health](./convenience/scorpio.md#print-health)          | Print health                                                     |
| [apis](./convenience/apis.md)         | -                                                          | -                                                        | print endpoints of FWARE Open APIs                               |
| [cp](./convenience/cp.md)             | -                                                          | -                                                        | copy entities                                                    |
| [wc](./convenience/wc.md)             | -                                                          | -                                                        | print number of entities, subscriptions, registrations, or types |
| [man](./convenience/man.md)           | -                                                          |                                                          | print urls of document                                           |
| [health](./convenience/health.md)     | -                                                          |                                                          | print health status of FIWARE GEs                                |
| [ls](./convenience/ls.md)             | -                                                          |                                                          | list entities                                                    |
| [rm](./convenience/rm.md)             | -                                                          |                                                          | remove entities                                                  |
| [receiver](./convenience/receiver.md) | -                                                          |                                                          | notification receiver                                            |
| [template](./convenience/template.md) | [subscription](./convenience/template.md#subscription)     |                                                          | create template of subscription                                  |
|                                       | [registration](./convenience/template.md#registration)     |                                                          | create template of registration                                  |
| [version](./convenience/version.md)   | -                                                          |                                                          | print the version of Context Broker                              |

<a name="ngsi-command"></a>

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

<a name="time-series-command"></a>

### Time series command

| command                             | sub-command | Description                                                           |
| ----------------------------------- | ----------- | --------------------------------------------------------------------- |
| [hdelete](./time_series/hdelete.md) | attr        | delete all the data associated to certain attribute of certain entity |
|                                     | entity      | delete historical data of a certain entity                            |
|                                     | entities    | delete historical data of all entities of a certain type              |
| [hget](./time_series/hget.md)       | attr        | get hstory of an attribute                                            |
|                                     | attrs       | get history of attributes                                             |
|                                     | entities    | list of all the entity id                                             |

<a name="cygnus-command"></a>

### Cygnus command

| command                                    | sub-command                                               | Description            |
| ------------------------------------------ | --------------------------------------------------------- | ---------------------- |
| [groupingrules](./cygnus/groupingrules.md) | [list](./cygnus/groupingrules.md#list-groupingrules)      | List grouping rules    |
|                                            | [create](./cygnus/groupingrules.md#create-a-groupingrule) | Create a grouping rule |
|                                            | [update](./cygnus/groupingrules.md#update-a-groupingrule) | Update a grouping rule |
|                                            | [delete](./cygnus/groupingrules.md#delete-a-groupingrule) | Delete a grouping rule |
| [namemappings](./cygnus/namemappings.md)   | [list](./cygnus/namemappings.md#list-namemappings)        | List name mappings     |
|                                            | [create](./cygnus/namemappings.md#create-a-namemapping)   | Create a name mapping  |
|                                            | [delete](./cygnus/namemappings.md#delete-a-namemapping)   | Delete a name mapping  |

<a name="iot-agent-command"></a>

### IoT Agent command

| command                           | sub-command                                                  | Description                  |
| --------------------------------- | ------------------------------------------------------------ | ---------------------------- |
| [services](iot_agent/services.md) | [list](iot_agent/services.md#list-configuration-group)       | List configuration groups    |
|                                   | [create](iot_agent/services.md#create-a-configuration-group) | Create a configuration group |
|                                   | [update](iot_agent/services.md#update-a-configuration-group) | Update a configuration group |
|                                   | [delete](iot_agent/services.md#delete-a-configuration-group) | Delete a configuration group |
| [devices](iot_agent/devices.md)   | [list](iot_agent/devices.md#list-all-devices)                | List all devices             |
|                                   | [create](iot_agent/devices.md#create-a-device)               | Create a device              |
|                                   | [get](iot_agent/devices.md#create-a-get-device)              | Get a device                 |
|                                   | [update](iot_agent/devices.md#update-a-device)               | Update a device              |
|                                   | [delete](iot_agent/devices.md#delete-a-device)               | Delete a device              |

<a name="context-aware-cep-command"></a>

### Context-Aware CEP command

| command               | sub-command                           | Description    |
| --------------------- | ------------------------------------- | -------------- |
| [rules](cep/rules.md) | [list](cep/rules.md#list-all-rules)   | List all rules |
|                       | [create](cep/rules.md#create-a-rule)  | Creates a rule |
|                       | [get](cep/rules.md#create-a-get-rule) | Get a rule     |
|                       | [delete](cep/rules.md#delete-a-rule)  | Delete a rule  |

<a name="keyrock-command"></a>

### Keyrock

| command                                   | sub-command                                                             | sub-sub-command                                                                                             | Description                                                                                   |
| ----------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| [users](keyrock/users.md)                 | [list](keyrock/users.md#list-users)                                     | -                                                                                                           | List users                                                                                    |
|                                           | [get](keyrock/users.md#get-a-user)                                      | -                                                                                                           | Get a user                                                                                    |
|                                           | [create](keyrock/users.md#create-a-user)                                | -                                                                                                           | Create a user                                                                                 |
|                                           | [update](keyrock/users.md#update-a-user)                                | -                                                                                                           | Update a user                                                                                 |
|                                           | [delete](keyrock/users.md#delete-a-user)                                | -                                                                                                           | Delete a user                                                                                 |
| [organizations](keyrock/organizations.md) | [list](keyrock/organizations.md#list-organizations)                     | -                                                                                                           | List organizations                                                                            |
|                                           | [get](keyrock/organizations.md#get-an-organization)                     | -                                                                                                           | Get an organization                                                                           |
|                                           | [create](keyrock/organizations.md#create-an-organization)               | -                                                                                                           | Create an organization                                                                        |
|                                           | [update](keyrock/organizations.md#update-an-organization)               | -                                                                                                           | Update an organization                                                                        |
|                                           | [delete](keyrock/organizations.md#delete-an-organization)               | -                                                                                                           | Delete an organization                                                                        |
|                                           | [users](keyrock/organizations-users.md)                                 | [list](keyrock/organizations-users.md#list-users-of-an-organization)                                        | List users of an organization                                                                 |
|                                           |                                                                         | [get](keyrock/organizations-users.md#get-info-of-user-organization-relationship)                            | Get info of user organization relationship                                                    |
|                                           |                                                                         | [add](keyrock/organizations-users.md#add-a-user-to-an-organization)                                         | Add a user to an organization                                                                 |
|                                           |                                                                         | [remove](keyrock/organizations-users.md#remove-a-user-from-an-organization)                                 | Remove a user from an organization                                                            |
| [applications](keyrock/applications.md)   | [list](keyrock/applications.md#list-applications)                       | -                                                                                                           | List applications                                                                             |
|                                           | [get](keyrock/applications.md#get-an-application)                       | -                                                                                                           | Get an application                                                                            |
|                                           | [create](keyrock/applications.md#create-an-application)                 | -                                                                                                           | Create an application                                                                         |
|                                           | [update](keyrock/applications.md#update-an-get-application)             | -                                                                                                           | Update an application                                                                         |
|                                           | [delete](keyrock/applications.md#delete-an-application)                 | -                                                                                                           | Delete an application                                                                         |
|                                           | [roles](keyrock/applications-roles.md)                                  | [list](keyrock/applications-roles.md#list-roles)                                                            | List roles                                                                                    |
|                                           |                                                                         | [get](keyrock/applications-roles.md#get-a-role)                                                             | Get a role                                                                                    |
|                                           |                                                                         | [create](keyrock/applications-roles.md#create-a-role)                                                       | Create a role                                                                                 |
|                                           |                                                                         | [update](keyrock/applications-roles.md#update-a-role)                                                       | Update a role                                                                                 |
|                                           |                                                                         | [delete](keyrock/applications-roles.md#delete-a-role)                                                       | Delete a role                                                                                 |
|                                           |                                                                         | [permissions](keyrock/applications-roles.md#list-permissions-associated-to-a-role)                          | List permissions associated to a role                                                         |
|                                           |                                                                         | [assign](keyrock/applications-roles.md#assign-a-permission-to-a-role)                                       | Assign a permission to a role                                                                 |
|                                           |                                                                         | [unassign](keyrock/applications-roles.md#delete-a-permission-to-a-role)                                     | Delete a permission from a role                                                               |
|                                           | [permissions](keyrock/applications-permissions.md)                      | [list](keyrock/applications-permissions.md#list-permissions)                                                | List permissions                                                                              |
|                                           |                                                                         | [get](keyrock/applications-permissions.md#get-a-permission)                                                 | Get a permission                                                                              |
|                                           |                                                                         | [create](keyrock/applications-permissions.md#create-a-permission)                                           | Create a permission                                                                           |
|                                           |                                                                         | [update](keyrock/applications-permissions.md#update-a-permission)                                           | Update a permission                                                                           |
|                                           |                                                                         | [delete](keyrock/applications-permissions.md#delete-a-permission)                                           | Delete a permission                                                                           |
|                                           | [organizations](keyrock/applications-organizations.md)                  | [ilst](keyrock/applications-organizations.md#list-organizations-in-an-application)                          | List organizations in an application                                                          |
|                                           |                                                                         | [get](keyrock/applications-organizations.md#get-roles-of-an-organization-in-an-application)                 | Get roles of an organization in an application                                                |
|                                           |                                                                         | [assign](keyrock/applications-organizations.md#assign-a-role-to-an-organization)                            | Assign a role to an organization                                                              |
|                                           |                                                                         | [unassign](keyrock/applications-organizations.md#delete-a-role-assignment-from-an-organization)             | Delete a role assignment from an organization                                                 |
|                                           | [pep-proxies](keyrock/applications-pep-proxies.md)                      | [list](keyrock/applications-pep-proxies.md#list-pep-proxies)                                                | List pep proxies                                                                              |
|                                           |                                                                         | [create](keyrock/applications-pep-proxies.md#create-a-pep-proxy)                                            | Create a pep proxy                                                                            |
|                                           |                                                                         | [reset](keyrock/applications-pep-proxies.md#reset-a-pep-proxy)                                              | Reset a pep proxy                                                                             |
|                                           |                                                                         | [delete](keyrock/applications-pep-proxies.md#delete-a-pep-proxy)                                            | Delete a pep proxy                                                                            |
|                                           | [trusted-applications.md](keyrock/applications-trusted-applications.md) | [list](keyrock/applications-trusted-applications.md#list-trusted-applications-associated-to-an-application) | List trusted applications associated to an application                                        |
|                                           |                                                                         | [add](keyrock/applications-trusted-applications.md#add-a-trusted-application)                               | Add a trusted application                                                                     |
|                                           |                                                                         | [delete](keyrock/applications-trusted-applications.md#delete-a-trusted-application)                         | Delete a trusted application                                                                  |
| [providers](keyrock/providers.md)         | -                                                                       | -                                                                                                           | Print service providers                                                                       |

<a name="management-commnad"></a>

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

<a name="global-options"></a>

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

<a name="common-options"></a>

### Common options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="datetime-options"></a>

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

<a name="data-option"></a>

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

<a name="safe-string"></a>

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

<a name="error-message"></a>

## Error message

An error message consists of a prefix and a body. E.g.

```text
entityCreate006 400 Bad Request {"error":"BadReqest","description":"Invalid characters in attribute value"}
```

The error message has `entityCreate006` as a prefix. A prefix consists of a Go function name and a position in the funciton.
The function name is `entityCreate`. The position is 6th.

<a name="detailed-error-information"></a>

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
