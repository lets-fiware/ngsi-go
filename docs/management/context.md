# Context - Management command

-   [List contexts](#list-contexts)
-   [Add context](#add-context)
-   [Update context](#update-context)
-   [Delete context](#delete-context)

## List contexts

```console
ngsi context list [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --name value, -n value          | specify @context name                        |
| --help                          | show help (default: false)                   |

#### Example 1

```console
ngsi context list
```

```text
etsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
ld https://schema.lab.fiware.org/ld/context
tutorial http://context-provider:3000/data-models/ngsi-context.jsonld
data-model http://context-provider:3000/data-models/ngsi-context.jsonld
```

#### Example 2

```console
ngsi context list --name etsi
```

```text
https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
```

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

#### Example

```console
ngsi context add --name tutorial --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

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

#### Example

```console
ngsi context update --name data-model --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

## Delete context

```console
ngsi context delete [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --help                 | show help (default: false)          |

#### Example

```console
ngsi context delete --name data-model
```
