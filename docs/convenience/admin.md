# admin - Convenience command

This command gives you various system information for FIWARE Orion. You can use these information for troubleshooting.

-   [Log](#log)
-   [Trace](#trace)
-   [Semaphore](#semaphore)
-   [Metrics](#metrics)
-   [Statistics](#statistics)
-   [Cache statistics](#cache-statistics)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="log"/>

## log

This command allows you to print or set logging level for FIWARE Orion.

```console
ngsi admin [command options] log [options]
```

### Options

| Options                 | Description                                                         |
| ----------------------- | ------------------------------------------------------------------- |
| --level value, -l value | specify log level (none, fatal, error, warn, info, debug)           | 
| --logging, -L           | logging output when logging level higher than Info (default: false) |
| --pretty, -P            | pretty format (default: false)                                      |
| --help                  | show help (default: false)                                          |

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

<a name="trace"/>

## Trace

This command allows you to print, set or delete trace level for FIWARE Orion.

```console
ngsi admin [common options] trace [options]
```

### Options

| Options                 | Description                                                         |
| ----------------------- | ------------------------------------------------------------------- |
| --level value, -l value | specify log level                                                   |
| --set, -s               | set (default: false)                                                |
| --delete, -d            | delete (default: false)                                             |
| --logging, -L           | logging output when logging level higher than Info (default: false) |
| --help                  | show help (default: false)                                          |

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

<a name="semaphore"/>

## Semaphore

This command prints semaphore for FIWARE Orion.

```console
ngsi admin [command options] semaphore [options]
```

### Options

| Options       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| --logging, -L | logging output when logging level higher than Info (default: false) |
| --pretty, -P  | pretty format (default: false)                                      |
| --help        | show help (default: false)                                          |

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

<a name="metrics"/>

## Metrics

This command allows you to print, reset or delete metrics for FIWARE Orion.

```console
ngsi admin [command options] metrics [options]
```

### Options

| Options       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| --delete, -d  | delete metrics (default: false)                                     |
| --reset, -r   | reset metrics (default: false)                                      |
| --logging, -L | logging output when logging level higher than Info (default: false) |
| --pretty, -P  | pretty format (default: false)                                      |
| --help        | show help (default: false)                                          |

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

<a name="statistics"/>

## Statistics

This command allows you to print or delete statistics for FIWARE Orion.

```console
ngsi admin [command options] statistics [options]
```

### Options

| Options       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| --delete, -d  | delete metrics (default: false)                                     |
| --logging, -L | logging output when logging level higher than Info (default: false) |
| --pretty, -P  | pretty format (default: false)                                      |
| --help        | show help (default: false)                                          |

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

<a name="cache-statistics"/>

## Cache statistics

This command allows you to print or delete cache statistics for FIWARE Orion.

```console
ngsi admin [command options] cacheStatistics [options]
```

### Options

| Options       | Description                                                         |
| ------------- | ------------------------------------------------------------------- |
| --delete, -d  | delete metrics (default: false)                                     |
| --logging, -L | logging output when logging level higher than Info (default: false) |
| --pretty, -P  | pretty format (default: false)                                      |
| --help        | show help (default: false)                                          |

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
