# Settings - Management command

-   [List current settings](#list-current-settings)
-   [Delete currnet settings](#delete-currnet-settings)
-   [Clear currnet settings](#clear-currnet-settings)
-   [Set previous args mode](#set-previousargs-mode)

<a name="list-current-settings"></a>

## List current settings

This command allows you to print currnet previous args. If you get confused about a behavior of NGSI Go,
please run this command to check previous values. You can clear previous values with `ngsi settings clear`.

```console
ngsi settings list [options]
```

### Options

| Options | Description                |
| ------- | -------------------------- |
| --all   | ail itmes (default: false) |
| --help  | show help (default: true)  |

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

<a name="delete-currnet-settings"></a>

## Delete currnet settings

```console
ngsi settings list [options]
```

### Options

| Options                 | Description               |
| ----------------------- | ------------------------- |
| --items VALUE, -i VALUE | itmes                     |
| --help                  | show help (default: true) |

#### Example

```console
ngsi settings delete --items service,syslog
```

<a name="clear-currnet-settings"></a>

## Clear currnet settings

This command allows you to clear currnet previous args. If you get confused about a behavior of NGSI Go,
please run `ngsi settings list` command to check previous values. You can clear previous values with this
command.

```console
ngsi settings clear [options]
```

### Options

| Options | Description               |
| ------- | ------------------------- |
| --help  | show help (default: true) |

<a name="set-previousargs-mode"></a>

## Set previous args mode

This command allows you to turns on or off the previous args mode. When multiple NGSI Go are run at the same time,
there is a conflict in writing to the config file. E.g. you run NGSI Go on an interactive shell and run one in
background as a batch process at same time. As a result, you may access an unexpected host.

The previous agrs are not stored in the config file when the PreviousArgs mode is off or you run NGSI Go with
`-B` (`--batch`) option. I recommend to use -Boption and specify explicitly a host, FIWARE Service and FIWARE
ServicePath when NGSI Go is used in a shell script.

```
ngsi settings previousArgs [options]
```

### Options

| Options   | Description                    |
| --------- | ------------------------------ |
| --off, -d | off (disable) (default: false) |
| --on, -e  | on (enable) (default: false)   |
| --help    | show help (default: true)      |
