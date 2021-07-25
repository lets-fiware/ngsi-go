# delete - NGSI command

-   [Delete an entity](#delete-an-entity)
-   [Delete multiple entities](#delete-multiple-entities)
-   [Delete temporal entity](#delete-temporal-entity)
-   [Delete an attribute from an entity](#delete-an-attribute-from-an-entity)
-   [Delete an attribute from a temporal entity](#delete-an-attribute-from-a-temporal-entity)
-   [Delete a subscription](#delete-a-subscription)
-   [Delete a registration](#delete-a-registration)
-   [Delete a JSON-LD context](#delete-a-json-ld-context)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="delete-an-entity"></a>

## Delete an entity

This command will delete entity

```console
ngsi delete [common options] entity [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | specify entity id          |
| --type value, -t value | specify entity Type        |
| --link value, -L value | specify @context           |
| --help                 | show help (default: false) |

### Example

#### Request:

```console
ngsi delete entity --id urn:ngsi-ld:Product:010
```

<a name="delete-multiple-entities"></a>

## Delete multiple entities

This command deletes entities or attributes from entities.

```console
ngsi delete [common options] entities [options]
```

### Options

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --keyValues, -k           | specify keyValues (default: false) |
| --data value, -d value    | specify data                       |
| --link value, -L value    | specify @context                   |
| --help                    | show help (default: false)         |

### Example

#### Request:

```console
ngsi delete entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:001", "type":"Product"
  },
  {
    "id":"urn:ngsi-ld:Product:002", "type":"Product"
  }
]'
```

#### Request:

```console
ngsi delete entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:003", "type":"Product",
    "price":{},
    "name": {}
  },
  {
    "id":"urn:ngsi-ld:Product:004", "type":"Product",
    "price":{},
    "name": {}
  }
]'
```

<a name="delete-temporal-entity"></a>

## Delete a temporal entity

This command will delete a temporal entity.

```console
ngsi delete [common options] tentity [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | specify entity id          |
| --link value, -L value | @context (LD)              |
| --help                 | show help (default: false) |

<a name="delete-an-attribute-from-an-entity"></a>

## Delete an Attribute from an Entity

This command will delete attribute.

```console
ngsi delete [common options] attr [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | specify entity id          |
| --type value, -t value | specify entity Type        |
| --attr value           | specify attribute name     |
| --link value, -L value | specify @context           |
| --help                 | show help (default: false) |

### Example

#### Request:

```console
ngsi delete attr --id urn:ngsi-ld:Product:001 --attr specialOffer
```

<a name="delete-an-attribute-from-a-temporal-entity"></a>

## Delete an attribute from a temporal entity

This command will delete an attribute from a temporal entity.

```console
ngsi delete [common options] tattr [options]
```

### Options

| Options                | Description                                         |
| ---------------------- | --------------------------------------------------- |
| --id value, -i value   | specify entity id                                   |
| --attr value           | specify attribute name                              |
| --deleteAll            | all atribute instances are deleted (default: false) |
| --datasetId value      | datasetId of the dataset to be deleted              |
| --instanceId value     | attribute instance id (LD)                          |
| --link value, -L value | @context (LD)                                       |
| --help                 | show help (default: false)                          |

<a name="delete-a-subscription"></a>

## Delete a subscription

This command deletes subscriptions.

```console
ngsi delete [common options] subscription [options]
```

### Options

| Options              | Description                |
| -------------------- | -------------------------- |
| --id value, -i value | specify subscription id    |
| --help               | show help (default: false) |

### Example

#### Request:

```console
ngsi delete subscription --id urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf
```

<a name="delete-a-registration"></a>

## Delete a registration

This command deletes registration.

```console
ngsi delete [common options] registration [options]
```

### Options

| Options              | Description                |
| -------------------- | -------------------------- |
| --id value, -i value | specify registration id    |
| --help               | show help (default: false) |

### Example

#### Request:

```console
ngsi delete registration --id urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0
```

<a name="delete-a-json-ld-context"></a>

## Delete a JSON-LD context

This command deletes a JSON-LD context.

```console
ngsi delete [common options] ldContext [options]
```

### Options

| Options              | Description                |
| -------------------- | -------------------------- |
| --id value, -i value | jsonldContexts id          |
| --help               | show help (default: false) |

### Example

#### Request:

```console
ngsi delete --host orion-ld ldContext --id d42e7ffe-ed21-11eb-bc92-0242c0a8a010
```
