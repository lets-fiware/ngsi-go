# admin scorpio - Convenience command

This command allows you to manage information about Scorpio.

-   [List information paths](#list-information-paths)
-   [Print types](#print-types)
-   [Print local types](#print-local-types)
-   [Print stats](#print-stats)
-   [Print health](#print-health)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-scorpio"></a>

## List information paths

This command lists information paths.

```console
ngsi admin [command options] scorpio list [options]
```

### Options

| Options         | Description                                                                   |
| --------------- | ----------------------------------------------------------------------------- |
| --help          | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host broker scorpio list
```

```console
available subresources:
/types
/localtypes
/stats
/health
```

<a name="print-types"></a>

## Print types

This command prints types.

```console
ngsi admin [command options] scorpio types [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host broker scorpio types
```

<a name="print-local-types"></a>

## Print local types

This command prints local types.

```console
ngsi admin [command options] scorpio localtypes [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host broker scorpio localtypes
```

<a name="print-stats"></a>

## Print stats

This command prints stats.

```console
ngsi admin [command options] scorpio stats [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host broker scorpio stats
```

```json
{
  "number of local available types": 0,
  "number of local available entities": 0
}
```

<a name="print-health"></a>

## Print health

This command prints health.

```console
ngsi admin [command options] scorpio health [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host scorpio@n5000 scorpio health
```

```json
{
  "Status of Registrymanager": "Up and running",
  "Status of Entitymanager": "Up and running",
  "Status of Subscriptionmanager": "Not running",
  "Status of Storagemanager": "Up and running",
  "Status of Querymanager": "Up and running",
  "Status of Historymanager": "Up and running"
}
```
