# Settings - Management command

-   [List current settings](#list-current-settings)
-   [Delete currnet settings](#delete-currnet-settings)
-   [Clear currnet settings](#clear-currnet-settings)

## List current settings

```
ngsi settings list [options]
```

### Options

| Options | Description                      |
| ------- | -------------------------------- |
| --all   | print ail itmes (default: false) |                         
| --help  | show help (default: false)       |

#### Example 1

```
$ ngsi settings list
Host: orion
FIWARE-Service: openiot
Syslog: debug
```

#### Example 2

```
$ ngsi settings list --all
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

```
ngsi settings list [options]
```

### Options

| Options                 | Description                                 |
| ----------------------- | ------------------------------------------- |
| --items value, -i value | specify the items in a comma-separated list |
| --help                  | show help (default: false)                  |

#### Example

```
$ ngsi settings delete --items service,syslog
```

## Clear currnet settings

```
ngsi settings clear [options]
```

### Options

| Options                 | Description                                 |
| ----------------------- | ------------------------------------------- |
| --help                  | show help (default: false)                  |

