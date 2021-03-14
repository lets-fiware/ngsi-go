# get - NGSI command

This command gets an entity, an attribute, multiple attributes, a subscription or a registration.

-   [Get an entity](#get-an-entity)
-   [Get an entities](#get-an-entities)
-   [Get temporal entity](#get-temporal-entity)
-   [Get an attribute](#get-an-attribute)
-   [Get multiple attributes](#get-multiple-attributes)
-   [Get a subscription](#get-a-subscription)
-   [Get a registration](#get-a-registration)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="get-an-entity"></a>

## Get an entity

This command gets entity.

```console
ngsi get [command options] entity [options]
```

### Options

| Options                | Description                                                      |
| ---------------------- | ---------------------------------------------------------------- |
| --id value, -i value   | specify id                                                       |
| --type value, -t value | specify entity type                                              |
| --attrs value          | specify attributes                                               |
| --keyValues, -k        | specify keyValues (default: false)                               |
| --values, -V           | specify values (default: false)                                  |
| --unique, -u           | specify unique (default: false)                                  |
| --sysAttrs, -s         | specify sysAttrs (default: false)                                |
| --link value, -L value | specify @context                                                 |
| --acceptJson           | set accecpt header to application/json (LD) (default: false)     |
| --acceptGeoJson        | set accecpt header to application/geo+json (LD) (default: false) |
| --acceptJson           | set accecpt header to application/json (LD)                      |
| --pretty, -P           | pretty format (default: false)                                   |
| --safeString value     | use safe string (value: on/off)                                  |
| --help                 | show help (default: false)                                       |

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

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --orderBy value        | specify orderBy                    |
| --count, -C            | specify count (default: false)     |
| --keyValues, -k        | specify keyValues (default: false) |
| --values, -V           | specify values (default: false)    |
| --unique, -u           | specify unique (default: false)    |
| --verbose, -v          | specfiy verbose (default: false)   |
| --lines, -1            | specify lines (default: false)     |
| --data value, -d value | specify data                       |
| --pretty, -P           | pretty format (default: false)     |
| --safeString value     | use safe string (value: on/off)    |
| --help                 | show help (default: false)         |

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

| Options                | Description                                                                    |
| ---------------------- | ------------------------------------------------------------------------------ |
| --id value, -i value   | id                                                                             |
| --attrs value          | attributes                                                                     |
| --timeProperty value   | timeProperty (LD)                                                              |
| --fromDate value       | starting date from which data should be retrieved                              |
| --toDate value         | final date until which data should be retrieved                                |
| --lastN value          | number of data entries to retrieve since the final date backwards (default: 0) |
| --link value, -L value | @context (LD)                                                                  |
| --temporalValues       | temporal simplified representation of entity (default: false)                  |
| --sysAttrs, -S         | sysAttrs (default: false)                                                      |
| --acceptJson           | set accecpt header to application/json (LD) (default: false)                   |
| --pretty, -P           | pretty format (default: false)                                                 |
| --safeString value     | use safe string (value: on/off)                                                |
| --etsi10               | ETSI CIM 009 V1.0 (default: false)                                             |
| --help                 | show help (default: false)                                                     |

<a name="get-an-attribute"></a>

## Get an attribute

This command gets an attribute value.

```console
ngsi get [common options] attr [options]
```

### Options

| Options                | Description                     |
| ---------------------- | ------------------------------- |
| --id value, -i value   | specify id                      |
| --type value, -t value | specify entity type             |
| --attr value           | specify attribute name          |
| --pretty, -P           | pretty format (default: false)  |
| --safeString value     | use safe string (value: on/off) |
| --help                 | show help (default: false)      |

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

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --id value, -i value   | specify id                         |
| --type value, -t value | specify Entity Type                |
| --attrs value          | specify attrs                      |
| --metadata value       | specify metadata                   |
| --keyValues, -k        | specify keyValues (default: false) |
| --values, -V           | specify values (default: false)    |
| --unique, -u           | specify unique (default: false)    |
| --pretty, -P           | pretty format (default: false)     |
| --safeString value     | use safe string (value: on/off)    |
| --help                 | show help (default: false)         |

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

<a name="get-a-subscription"></a>

## Get a subscription

This command gets a subscription.

```console
ngsi get [common options] subscription [options]
```

### Options

| Options                | Description                     |
| ---------------------- | ------------------------------- |
| --id value, -i value   | specify id                      |
| --localTime            | localTime (default: false)      |
| --pretty, -P           | pretty format (default: false)  |
| --safeString value     | use safe string (value: on/off) |
| --help                 | show help (default: false)      |

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

| Options              | Description                        |
| -------------------- | ---------------------------------- |
| --id value, -i value | specify id                         |
| --localTime          | specify localTime (default: false) |
| --pretty, -P           | pretty format (default: false)  |
| --safeString value   | use safe string (value: on/off)    |
| --help               | show help (default: false)         |

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
