# NGSI Go walkthrough for NGSIv2

## Prerequisites

```
git clone https://github.com/FIWARE/tutorials.CRUD-Operations.git
cd tutorials.CRUD-Operations
git checkout NGSI-v2

./services start
sed -i -e 's/orion:1026/localhost:1026/' ./import-data
./import-data

curl http://localhost:1026/version
```

## NGSI v2 CRUD operation

### Add host

```
ngsi broker add --host orion --brokerHost http://localhost:1026 --ngsiType v2
```

### Version

```
ngsi version -h orion
```

### Create a New Data Entity

This example adds a new **Product** entity ("Lemonade" at 99 cents) to the context.

#### :one: Request:

```
ngsi create entity \
  --data ' {
      "id":"urn:ngsi-ld:Product:010", "type":"Product",
      "name":{"type":"Text", "value":"Lemonade"},
      "size":{"type":"Text", "value": "S"},
      "price":{"type":"Integer", "value": 99}
}'
```

#### :two: Request:

```
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product
```

### Create a New Attribute

This example adds a new `specialOffer` attribute to the existing **Product** entity with `id=urn:ngsi-ld:Product:001`.

#### :three: Request:

```
ngsi append attrs --id urn:ngsi-ld:Product:010 \
  --data '{
      "specialOffer":{"value": true}
}'
```

#### :four: Request:

```
ngsi get entity --id urn:ngsi-ld:Product:001 --type Product
```

### Batch Create New Data Entities or Attributes

This example uses the convenience batch processing endpoint to add two new **Product** entities and one new attribute
(`offerPrice`) to the context.

#### :five: Request:

```
ngsi create entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:011", "type":"Product",
    "name":{"type":"Text", "value":"Brandy"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:012", "type":"Product",
    "name":{"type":"Text", "value":"Port"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1099}
  },
  {
    "id":"urn:ngsi-ld:Product:001", "type":"Product",
    "offerPrice":{"type":"Integer", "value": 89}
  }
]'
```

### Batch Create/Overwrite New Data Entities

This example uses the convenience batch processing endpoint to add or amend two **Product** entities and one attribute
(`offerPrice`) to the context.

-   if an entity already exists, the request will update that entity's attributes.
-   if an entity does not exist, a new entity will be created.

#### :six: Request:

```
ngsi upsert entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:011", "type":"Product",
    "name":{"type":"Text", "value":"Brandy"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:012", "type":"Product",
    "name":{"type":"Text", "value":"Port"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1099}
  }
]'
```

### Read a Data Entity (verbose)

This example reads the full context from an existing **Product** entity with a known `id`.

#### :seven: Request:

```
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product
```

### Read an Attribute from a Data Entity

This example reads the value of a single attribute (`name`) from an existing **Product** entity with a known `id`.

#### :eight: Request:

```
ngsi get attr --id urn:ngsi-ld:Product:001 --attrName name
```

### Read a Data Entity (key-value pairs)

This example reads the key-value pairs of two attributes (`name` and `price`) from the context of existing **Product**
entities with a known `id`.

#### :nine: Request:

```
ngsi get entity --keyValues --type Product --id urn:ngsi-ld:Product:001 --attrs name,price
```

```
{"id":"urn:ngsi-ld:Product:001","type":"Product","name":"Lemonade","price":99}
```

```
ngsi get attrs --keyValues --type Product --id urn:ngsi-ld:Product:001 --attrs name,price
```

```
{"name":"Lemonade","price":99}
```

### Read Multiple attributes values from a Data Entity

This example reads the value of two attributes (`name` and `price`) from the context of existing **Product** entities
with a known ID.

#### :one::zero: Request:

```
ngsi get attrs --id urn:ngsi-ld:Product:001 --attrs name,price --values
```

### List all Data Entities (verbose)

This example lists the full context of all **Product** entities.

#### :one::one: Request:

```
ngsi list entities --type Product
```

### List all Data Entities (key-value pairs)

This example lists the `name` and `price` attributes of all **Product** entities.

#### :one::two: Request:

```
ngsi list entities --type Product -attrs name,price --keyValues
```

### List Data Entity by ID

This example lists the `id` and `type` of all **Product** entities.

#### :one::three: Request:

```
ngsi list entities --type Product -attrs id
```

## Update Operations

### Overwrite the value of an Attribute value

This example updates the value of the price attribute of the Entity with `id=urn:ngsi-ld:Product:001`

#### :one::four: Request:

```
ngsi update attr --id urn:ngsi-ld:Product:001 --attrName price --data 89
```

### Overwrite Multiple Attributes of a Data Entity

This example simultaneously updates the values of both the price and name attributes of the Entity with
`id=urn:ngsi-ld:Product:001`.

#### :one::five: Request:

```
ngsi update attrs --id urn:ngsi-ld:Product:001 \
--data ' {
    "price":{"type":"Integer", "value": 89},
    "name": {"type":"Text", "value": "Ale"}
}'
```

### Batch Overwrite Attributes of Multiple Data Entities

This example uses the convenience batch processing endpoint to update existing products.

#### :one::six: Request:

```
ngsi update entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:001", "type":"Product",
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:002", "type":"Product",
    "price":{"type":"Integer", "value": 1199},
    "size": {"type":"Text", "value": "L"}
  }
]'
```

### Batch Create/Overwrite Attributes of Multiple Data Entities

This example uses the convenience batch processing endpoint to update existing products.

#### :one::seven: Request:

```
ngsi upsert entities \
--data '[
{
    "id":"urn:ngsi-ld:Product:001", "type":"Product",
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:002", "type":"Product",
    "price":{"type":"Integer", "value": 1199},
    "specialOffer": {"type":"Boolean", "value":  true}
  }
]'
```

### Batch Replace Entity Data

This example uses the convenience batch processing endpoint to replace entity data of existing products.

#### :one::eight: Request:

```
ngsi replace entities \
--data '[
{
    "id":"urn:ngsi-ld:Product:010", "type":"Product",
    "price":{"type":"Integer", "value": 1199}
  }
]'
```

## Delete Operations

### Delete an Entity

This example deletes the entity with `id=urn:ngsi-ld:Product:001` from the context.

#### :one::nine: Request:

```
ngsi delete entity --id urn:ngsi-ld:Product:010
```

### Delete an Attribute from an Entity

This example removes the `specialOffer` attribute from the entity with `id=urn:ngsi-ld:Product:001`.

#### :two::zero: Request:

```
ngsi delete attr --id urn:ngsi-ld:Product:001 --attrName specialOffer
```

### Batch Delete Multiple Entities

This example uses the convenience batch processing endpoint to delete some **Product** entities.

#### :two::one: Request:

```
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

### Batch Delete Multiple Attributes from an Entity

This example uses the convenience batch processing endpoint to delete some attributes from a **Product** entity.

#### :two::two: Request:

```
ngsi delete entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:003", "type":"Product",
    "price":{},
    "name": {}
  }
]'
```
### Find existing data relationships

This example returns the key of all entities directly associated with the `urn:ngsi-ld:Product:001`.

#### :two::three: Request:

```
ngsi list entities -q "refProduct%==urn:ngsi-ld:Product:001" --attrs type
```
