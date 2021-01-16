# Settings - Management command

-   [List current settings](#list-current-settings)
-   [Delete currnet settings](#delete-currnet-settings)
-   [Clear currnet settings](#clear-currnet-settings)

## List current settings

```console
ngsi settings list [options]
```

### Options

| Options | Description                      |
| ------- | -------------------------------- |
| --all   | print ail itmes (default: false) |
| --help  | show help (default: false)       |

#### Example 1

```console
ngsi settings list
```

```text
Host: orion
FIWARE-Service: openiot
Syslog: debug
```

#### Example 2

```console
ngsi settings list --all
```

```text
Host: orion
FIWARE-Service:
FIWARE-ServicePath:
Token:
Syslog:
Stderr:
LogFile:
LogLevel:
```

## Delete currnet settings

```console
ngsi settings list [options]
```

### Options

| Options                 | Description                                 |
| ----------------------- | ------------------------------------------- |
| --items value, -i value | specify the items in a comma-separated list |
| --help                  | show help (default: false)                  |

#### Example

```console
ngsi settings delete --items service,syslog
```

## Clear currnet settings

```console
ngsi settings clear [options]
```

### Options

| Options                 | Description                                 |
| ----------------------- | ------------------------------------------- |
| --help                  | show help (default: false)                  |
