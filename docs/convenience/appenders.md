# admin appenders - Convenience command

This command allows you to manage appenders for Cygnus.

-   [List appenders](#list-appenders)
-   [Get a appender](#get-a-appender)
-   [Create a appender](#create-a-appender)
-   [Update a appender](#update-a-appender)
-   [Delete a appender](#delete-a-appender)

## Common Options


| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-appenders"></a>

## List appenders

This command lists appenders for Cygnus

```console
ngsi admin [command options] appenders list [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required)                                        |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: true)                                                     |

### Example

#### Request:

```console
ngsi admin --host cygnus appenders list --pretty
```

```json
{
  "success": "true",
  "appenders": [
    {
      "name": "DAILY",
      "layout": "time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n",
      "active": "false"
    },
    {
      "name": "LOGFILE",
      "layout": "time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n",
      "active": "true"
    },
    {
      "name": "console",
      "layout": "time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n",
      "active": "false"
    }
  ]
}
```

<a name="get-a-appender"></a>

## Get a appender

This command gets a appender for Cygnus

```console
ngsi admin [command options] appenders get [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required)                                        |
| --name VALUE, -n VALUE | appender name (required)                                                      |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: true)                                                     |

### Example

#### Request:

```console
ngsi admin --host cygnus appenders get --name console
```

```json
{"success":"true","appender":"[{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n","active":"false"}"}

```

<a name="create-a-appender"></a>

## Create a appender

This command creates a appender for Cygnus

```console
ngsi admin [command options] appenders create [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required)                                        |
| --name VALUE, -n VALUE | appender name                                                                 |
| --data VALUE, -d VALUE | appender information (required)                                               |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: true)                                                     |

### Example

#### Request:

```console
ngsi admin --host cygnus appenders create --name test --data \
'{
    "appender": {
        "name":"test",
        "class":""
    },
    "pattern": {
        "layout":"",
        "ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n"
    }
}'
```

```json
{"success":"true","result":"Appender 'test' posted"}
```

<a name="update-a-appender"></a>

## Update a appender

This command updates a appender for Cygnus

```console
ngsi admin [command options] appenders update [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required)                                        |
| --name VALUE, -n VALUE | appender name (required)                                                      |
| --data VALUE, -d VALUE | appender information (required)                                               |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: true)                                                     |

### Example

#### Request:

```console
ngsi admin --host cygnus appenders update --name test --data \
'{
    "appender": {
        "name":"test",
        "class":""
    },
    "pattern": {
        "layout":"",
        "ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z | lvl=%p | corr=%X{correlatorId} | trans=%X{transactionId} | srv=%X{service} | subsrv=%X{subservice} | comp=%X{agent} | op=%M | msg=%C[%L] : %m%n"
    }
}'
```

```json
{"success":"true","result":"Appender 'test' put"}
```

<a name="delete-a-appender"></a>

## Delete a appender

This command deletes a appender for Cygnus

```console
ngsi admin [command options] appenders delete [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required)                                        |
| --name VALUE, -n VALUE | appender name (required)                                                      |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: true)                                                     |

### Example

#### Request:

```console
ngsi admin --host cygnus appenders delete --name test 
```

```json
{"success":"true","result":" Appender 'test' removed successfully"}
```
