# context - Management command

-   [List contexts](#list-contexts)
-   [Add context](#add-context)
-   [Update context](#update-context)
-   [Delete context](#delete-context)

## List contexts

```
ngsi context list [options]
```

### Options

| Options                         | Description                                  |
| ------------------------------- | -------------------------------------------- |
| --name value, -n value          | specify @context name                        |
| --help                          | show help (default: false)                   |

#### Example 1

```
$ ngsi context list
etsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
ld https://schema.lab.fiware.org/ld/context
tutorial http://context-provider:3000/data-models/ngsi-context.jsonld
data-model http://context-provider:3000/data-models/ngsi-context.jsonld
```

#### Example 2

```
$ ngsi context list --name etsi
https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
```

## Add context

```
ngsi context add [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --url value, -u value  | specify URL for @context (Required) |
| --help                 | show help (default: false)          |

#### Example

```
$ ngsi context add --name tutorial --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

## Update context

```
ngsi context update [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --url value, -u value  | specify URL for @context (Required) |
| --help                 | show help (default: false)          |

#### Example

```
$ ngsi context update --name data-model --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

## Delete context

```
ngsi context delete [options]
```

### Options

| Options                | Description                         |
| ---------------------- | ----------------------------------- |
| --name value, -n value | specify @context name (Required)    |
| --help                 | show help (default: false)          |

#### Example

```
$ ngsi context delete --name data-model
```
