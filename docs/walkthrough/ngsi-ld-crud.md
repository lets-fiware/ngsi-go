# NGSI Go walkthrough for NGSI-LD

## Prerequisites

```
git clone https://github.com/FIWARE/tutorials.CRUD-Operations.git
cd tutorials.CRUD-Operations
git checkout NGSI-LD

./services orion

curl http://localhost:1026/version
```

## NGSI-LD CRUD operation

### Add Host

```
ngsi broker add --host orion-ld --brokerHost http://localhost:1026
```

### Add Context

```
ngsi context add --name tutorial --url http://context-provider:3000/data-models/ngsi-context.jsonld
```

### Version

```
ngsi version -h orion-ld
```

## Create Operations

### Create a New Data Entity

This example adds a new **TemperatureSensor** entity to the context.

#### :one: Request:

```
ngsi create entity --link tutorial \
--data '{
      "id": "urn:ngsi-ld:TemperatureSensor:001",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 25,
            "unitCode": "CEL"
      }
}'
```

#### :two: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001
```

### Create New Attributes

This example adds a new `batteryLevel` Property and a `controlledAsset` Relationship to the existing
**TemperatureSensor** entity with `id=urn:ngsi-ld:TemperatureSensor:001`.

#### :three: Request:

```
ngsi append attrs --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 \
--data '{
       "batteryLevel": {
            "type": "Property",
            "value": 0.9,
            "unitCode": "C62"
      },
      "controlledAsset": {
            "type": "Relationship",
            "object": "urn:ngsi-ld:Building:barn002"
      }
}'
```
#### :four: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001
```

#### :five: Request:

```
ngsi create entities --link tutorial \
--data '[
    {
      "id": "urn:ngsi-ld:TemperatureSensor:002",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 20,
            "unitCode": "CEL"
      }
    },
    {
      "id": "urn:ngsi-ld:TemperatureSensor:003",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 2,
            "unitCode": "CEL"
      }
    },
     {
      "id": "urn:ngsi-ld:TemperatureSensor:004",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 100,
            "unitCode": "CEL"
      }
    }
]'
```

### Batch Create/Overwrite New Data Entities

This example uses the convenience batch processing endpoint to add or amend two **TemperatureSensor** entities in the
context.

-   if an entity already exists, the request will update that entity's attributes.
-   if an entity does not exist, a new entity will be created.

#### :six: Request:

```
ngsi upsert entities --link tutorial \
--data '[
    {
      "id": "urn:ngsi-ld:TemperatureSensor:002",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 21,
            "unitCode": "CEL"
      }
    },
    {
      "id": "urn:ngsi-ld:TemperatureSensor:003",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 27,
            "unitCode": "CEL"
      }
    }
]'
```

## Read Operations

-   The `/ngsi-ld/v1/entities` endpoint is used for listing entities
-   The `/ngsi-ld/v1/entities/<entity>` endpoint is used for retrieving the details of a single entity.

For read operations the `@context` must be supplied in a `Link` header.

### Filtering

-   The `options` parameter (combined with the `attrs` parameter) can be used to filter the returned fields
-   The `q` parameter can be used to filter the returned entities

### Read a Data Entity (verbose)

This example reads the full context from an existing **TemperatureSensor** entity with a known `id`.

#### :seven: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --sysAttrs
```

### Read an Attribute from a Data Entity

This example reads the value of a single attribute (`temperature`) from an existing **TemperatureSensor** entity with a
known `id`.

#### :eight: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --attrs temperature
```

### Read a Data Entity (key-value pairs)

This example reads the key-value pairs from the context of an existing **TemperatureSensor** entities with a known `id`.

#### :nine: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --keyValues
```

### Read Multiple attributes values from a Data Entity

This example reads the value of two attributes (`category` and `temperature`) from the context of an existing
**TemperatureSensor** entity with a known `id`.

#### :one::zero: Request:

```
ngsi get entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --keyValues --attrs category,temperature
```

### List all Data Entities (verbose)

This example lists the full context of all **TemperatureSensor** entities.

#### :one::one: Request:

```
ngsi list entities --link tutorial --type TemperatureSensor
```

### List all Data Entities (key-value pairs)

This example lists the `temperature` attribute of all **TemperatureSensor** entities.

#### :one::two: Request:

```
ngsi list entities -link tutorial --type TemperatureSensor --attrs temperature --keyValues
```

### Filter Data Entities by ID

This example lists selected data from two **TemperatureSensor** entities chosen by `id`. Note that every `id` must be
unique, so `type` is not required for this request. To filter by `id` add the entries in a comma delimited list.

#### :one::three: Request:

```
ngsi list entities --link tutorial --id urn:ngsi-ld:TemperatureSensor:001,urn:ngsi-ld:TemperatureSensor:002 --attrs temperature --keyValues
```

## Update Operations

Overwrite operations are mapped to HTTP PATCH:

-   The `/ngsi-ld/v1/entities/<entity-id>/attrs/<attribute>` endpoint is used to update an attribute
-   The `/ngsi-ld/v1/entities/<entity-id>/attrs` endpoint is used to update multiple attributes

### Overwrite the value of an Attribute value

This example updates the value of the `category` attribute of the Entity with `id=urn:ngsi-ld:TemperatureSensor:001`

#### :one::four: Request:

```
ngsi update attr --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --attrName category \
--data '{
    "value": ["sensor", "actuator"],
    "type": "Property"
}'
```

### Overwrite Multiple Attributes of a Data Entity

This example simultaneously updates the values of both the `category` and `controlledAsset` attributes of the Entity
with `id=urn:ngsi-ld:TemperatureSensor:001`.

#### :one::five: Request:

```
ngsi update attrs --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 \
--data '{
      "category": {
            "value": [
                  "sensor",
                  "actuator"
            ],
            "type": "Property"
      },
      "controlledAsset": {
            "type": "Relationship",
            "object": "urn:ngsi-ld:Building:barn001"
      }
}'
```

### Batch Update Attributes of Multiple Data Entities

This example uses the convenience batch processing endpoint to update existing sensors.

#### :one::six: Request:

```
ngsi upsert entities --link tutorial --update \
--data '[
  {
    "id": "urn:ngsi-ld:TemperatureSensor:003",
    "type": "TemperatureSensor",
    "category": {
      "type": "Property",
      "value": [
        "actuator",
        "sensor"
      ]
    }
  },
  {
    "id": "urn:ngsi-ld:TemperatureSensor:004",
    "type": "TemperatureSensor",
    "category": {
      "type": "Property",
      "value": [
        "actuator",
        "sensor"
      ]
    }
  }
]'
```

### Batch Replace Entity Data

This example uses the convenience batch processing endpoint to replace entity data of existing sensors.

#### :one::seven: Request:

```
ngsi update entities --link tutorial --replace \
--data '[
  {
    "id": "urn:ngsi-ld:TemperatureSensor:003",
    "type": "TemperatureSensor",
    "category": {
      "type": "Property",
      "value": [
        "actuator",
        "sensor"
      ]
    }
  },
  {
    "id": "urn:ngsi-ld:TemperatureSensor:004",
    "type": "TemperatureSensor",
    "temperature": {
      "type": "Property",
      "value": [
        "actuator",
        "sensor"
      ]
    }
  }
]'
```

## Delete Operations

Delete Operations map to HTTP DELETE.

-   The `/ngsi-ld/v1/entities/<entity-id>` endpoint can be used to delete an entity
-   The `/ngsi-ld/v1/entities/<entity-id>/attrs/<attribute>` endpoint can be used to delete an attribute

The response will be **204 - No Content** if the operation is successful or **404 - Not Found** if the operation fails.

### Data Relationships

If there are entities within the context which relate to one another, you must be careful when deleting an entity. You
will need to check that no references are left dangling once the entity has been deleted.

Organizing a cascade of deletions is beyond the scope of this tutorial, but it would be possible using a batch delete
request.

### Delete an Entity

This example deletes the entity with `id=urn:ngsi-ld:TemperatureSensor:004` from the context.

#### :one::eight: Request:

```
ngsi delete entity --link tutorial --id urn:ngsi-ld:TemperatureSensor:004
```

### Delete an Attribute from an Entity

This example removes the `batteryLevel` attribute from the entity with `id=urn:ngsi-ld:TemperatureSensor:001`.

#### :one::nine: Request:

```
ngsi delete attr --link tutorial --id urn:ngsi-ld:TemperatureSensor:001 --attrName batteryLevel
```

### Batch Delete Multiple Entities

This example uses the convenience batch processing endpoint to delete some **TemperatureSensor** entities.

#### :two::zero: Request:

```
ngsi delete entities --link tutorial \
--data '[
  "urn:ngsi-ld:TemperatureSensor:002",
  "urn:ngsi-ld:TemperatureSensor:003"
]'
```

### Batch Delete Multiple Attributes from an Entity

This example uses the PATCH `/ngsi-ld/v1/entities/<entity-id>/attrs` endpoint to delete some attributes from a
**TemperatureSensor** entity.

#### :two::one: Request:

```
ngsi update attrs --link tutorial --id urn:ngsi-ld:TemperatureSensor:001
--data '{
      "category": {
            "value": null,
            "type": "Property"
      },
      "controlledAsset": {
            "type": "Relationship",
            "object": null
      }
}'
```

### Find existing data relationships

This example returns a header indicating whether any linked data relationships remain against the entity
`urn:ngsi-ld:Building:barn002`

#### :two::two: Request:

```
ngsi list entities --link tutorial --type TemperatureSensor --query controlledAsset==\"urn:ngsi-ld:Building:barn002\" --count
```
