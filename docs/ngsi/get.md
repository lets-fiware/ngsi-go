# get - NGSI command

This command gets an entity, an attribute, multiple attributes, a subscription or a registration.

-   [Get an entity](#get-an-entity)
-   [Get an entities](#get-an-entities)
-   [Get temporal entity](#get-temporal-entity)
-   [Get an attribute](#get-an-attribute)
-   [Get multiple attributes](#get-multiple-attributes)
-   [Get a type](#get-a-type)
-   [Get a subscription](#get-a-subscription)
-   [Get a registration](#get-a-registration)
-   [Get a JSON-LD context](#get-a-json-ld-context)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="get-an-entity"></a>

## Get an entity

This command gets entity.

```console
ngsi get [command options] entity [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --orderBy VALUE           | orderBy                                |
| --count, -C               | count (default: false)                 |
| --keyValues, -K           | keyValues (default: false)             |
| --values, -V              | values (default: false)                |
| --unique, -U              | unique (default: false)                |
| --verbose, -v             | verbose (default: false)               |
| --lines, -1               | lines (default: false)                 |
| --data VALUE, -d VALUE    | entities data                          |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product
```

```json
{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade","metadata":{}},"price":{"type":"Integer","value":99,"metadata":{}},"size":{"type":"Text","value":"S","metadata":{}}}
```

#### Request:

```console
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product --keyValues
```

```json
{"id":"urn:ngsi-ld:Product:010","type":"Product","name":"Lemonade","price":99,"size":"S"}
```

#### Request:

```console
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product --keyValues --attrs size
```

```json
{"id":"urn:ngsi-ld:Product:010","type":"Product","size":"S"}
```

<a name="get-an-entities"></a>

## Get multiple entities

This command gets multiple entities.

```console
ngsi get [command options] entities [options]
```

### Options

| Options                   | Description                                                      |
| ------------------------- | ---------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                           |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                             |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                         |
| --id VALUE, -i VALUE      | entity id (required)                                             |
| --type VALUE, -t VALUE    | entity type                                                      |
| --attrs VALUE             | attributes                                                       |
| --keyValues, -K           | keyValues (default: false)                                       |
| --values, -V              | values (default: false)                                          |
| --unique, -U              | unique (default: false)                                          |
| --sysAttrs, -S            | sysAttrs (default: false)                                        |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                              |
| --acceptJson              | set accecpt header to application/json (LD) (default: false)     |
| --acceptGeoJson           | set accecpt header to application/geo+json (LD) (default: false) |
| --pretty, -P              | pretty format (default: false)                                   |
| --safeString VALUE        | use safe string (VALUE: on/off)                                  |
| --help                    | show help (default: true)                                        |

### Examples

#### Request:

```console
ngsi get entities --data '{"entities": [{"type": "Device", "idPattern": ".*"}],"attrs":["name"]}'
```

<a name="Get temporal entity"></a>

## Get temporal entity

This command gets a temporal entity.

```console
ngsi get [common options] tentity [options]
```

### Options

| Options                   | Description                                                       |
| ------------------------- | ----------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                            |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                              |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                          |
| --id VALUE, -i VALUE      | temporal entity id (required)                                     |
| --attrs VALUE             | attributes                                                        |
| --timeProperty VALUE      | timeProperty (LD)                                                 |
| --fromDate VALUE          | starting date from which data should be retrieved                 |
| --toDate VALUE            | final date until which data should be retrieved                   |
| --lastN VALUE             | number of data entries to retrieve since the final date backwards |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                               |
| --temporalValues          | temporal simplified representation of entity (default: false)     |
| --sysAttrs, -S            | sysAttrs (default: false)                                         |
| --acceptJson              | set accecpt header to application/json (LD) (default: false)      |
| --pretty, -P              | pretty format (default: false)                                    |
| --safeString VALUE        | use safe string (VALUE: on/off)                                   |
| --etsi10                  | ETSI CIM 009 V1.0 (default: false)                                |
| --help                    | show help (default: true)                                         |

<a name="get-an-attribute"></a>

## Get an attribute

This command gets an attribute value.

```console
ngsi get [common options] attr [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --attr VALUE              | attribute name (required)              |
| --type VALUE, -t VALUE    | entity type                            |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi get attr --id urn:ngsi-ld:Product:010 --type Product --attr size "S"
```

<a name="get-multiple-attributes"></a>

## Get multiple attributes

This command gets attributes.

```console
ngsi get [common options] attrs [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --type VALUE, -t VALUE    | entity type                            |
| --attrs VALUE             | attributes                             |
| --metadata VALUE          | metadata (v2)                          |
| --keyValues, -K           | keyValues (default: false)             |
| --values, -V              | values (default: false)                |
| --unique, -U              | unique (default: false)                |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi get attrs --type Product --id urn:ngsi-ld:Product:001 --attrs name,price
```

```json
{"name":{"type":"Text","value":"Beer","metadata":{}},"price":{"type":"Integer","value":99,"metadata":{}}}
```

#### Request:

```console
ngsi get attrs --type Product --id urn:ngsi-ld:Product:001 --attrs name,price --keyValues
```

```json
{"name":"Beer","price":99}
```

#### Request:

```console
ngsi get attrs --type Product --id urn:ngsi-ld:Product:001 --attrs name,price --values
```

```json
["Beer",99]
```

<a name="get-a-type"></a>

## Get a type

This command gets type.

```console
ngsi get [common options] type [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --type VALUE, -t VALUE    | entity type                            |
| --pretty, -P              | pretty format (default: false)         |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --help                    | show help (default: true)              |

### Examples for NGSIv2

#### Request:

```console
ngsi get --host orion type --type Product --pretty
```

```json
{
  "attrs": {
    "name": {
      "types": [
        "Text"
      ]
    },
    "price": {
      "types": [
        "Integer"
      ]
    },
    "size": {
      "types": [
        "Text"
      ]
    }
  },
  "count": 1
}
```

### Examples for NGSI-LD

#### Request:

```console
ngsi get --host orion-ld type --pretty https://uri.fiware.org/ns/data-models#TemperatureSensor
```

```json
{
  "@context": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld",
  "id": "https://uri.fiware.org/ns/data-models#TemperatureSensor",
  "type": "EntityTypeInformation",
  "typeName": "https://uri.fiware.org/ns/data-models#TemperatureSensor",
  "entityCount": 1,
  "attributeDetails": [
    {
      "id": "https://uri.fiware.org/ns/data-models#category",
      "type": "Attribute",
      "attributeName": "https://uri.fiware.org/ns/data-models#category",
      "attributeTypes": [
        "Property"
      ]
    },
    {
      "id": "https://w3id.org/saref#temperature",
      "type": "Attribute",
      "attributeName": "https://w3id.org/saref#temperature",
      "attributeTypes": [
        "Property"
      ]
    }
  ]
}
```

<a name="get-a-subscription"></a>

## Get a subscription

This command gets a subscription.

```console
ngsi get [common options] subscription [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | subscription id (required)             |
| --localTime               | localTime (default: false)             |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --raw                     | handle raw data (default: false)       |
| --help                    | show help (default: true)              |

### Examples for NGSIv2

#### Request:

```console
ngsi get subscription --id 5fa7988a627088ba9b91b1c1 --pretty
```

```json
{
  "id": "5fa7988a627088ba9b91b1c1",
  "subject": {
    "entities": [
      {
        "idPattern": ".*",
        "type": "Sensor"
      }
    ],
    "condition": {
      "attrs": [
        "temperature"
      ]
    }
  },
  "notification": {
    "onlyChangedAttrs": false,
    "http": {
      "url": "http://192.168.0.1/"
    },
    "attrs": [
      "notification"
    ],
    "attrsFormat": "normalized"
  },
  "expires": "2020-11-09T07:04:42.000Z",
  "status": "active"
}
```

#### Request:

```console
ngsi get subscription --id 5fa7988a627088ba9b91b1c1 --localTime --pretty
```

```json
{
  "id": "5fa7988a627088ba9b91b1c1",
  "subject": {
    "entities": [
      {
        "idPattern": ".*",
        "type": "Sensor"
      }
    ],
    "condition": {
      "attrs": [
        "temperature"
      ]
    }
  },
  "notification": {
    "onlyChangedAttrs": false,
    "http": {
      "url": "http://192.168.0.1/"
    },
    "attrs": [
      "notification"
    ],
    "attrsFormat": "normalized"
  },
  "expires": "2020-11-09T16:04:42.000+0900",
  "status": "active"
}
```

### Examples for NGSI-LD

#### Request:

```console
ngsi get subscription --id urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce --pretty
```

```json
{
  "id": "urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce",
  "type": "Subscription",
  "description": "Notify me of low stock in Store 001",
  "entities": [
    {
      "type": "Shelf"
    }
  ],
  "watchedAttributes": [
    "numberOfItems"
  ],
  "q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store001",
  "notification": {
    "attributes": [
      "numberOfItems",
      "stocks",
      "locatedIn"
    ],
    "format": "keyValues",
    "endpoint": {
      "uri": "https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld",
      "accept": "application/ld+json"
    }
  }
}
```

<a name="get-a-registration"></a>

## Get a registration

This command gets a registration.

```console
ngsi get [common options] registration [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | registration id (required)             |
| --localTime               | localTime (default: false)             |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Examples for NGSI-LD

#### Request:

```console
ngsi get registration --id urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0 --pretty
```

```json
{
  "id": "urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0",
  "type": "ContextSourceRegistration",
  "description": "ContextSourceRegistration",
  "endpoint": "http://context-provider:3000/static/tweets",
  "information": [
    {
      "entities": [
        {
          "id": "urn:ngsi-ld:Building:store001",
          "type": "Building"
        }
      ],
      "properties": [
        "tweets"
      ]
    }
  ]
}
```

<a name="get-a-json-ld-context"></a>

## Get a JSON-LD context

This command gets a JSON-LD context.

```console
ngsi get [common options] ldContext [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | jsonldContexts id (LD) (required)      |
| --pretty, -P              | pretty format (default: false)         |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi get --host orion-ld ldContext --id 2fa4dbc4-ece8-11eb-a645-0242c0a8a010
```
#### Response:

```json
{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}
```
