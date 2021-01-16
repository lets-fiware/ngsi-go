# delete - NGSI command

-   [Delete an entity](#delete-an-entity)
-   [Delete multiple entities](#delete-multiple-entities)
-   [Delete an attribute from an Entity](#delete-an-attribute)
-   [Delete a subscription](#delete-a-subscription)
-   [Delete a registration](#delete-a-registration)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="delete-an-entity"/>

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

<a name="delete-multiple-entities"/>

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

<a name="delete-an-attribute"/>

## Delete an Attribute from an Entity

This commnad will delete attribute.

```console
ngsi delete [common options] attr [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | specify entity id          |
| --type value, -t value | specify entity Type        |
| --attrName value       | specify attribute name     |
| --link value, -L value | specify @context           |
| --help                 | show help (default: false) |

### Example

#### Request:

```console
ngsi delete attr --id urn:ngsi-ld:Product:001 --attrName specialOffer
```

<a name="delete-a-subscription"/>

## Delete a subscription

This commnad deletes subscriptions.

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

<a name="delete-a-registration"/>

## Delete a registration

This commnad deletes registrations.

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
