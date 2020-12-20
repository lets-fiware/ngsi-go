# Context - Management command

-   [List contexts](#list-contexts)
-   [Add context](#add-context)
-   [Update context](#update-context)
-   [Delete context](#delete-context)
-   [Serve context](#serve-context)

<a name="list-contexts">

## List contexts

```console
ngsi context list [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --name value, -n value          | specify @context name                        |
| --help                          | show help (default: false)                   |

### Example 1

```console
ngsi context list
```

```text
etsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
ld https://schema.lab.fiware.org/ld/context
tutorial http://context-provider:3000/data-models/ngsi-context.jsonld
data-model http://context-provider:3000/data-models/ngsi-context.jsonld
```

### Example 2

```console
ngsi context list --name etsi
```

```text
https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
```

<a name="add-context">

## Add context

```console
ngsi context add [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --url value, -u value  | specify URL for @context (Required) |
| --help                 | show help (default: false)          |

### Example

```console
ngsi context add --name tutorial --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

<a name="update-context">

## Update context

```console
ngsi context update [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --url value, -u value  | specify URL for @context (Required) |
| --help                 | show help (default: false)          |

### Example

```console
ngsi context update --name data-model --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

<a name="delete-context">

## Delete context

```console
ngsi context delete [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --help                 | show help (default: false)          |

### Example

```console
ngsi context delete --name data-model
```

<a name="serve-context">

## Serve context

```console
ngsi context server [options]
```

### Options

| Options                | Description                                         |
| ---------------------- | --------------------------------------------------- |
| --name value, -n value | specify @context name                               |
| --data value, -d value | specify @context data                               |
| --host value, -h value | specify host for receiver (default: "0.0.0.0")      |
| --port value, -p value | specify port for receiver (default: "1028")         |
| --url value, -u value  | specify url for receiver (default: "/")             |
| --https, -s            | start in https (default: false)                     |
| --key value, -k value  | specify key file (only needed if https is enabled)  |
| --cert value, -c value | specify cert file (only needed if https is enabled) |
| --help                 | specify show help (default: false)                  |

### Example

```console
ngsi context server --name ld
```
