# list - NGSI command

This command lists types, entities, subscriptions or registrations.

-   [List multiple types](#list-multiple-types)
-   [List multiple attributes](#list-multiple-attributes)
-   [List multiple entities](#list-multiple-entities)
-   [List multipletemporal entities](#list-multiple-temporal-entities)
-   [List multiple subscriptions](#list-multiple-subscriptions)
-   [List multiple registrations](#list-multiple-registrations)
-   [List JSON-LD contexts](#list-json-ld-contexts)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="list-multiple-types"></a>

## List multiple types

This command lists types.

```console
ngsi list [common options] types [options]
```

### Options

| Options                   | Description                                            |
| ------------------------- | ------------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                 |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                               |
| --details, -d             | detailed entity type information (LD) (default: false) |
| --json, -j                | JSON format (default: false)                           |
| --pretty, -P              | pretty format (default: false)                         |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                    |
| --help                    | show help (default: true)                              |

### Examples for NGSIv2

#### Request:

```console
ngsi list --host orion types
```

```text
InventoryItem
Product
Shelf
Store
```

#### Request:

```console
ngsi list --host orion types --json
```

```json
["InventoryItem","Product","Shelf","Store"]
```
### Examples for NGSI-LD

#### Request:

```console
ngsi list --host orion-ld types --details --pretty
```

```
[
  {
    "@context": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld",
    "id": "https://uri.fiware.org/ns/data-models#TemperatureSensor",
    "type": "EntityType",
    "typeName": "https://uri.fiware.org/ns/data-models#TemperatureSensor",
    "attributeNames": [
      "https://uri.fiware.org/ns/data-models#category",
      "https://w3id.org/saref#temperature"
    ]
  }
]
```

<a name="list-multiple-attributes"></a>

## List multiple attributes

This command lists attributes.

```console
ngsi list [common options] attributes [options]
```

### Options

| Options                   | Description                                          |
| ------------------------- | ---------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)               |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                 |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                             |
| --attr VALUE              | attribute name                                       |
| --details, -d             | detailed attribute information (LD) (default: false) |
| --pretty, -P              | pretty format (default: false)                       |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                  |
| --help                    | show help (default: true)                            |

### Examples

```console
ngsi list --host orion-ld attributes
```

```console
ngsi list --host orion-ld attributes --link ctx
```

```console
ngsi list --host orion-ld attributes --attr "https://w3id.org/saref#temperature"
```

```console
ngsi list --host orion-ld attributes --attr temperature --link ctx
```

<a name="list-multiple-entities"></a>

## List multiple entities

This command lists multiple entities.

```console
ngsi list [common options] entities [options]
```

### Options

| Options                   | Description                                                      |
| ------------------------- | ---------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                           |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                             |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                         |
| --id VALUE, -i VALUE      | entity id                                                        |
| --type VALUE, -t VALUE    | entity type                                                      |
| --idPattern VALUE         | idPattern                                                        |
| --typePattern VALUE       | typePattern (v2)                                                 |
| --query VALUE, -q VALUE   | filtering by attribute value                                     |
| --mq VALUE, -m VALUE      | filtering by metadata (v2)                                       |
| --georel VALUE            | georel                                                           |
| --geometry VALUE          | geometry                                                         |
| --coords VALUE            | coords                                                           |
| --attrs VALUE             | attributes                                                       |
| --metadata VALUE          | metadata (v2)                                                    |
| --orderBy VALUE           | orderBy                                                          |
| --count, -C               | count (default: false)                                           |
| --keyValues, -K           | keyValues (default: false)                                       |
| --values, -V              | values (default: false)                                          |
| --unique, -U              | unique (default: false)                                          |
| --skipForwarding          | skip forwarding to CPrs (v2) (default: false)                    |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                              |
| --acceptJson              | set accecpt header to application/json (LD) (default: false)     |
| --acceptGeoJson           | set accecpt header to application/geo+json (LD) (default: false) |
| --verbose, -v             | verbose (default: false)                                         |
| --lines, -1               | lines (default: false)                                           |
| --pretty, -P              | pretty format (default: false)                                   |
| --safeString VALUE        | use safe string (VALUE: on/off)                                  |
| --help                    | show help (default: true)                                        |

### Example

#### Request:

```console
ngsi list entities --type Product
```

```text
urn:ngsi-ld:Product:001
urn:ngsi-ld:Product:002
urn:ngsi-ld:Product:003
urn:ngsi-ld:Product:004
urn:ngsi-ld:Product:005
urn:ngsi-ld:Product:006
urn:ngsi-ld:Product:007
urn:ngsi-ld:Product:008
urn:ngsi-ld:Product:009
urn:ngsi-ld:Product:010
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112
urn:ngsi-ld:Product:101
```

#### Request:

```console
ngsi list entities --type Product --count
```

```text
14
```

#### Request:

```console
ngsi list entities --type Product --idPattern '0{2}'
```

```text
urn:ngsi-ld:Product:001
urn:ngsi-ld:Product:002
urn:ngsi-ld:Product:003
urn:ngsi-ld:Product:004
urn:ngsi-ld:Product:005
urn:ngsi-ld:Product:006
urn:ngsi-ld:Product:007
urn:ngsi-ld:Product:008
urn:ngsi-ld:Product:009
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}'
```

```text
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' --count
```

```text
3
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' --verbose --pretty
```

```json
[
  {
    "id": "urn:ngsi-ld:Product:110",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Lemonade"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 99
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "S"
    },
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:111",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Brandy"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 1199
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "M"
    },
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:112",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Port"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 1099
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "M"
    },
    "type": "Product"
  }
]
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' --verbose --keyValues --pretty
```

```json
[
  {
    "id": "urn:ngsi-ld:Product:110",
    "name": "Lemonade",
    "price": 99,
    "size": "S",
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:111",
    "name": "Brandy",
    "price": 1199,
    "size": "M",
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:112",
    "name": "Port",
    "price": 1099,
    "size": "M",
    "type": "Product"
  }
]
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' --count
```

```text
3
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' | xargs -L 1 ngsi delete entity --id
```

#### Request:

```console
ngsi list entities --type Product --idPattern '1{2}' --count
```

```text
0
```

#### Request:

```console
ngsi list entities -q "refProduct%==urn:ngsi-ld:Product:001" --attrs type
```

<a name="list-temporal-entities"></a>

# List temporal entities

This command lists multiple tempral entities.

<a name="list-multiple-subscriptions"></a>

```console
ngsi list [common options] tentities [options]
```

### Options

| Options                   | Description                                                       |
| ------------------------- | ----------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                            |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                              |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                          |
| --id VALUE, -i VALUE      | temporal entity id                                                |
| --type VALUE, -t VALUE    | entity type                                                       |
| --idPattern VALUE         | idPattern                                                         |
| --attrs VALUE             | attributes                                                        |
| --query VALUE, -q VALUE   | filtering by attribute value                                      |
| --csf VALUE               | context source filter (LD)                                        |
| --georel VALUE            | georel                                                            |
| --geometry VALUE          | geometry                                                          |
| --coords VALUE            | coords                                                            |
| --geoProperty VALUE       | geo property (LD)                                                 |
| --timeProperty VALUE      | timeProperty (LD)                                                 |
| --fromDate VALUE          | starting date from which data should be retrieved                 |
| --toDate VALUE            | final date until which data should be retrieved                   |
| --lastN VALUE             | number of data entries to retrieve since the final date backwards |
| --temporalValues          | temporal simplified representation of entity (default: false)     |
| --sysAttrs, -S            | sysAttrs (default: false)                                         |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                               |
| --acceptJson              | set accecpt header to application/json (LD) (default: false)      |
| --verbose, -v             | verbose (default: false)                                          |
| --lines, -1               | lines (default: false)                                            |
| --pretty, -P              | pretty format (default: false)                                    |
| --safeString VALUE        | use safe string (VALUE: on/off)                                   |
| --etsi10                  | ETSI CIM 009 V1.0 (default: false)                                |
| --help                    | show help (default: true)                                         |

## List multiple subscriptions

This command lists multiple subscriptions.

```console
ngsi list [common options] subscriptions [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --verbose, -v             | verbose (default: false)               |
| --json, -j                | JSON format (default: false)           |
| --status VALUE            | status                                 |
| --localTime               | localTime (default: false)             |
| --query VALUE, -q VALUE   | filtering by attribute value           |
| --items VALUE, -i VALUE   | itmes                                  |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --count, -C               | count (default: false)                 |
| --help                    | show help (default: true)              |

### Examples for NGSI-LD

#### Request:

```console
ngsi list subscriptions
```

```text
urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce
urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf
```

#### Request:

```console
ngsi list subscriptions --verbose
```

```text
urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce Notify me of low stock in Store 001
urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf LD Notify me of low stock in Store 002
```

#### Request:

```console
ngsi list subscriptions --json --pretty
```

```json
[
  {
    "description": "Notify me of low stock in Store 001",
    "entities": [
      {
        "type": "Shelf"
      }
    ],
    "id": "urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce",
    "notification": {
      "attributes": [
        "numberOfItems",
        "stocks",
        "locatedIn"
      ],
      "endpoint": {
        "accept": "application/ld+json",
        "uri": "https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld"
      },
      "format": "keyValues"
    },
    "q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store001",
    "type": "Subscription",
    "watchedAttributes": [
      "numberOfItems"
    ]
  },
  {
    "description": "LD Notify me of low stock in Store 002",
    "entities": [
      {
        "type": "Shelf"
      }
    ],
    "id": "urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf",
    "notification": {
      "attributes": [
        "numberOfItems",
        "stocks",
        "locatedIn"
      ],
      "endpoint": {
        "accept": "application/ld+json",
        "uri": "http://tutorial:3000/subscription/low-stock-store002"
      },
      "format": "normalized"
    },
    "q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002",
    "type": "Subscription",
    "watchedAttributes": [
      "numberOfItems"
    ]
  }
]
```

<a name="list-multiple-registrations"></a>

## List multiple registrations

This command lists multiple registrations.

```console
ngsi list [common options] registrations [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --verbose, -v             | verbose (default: false)               |
| --json, -j                | JSON format (default: false)           |
| --localTime               | localTime (default: false)             |
| --pretty, -P              | pretty format (default: false)         |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Examples for NGSI-LD

#### Request:

```console
ngsi list registrations
```

```text
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0
```

#### Request:

```console
ngsi list registrations -v
```

```text
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0 ContextSourceRegistration
```

#### Request:

```console
ngsi list registrations -j --pretty
```

```json
[
  {
    "description": "ContextSourceRegistration",
    "endpoint": "http://context-provider:3000/static/tweets",
    "id": "urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0",
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
    ],
    "type": "ContextSourceRegistration"
  }
]
```

<a name="list-json-ld-contexts"></a>

## List JSON-LD contexts

This command lists JSON-LD contexts.

```console
ngsi list [common options] ldContexts [options]
```

### Options

| Options                   | Description                                               |
| ------------------------- | --------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                    |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                      |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                  |
| --details, -d             | detailed jsonldContexts information (LD) (default: false) |
| --json, -j                | JSON format (default: false)                              |
| --pretty, -P              | pretty format (default: false)                            |
| --help                    | show help (default: true)                                 |

### Examples

#### Request:

```console
ngsi list --host orion-ld ldContexts
```

#### Response:

```console
fd564040-ece7-11eb-8e4a-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld
08d25d00-ece8-11eb-8d65-0242c0a8a010 http://atcontext:8000/ngsi-context.jsonld
0c6484d4-ece8-11eb-a312-0242c0a8a010 http://atcontext:8000/test-context.jsonld
30abb6fa-ece8-11eb-a645-0242c0a8a010 https://fiware.github.io/data-models/context.jsonld
31443434-ece8-11eb-a645-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld
2fa4dbc4-ece8-11eb-a645-0242c0a8a010 http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010
```
