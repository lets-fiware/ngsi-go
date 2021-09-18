# admin - Convenience command

This command gives you various system information for FIWARE Orion, Cygnus, Perseo. You can use these information for troubleshooting.

-   [Log](#log)
-   [Trace](#trace)
-   [Semaphore](#semaphore)
-   [Metrics](#metrics)
-   [Statistics](#statistics)
-   [Cache statistics](#cache-statistics)
-   Sub commands
    -   [appenders](appenders.md)
    -   [loggers](loggers.md)
    -   [scorpio](scorpio.md)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="log"></a>

## log

This command allows you to print or set logging level for FIWARE Orion.

```console
ngsi admin [command options] log [options]
```

### Options

| Options                 | Description                            |
| ----------------------- | -------------------------------------- |
| --host VALUE, -h VALUE  | broker or server host VALUE (required) |
| --level VALUE, -l VALUE | log level                              |
| --logging, -L           | logging (default: false)               |
| --pretty, -P            | pretty format (default: false)         |
| --help                  | show help (default: true)              |

### Example

#### Request:

```console
ngsi admin log
```

```json
{"level":"DEBUG"}
```

#### Request:

```console
ngsi admin log --level info
```

<a name="trace"></a>

## Trace

This command allows you to print, set or delete trace level for FIWARE Orion.

```console
ngsi admin [common options] trace [options]
```

### Options

| Options                 | Description                            |
| ----------------------- | -------------------------------------- |
| --host VALUE, -h VALUE  | broker or server host VALUE (required) |
| --level VALUE, -l VALUE | log level                              |
| --set, -s               | set (default: false)                   |
| --delete, -d            | delete (default: false)                |
| --logging, -L           | logging (default: false)               |
| --help                  | show help (default: true)              |

### Example

#### Request:

```console
ngsi admin trace
```

```json
{"tracelevels":"empty"}
```

#### Request:

```console
ngsi admin trace --level t1
ngsi admin trace --level t1-t2
ngsi admin trace --level t1-t2,t3-t4
ngsi admin trace --level 180-199
```

#### Request:

```console
ngsi admin trace --delete
ngsi admin trace --delete --level t1
ngsi admin trace --delete --level t1-t2
ngsi admin trace --delete --level t1-t2,t3-t4
```

<a name="semaphore"></a>

## Semaphore

This command prints semaphore for FIWARE Orion.

```console
ngsi admin [command options] semaphore [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --logging, -L          | logging (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

#### Request:

```console
ngsi admin semaphore --pretty
```

```json
{
  "dbConnectionPool": {
    "status": "free"
  },
  "dbConnection": {
    "status": "free"
  },
  "request": {
    "status": "free"
  },
  "subCache": {
    "status": "free"
  },
  "transaction": {
    "status": "free"
  },
  "timeStat": {
    "status": "free"
  },
  "logMsg": {
    "status": "free"
  },
  "alarmMgr": {
    "status": "free"
  },
  "metricsMgr": {
    "status": "free"
  },
  "connectionContext": {
    "status": "free"
  },
  "connectionEndpoints": {
    "status": "free"
  }
}
```

<a name="metrics"></a>

## Metrics

This command allows you to print, reset or delete metrics for FIWARE Orion, Cygnus.

```console
ngsi admin [command options] metrics [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --delete, -d           | delete (default: false)                |
| --reset, -r            | reset (default: false)                 |
| --logging, -L          | logging (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

#### Request:

```console
ngsi admin metrics --pretty
```

```json
{
  "services": {
    "default-service": {
      "subservs": {
        "root-subserv": {
          "incomingTransactionResponseSize": 1734,
          "serviceTime": 0.000461688,
          "incomingTransactions": 16
        }
      },
      "sum": {
        "incomingTransactionResponseSize": 1734,
        "serviceTime": 0.000461688,
        "incomingTransactions": 16
      }
    }
  },
  "sum": {
    "subservs": {
      "root-subserv": {
        "incomingTransactionResponseSize": 1734,
        "serviceTime": 0.000461688,
        "incomingTransactions": 16
      }
    },
    "sum": {
      "incomingTransactionResponseSize": 1734,
      "serviceTime": 0.000461688,
      "incomingTransactions": 16
    }
  }
}
```

#### Request:

```console
ngsi admin metrics --reset --pretty
```

```json
{
  "services": {
    "default-service": {
      "subservs": {
        "root-subserv": {
          "incomingTransactionResponseSize": 482,
          "serviceTime": 0.000316,
          "incomingTransactions": 1
        }
      },
      "sum": {
        "incomingTransactionResponseSize": 482,
        "serviceTime": 0.000316,
        "incomingTransactions": 1
      }
    }
  },
  "sum": {
    "subservs": {
      "root-subserv": {
        "incomingTransactionResponseSize": 482,
        "serviceTime": 0.000316,
        "incomingTransactions": 1
      }
    },
    "sum": {
      "incomingTransactionResponseSize": 482,
      "serviceTime": 0.000316,
      "incomingTransactions": 1
    }
  }
}
```

#### Request:

```console
ngsi admin metrics --delete
```

<a name="statistics"></a>

## Statistics

This command allows you to print or delete statistics for FIWARE Orion, Cygnus.

```console
ngsi admin [command options] statistics [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --delete, -d           | delete (default: false)                |
| --logging, -L          | logging (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

#### Request:

```console
 ngsi admin statistics --pretty
```

```json
{
  "uptime_in_secs": 156334,
  "measuring_interval_in_secs": 156334
}
```

#### Request:

```console
ngsi admin statistics --delete
```

<a name="cache-statistics"></a>

## Cache statistics

This command allows you to print or delete cache statistics for FIWARE Orion.

```console
ngsi admin [command options] cacheStatistics [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --delete, -d           | delete (default: false)                |
| --logging, -L          | logging (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

#### Request:

```console
ngsi admin cacheStatistics --pretty
```

```json
{
  "ids": "",
  "refresh": 2011,
  "inserts": 0,
  "removes": 0,
  "updates": 0,
  "items": 0
}
```

#### Request:

```console
ngsi admin cacheStatistics --delete
```
