# hdelete - NGSI command

-   [Delete all the data associated to certain attribute of certain entity](#delete-all-the-data-associated-to-certain-attribute-of-certain-entity)
-   [Delete historical data of a certain entity](#delete-historical-data-of-a-certain-entity)
-   [Delete historical data of all entities of a certain type](#delete-historical-data-of-all-entities-of-a-certain-type)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="delete-all-the-data-associated-to-certain-attribute-of-certain-entity"></a>

## Delete all the data associated to certain attribute of certain entity

This command deletes all the data associated to certain attribute of certain entity.

```console
ngsi hdelete [common options] attr [options]
```

### Options

| Options                | Description                  |
| ---------------------- | ---------------------------- |
| --id value, -i value   | entity id                    |
| --type value, -t value | entity Type                  |
| --attrName value       | attribute name               |
| --run                  | run command (default: false) |
| --help                 | show help (default: false)   |

### Example

#### Request:

```console
ngsi hdelete --host comet attr --type device --id device001 --attrName A1
```
<a name="delete-historical-data-of-a-certain-entity"></a>

## Delete historical data of a certain entity

This command deletes historical data of a certain entity.

```console
ngsi hdelete [common options] entity [options]
```

### Options

| Options                | Description                                       |
| ---------------------- | ------------------------------------------------- |
| --id value, -i value   | Entity id                                         |
| --type value, -t value | Entity Type                                       |
| --fromDate value       | starting date from which data should be retrieved |
| --toDate value         | final date until which data should be retrieved   |
| --run                  | run command (default: false)                      |
| --help                 | show help (default: false)                        |

### Example

#### Request:

```console
ngsi hdelete --host comet entity --type device --id device001
```

#### Request:

```console
ngsi hdelete --host quantumleap entity --id device003
```

<a name="delete-historical-data-of-all-entities-of-a-certain-type"></a>

## Delete historical data of all entities of a certain type

This command deletes historical data of all entities of a certain type.

```console
ngsi hdelete [common options] entities [options]
```

### Options

| Options                | Description                                       |
| ---------------------- | ------------------------------------------------- |
| --id value, -i value   | Entity id                                         |
| --type value, -t value | Entity Type                                       |
| --dropTable value      | drop the table storing an entity type             |
| --fromDate value       | starting date from which data should be retrieved |
| --toDate value         | final date until which data should be retrieved   |
| --run                  | run command (default: false)                      |
| --help                 | show help (default: false)                        |

### Example

#### Request:

```console
ngsi hdelete --host comet entities
```

#### Request:

```console
ngsi hdelete --host quantumleap entities --type Thing
```
