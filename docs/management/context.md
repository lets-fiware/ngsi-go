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

| Options                | Description               |
| ---------------------- | ------------------------- |
| --name VALUE, -n VALUE | @context name             |
| --help                 | show help (default: true) |

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

| Options                | Description               |
| ---------------------- | ------------------------- |
| --name VALUE, -n VALUE | @context name (required)  |
| --url VALUE, -u VALUE  | url for @context          |
| --json VALUE, -j VALUE | url for @context          |
| --help                 | show help (default: true) |

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

| Options                | Description                 |
| ---------------------- | --------------------------- |
| --name VALUE, -n VALUE | @context name (required)    |
| --url VALUE, -u VALUE  | url for @context (required) |
| --help                 | show help (default: true)   |

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

| Options                | Description               |
| ---------------------- | ------------------------- |
| --name VALUE, -n VALUE | @context name (required)  |
| --help                 | show help (default: true) |

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

| Options                | Description                                 |
| ---------------------- | ------------------------------------------- |
| --name VALUE, -n VALUE | @context name                               |
| --data VALUE, -d VALUE | @context data                               |
| --host VALUE, -h VALUE | host for server                             |
| --port VALUE, -p VALUE | port for server                             |
| --url VALUE, -u VALUE  | url for server                              |
| --https, -s            | start in https (default: false)             |
| --key VALUE, -k VALUE  | key file (only needed if https is enabled)  |
| --cert VALUE, -c VALUE | cert file (only needed if https is enabled) |
| --help                 | show help (default: true)                   |

### Example

```console
ngsi context server --name ld
```
