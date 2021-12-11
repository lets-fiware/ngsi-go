# create - NGSI command

-   [Create an entity](#create-an-entity)
-   [Create multiple entities](#create-multiple-entities)
-   [create temporal entity](#create-temporal-entity)
-   [Create a subscription](#create-a-subscription)
-   [Create a registration](#create-a-registration)
-   [Create a JSON-LD context](#create-a-json-ld-context)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="create-an-entity"></a>

## Create an entity

This command create an entity.

```console
ngsi create [command options] entity [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | entity data (required)                 |
| --keyValues, -K           | keyValues (default: false)             |
| --upsert                  | upsert (default: false)                |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi create entity \
--data ' {
      "id":"urn:ngsi-ld:Product:010",
      "type":"Product",
      "name":{"type":"Text", "value":"Lemonade"},
      "size":{"type":"Text", "value": "S"},
      "price":{"type":"Integer", "value": 99}
}'
```

#### Request:

```console
ngsi get entity --id urn:ngsi-ld:Product:010 --type Product
```

```json
{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade","metadata":{}},"price":{"type":"Integer","value":99,"metadata":{}},"size":{"type":"Text","value":"S","metadata":{}}}
```

#### Request:

```console
ngsi create entity --keyValues \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "Lemonade",
      "size": "S",
      "price": 99
}'
```

#### Request:

```console
ngsi get entity --id urn:ngsi-ld:Product:110 --type Product
```

```json
{"id":"urn:ngsi-ld:Product:110","type":"Product","name":{"type":"Text","value":"Lemonade","metadata":{}},"price":{"type":"Number","value":99,"metadata":{}},"size":{"type":"Text","value":"S","metadata":{}}}
```

<a name="create-multiple-entities"></a>

## Create multiple entities

This command create entities.

```console
ngsi create [common options] entities [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | entities data (required)               |
| --keyValues, -K           | keyValues (default: false)             |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi create entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:011",
    "type":"Product",
    "name":{"type":"Text", "value":"Brandy"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:012",
    "type":"Product",
    "name":{"type":"Text", "value":"Port"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1099}
  },
  {
    "id":"urn:ngsi-ld:Product:001",
    "type":"Product",
    "offerPrice":{"type":"Integer", "value": 89}
  }
]'
```

#### Request:

```console
ngsi create entities --keyValues \
--data '[
  {
    "id":"urn:ngsi-ld:Product:111",
    "type":"Product",
    "name": "Brandy",
    "size": "M",
    "price": 1199
  },
  {
    "id":"urn:ngsi-ld:Product:112",
    "type":"Product",
    "name":"Port",
    "size":"M",
    "price":1099
  },
  {
    "id":"urn:ngsi-ld:Product:101",
    "type":"Product",
    "offerPrice":89
  }
]'
```

<a name="create-temporal-entity"></a>

## Create a temporal entity

This command creates a temporal entity.

```
ngsi create [command options] tentity [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | temporal entity data (required)        |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

<a name="create-a-subscription"></a>

## Create a subscription

This command reads a query that the template command generated and creates a subscription.

```console
ngsi create [command options] subscription [options]
```

### Options

| Options                   | Description                                            |
| ------------------------- | ------------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                 |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                               |
| --data VALUE, -d VALUE    | subscription data                                      |
| --skipInitialNotification | skipInitialNotification (default: false)               |
| --subscriptionId VALUE    | subscription id (LD)                                   |
| --name VALUE              | subscription name (LD)                                 |
| --description VALUE       | description                                            |
| --entityId VALUE          | entity id                                              |
| --idPattern VALUE         | idPattern                                              |
| --type VALUE, -t VALUE    | entity type                                            |
| --typePattern VALUE       | typePattern (v2)                                       |
| --wAttrs VALUE            | watched attributes                                     |
| --timeInterval VALUE      | time interval (LD)                                     |
| --query VALUE, -q VALUE   | filtering by attribute value                           |
| --mq VALUE, -m VALUE      | filtering by metadata (v2)                             |
| --geometry VALUE          | geometry                                               |
| --coords VALUE            | coords                                                 |
| --georel VALUE            | georel                                                 |
| --geoproperty VALUE       | geoproperty (LD)                                       |
| --csf VALUE               | context source filter (LD)                             |
| --active                  | active (LD) (default: false)                           |
| --inactive                | inactive (LD) (default: false)                         |
| --nAttrs VALUE            | attributes to be notified                              |
| --keyValues, -K           | keyValues (default: false)                             |
| --uri VALUE, -u VALUE     | uri/url to be invoked when a notification is generated |
| --accept VALUE            | accept header (json or ld+json)                        |
| --expires VALUE, -e VALUE | expires                                                |
| --throttling VALUE        | throttling                                             |
| --timeRel VALUE           | temporal relationship (LD)                             |
| --timeAt VALUE            | timeAt (LD)                                            |
| --endTimeAt VALUE         | endTimeAt (LD)                                         |
| --timeProperty VALUE      | timeProperty (LD)                                      |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                                    |
| --status VALUE            | status                                                 |
| --headers VALUE           | headers (v2)                                           |
| --qs VALUE                | qs (v2)                                                |
| --method VALUE            | method (v2)                                            |
| --payload VALUE           | payload (v2)                                           |
| --metadata VALUE          | metadata (v2)                                          |
| --exceptAttrs VALUE       | exceptAttrs (v2)                                       |
| --attrsFormat VALUE       | attrsFormat (v2)                                       |
| --safeString VALUE        | use safe string (VALUE: on/off)                        |
| --raw                     | handle raw data (default: false)                       |
| --help                    | show help (default: true)                              |

### Example for NGSIv2

#### Request:

```console
ngsi create subscription --idPattern ".*" --type Sensor \
--wAttrs temperature --nAttrs temperature \
--url http://192.168.0.1/ --expires 1day
```

```text
5fa7988a627088ba9b91b1c1
```

#### Request:

```console
ngsi get subscription --id 5fa7988a627088ba9b91b1c1 --pretty
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
      "temperature"
    ],
    "attrsFormat": "normalized"
  },
  "expires": "2020-11-09T07:04:42.000Z",
  "status": "active"
}
```

### Example for NGSI-LD

#### Request:

```console
ngsi create subscription --data @subscription.json
```

```text
urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce
```

#### Request:

```console
ngsi create subscription \
  --link https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld \
  --data @subscription.json
```

```text
urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf
```

<a name="create-a-registration"></a>

## Create a registration

This command reads a query that the template command generated and creates a registration.

```console
ngsi create [command options] registration [options]
```

### Options

| Options                    | Description                                  |
| -------------------------- | -------------------------------------------- |
| --host VALUE, -h VALUE     | broker or server host VALUE (required)       |
| --service VALUE, -s VALUE  | FIWARE Service VALUE                         |
| --path VALUE, -p VALUE     | FIWARE ServicePath VALUE                     |
| --data VALUE, -d VALUE     | registration data                            |
| --link VALUE, -L VALUE     | @context VALUE (LD)                          |
| --context VLAUE, -C VLAUE  | @context VLAUE (LD)                          |
| --providedId VALUE         | providedId                                   |
| --idPattern VALUE          | idPattern                                    |
| --type VALUE, -t VALUE     | entity type                                  |
| --attrs VALUE              | attributes                                   |
| --provider VALUE, -p VALUE | Url of context provider/source               |
| --description VALUE        | description                                  |
| --legacy                   | legacy forwarding mode (V2) (default: false) |
| --forwardingMode VALUE     | forwarding mode (V2)                         |
| --expires VALUE, -e VALUE  | expires                                      |
| --status VALUE             | status                                       |
| --safeString VALUE         | use safe string (VALUE: on/off)              |
| --help                     | show help (default: true)                    |

### Example for NGSI-LD

#### Request:

```console
ngsi create registration --data @registration.json
```

```text
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0
```

<a name="create-a-json-ld-context"></a>

## Create a JSON-LD context

This command create a JSON-LD context.

```console
ngsi create [command options] ldContext [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | jsonldContexts data (LD) (required)    |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi create --host orion-ld ldContext \
  --data '["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]'
```

#### Response:

```console
d42e7ffe-ed21-11eb-bc92-0242c0a8a010
```
