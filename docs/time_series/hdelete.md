# hdelete - NGSI command

-   [Delete all the data associated to certain attribute of certain entity](#delete-all-the-data-associated-to-certain-attribute-of-certain-entity)
-   [Delete historical data of a certain entity](#delete-historical-data-of-a-certain-entity)
-   [Delete historical data of all entities of a certain type](#delete-historical-data-of-all-entities-of-a-certain-type)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="delete-all-the-data-associated-to-certain-attribute-of-certain-entity"></a>

## Delete all the data associated to certain attribute of certain entity

This command deletes all the data associated to certain attribute of certain entity.

```console
ngsi hdelete [common options] attr [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | id                                     |
| --type VALUE, -t VALUE    | Entity Type                            |
| --attr VALUE              | attribute name                         |
| --run                     | run command (default: false)           |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi hdelete --host comet attr --type device --id device001 --attr A1
```
<a name="delete-historical-data-of-a-certain-entity"></a>

## Delete historical data of a certain entity

This command deletes historical data of a certain entity.

```console
ngsi hdelete [common options] entity [options]
```

### Options

| Options                   | Description                                       |
| ------------------------- | ------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)            |
| --service VALUE, -s VALUE | FIWARE Service VALUE                              |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                          |
| --id VALUE, -i VALUE      | id                                                |
| --type VALUE, -t VALUE    | Entity Type                                       |
| --fromDate VALUE          | starting date from which data should be retrieved |
| --toDate VALUE            | final date until which data should be retrieved   |
| --run                     | run command (default: false)                      |
| --help                    | show help (default: true)                         |

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

| Options                   | Description                                            |
| ------------------------- | ------------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                 |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                               |
| --id VALUE, -i VALUE      | id                                                     |
| --type VALUE, -t VALUE    | Entity Type                                            |
| --dropTable               | drop the table storing an entity type (default: false) |
| --fromDate VALUE          | starting date from which data should be retrieved      |
| --toDate VALUE            | final date until which data should be retrieved        |
| --run                     | run command (default: false)                           |
| --help                    | show help (default: true)                              |

### Example

#### Request:

```console
ngsi hdelete --host comet entities
```

#### Request:

```console
ngsi hdelete --host quantumleap entities --type Thing
```
