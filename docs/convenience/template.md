# template - Convenience command

This command generates a json-style query text for subscription or registration.

-   [Subscription](#subscription)
-   [Registration](#registration)

### Options

| Options                   | Description                 |
| ------------------------- | --------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE |
| --service VALUE, -s VALUE | FIWARE Service VALUE        |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE    |
| --link VALUE, -L VALUE    | @context VALUE (LD)         |
| --help                    | show help (default: true)   |

<a name="subscription"></a>

## Subscription

This command generates a json-style query text to create a subscription and print it to stdout.

```console
ngsi template subscription [options]
```

### Options

| Options                   | Description                                            |
| ------------------------- | ------------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE                            |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                               |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                    |
| --ngsiType VALUE          | NGSI type: v2 or ld                                    |
| --data VALUE, -d VALUE    | subscription data                                      |
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
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                                    |
| --status VALUE            | status                                                 |
| --headers VALUE           | headers (v2)                                           |
| --qs VALUE                | qs (v2)                                                |
| --method VALUE            | method (v2)                                            |
| --payload VALUE           | payload (v2)                                           |
| --metadata VALUE          | metadata (v2)                                          |
| --exceptAttrs VALUE       | exceptAttrs (v2)                                       |
| --attrsFormat VALUE       | attrsFormat (v2)                                       |
| --pretty, -P              | pretty format (default: false)                         |
| --help                    | show help (default: true)                              |

### Example for NGSI-LD

#### Reuqest:

```console
ngsi template subscription --ngsiType ld
  --type Shelf \
  --query "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store001" \
  --wAttrs "numberOfItems" \
  --nAttrs "numberOfItems,stocks,locatedIn" --keyValues \
  --uri "http://tutorial:3000/subscription/low-stock-store001" \
  --link tutorial \
  --description "Notify me of low stock in Store 001"
```

#### Response:

```json
{
  "subscriptionName": "subscription name",
  "description": "Notify me of low stock in Store 001",
  "type": "Subscription",
  "entities": [
    {
      "type": "Shelf"
    }
  ],
  "watchedAttributes": [
    "numberOfItems"
  ],
  "timeInterval": 0,
  "q": "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store001",
  "csf": "",
  "isActive": true,
  "expiresAt": "2099-12-31T14:00:00Z",
  "notification": {
    "attributes": [
      "numberOfItems",
      "stocks",
      "locatedIn"
    ],
    "format": "keyValues",
    "endpoint": {
      "uri": "http://tutorial:3000/subscription/low-stock-store001",
      "accept": "application/ld+json"
    }
  },
  "@context": "http://context-provider:3000/data-models/ngsi-context.jsonld"
}
```

#### Reuqest:

```console
ngsi template subscription --ngsiType ld \
  --description "LD Notify me of low stock in Store 002" \
  --type Shelf \
  --wAttrs "numberOfItems" \
  --query "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store002" \
  --nAttrs "numberOfItems,stocks,locatedIn" \
  --uri "http://tutorial:3000/subscription/low-stock-store002"
```

#### Respose:

```json
{
  "description": "LD Notify me of low stock in Store 002",
  "entities": [
    {
      "type": "Shelf"
    }
  ],
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
  "q": "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store002",
  "type": "Subscription",
  "watchedAttributes": [
    "numberOfItems"
  ]
}
```

### Example for NGSIv2

#### Request:

```console
ngsi template subscription --ngsiType v2 --idPattern ".*" --type Sensor --wAttrs temperature --nAttrs temperature --url http://192.168.0.1/ --expires 1day
```

#### Response:

```json
{
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
    "http": {
      "url": "http://192.168.0.1/"
    },
    "attrs": [
      "temperature"
    ]
  },
  "expires": "2020-11-09T11:43:17.510Z"
}
```

<a name="registration"></a>

## Registration

This command will generate a json-style query to create a registration and print it to stdout.

```console
ngsi template registration
```

### Options

| Options                    | Description                                  |
| -------------------------- | -------------------------------------------- |
| --host VALUE, -h VALUE     | broker or server host VALUE                  |
| --service VALUE, -s VALUE  | FIWARE Service VALUE                         |
| --path VALUE, -p VALUE     | FIWARE ServicePath VALUE                     |
| --link VALUE, -L VALUE     | @context VALUE (LD)                          |
| --ngsiType VALUE           | NGSI type: v2 or ld                          |
| --data VALUE, -d VALUE     | registration data                            |
| --description VALUE        | description                                  |
| --type VALUE, -t VALUE     | entity type                                  |
| --providedId VALUE         | providedId                                   |
| --idPattern VALUE          | idPattern                                    |
| --properties VALUE         | properties (LD)                              |
| --relationships VALUE      | relationships (LD)                           |
| --expires VALUE, -e VALUE  | expires                                      |
| --provider VALUE, -p VALUE | Url of context provider/source               |
| --attrs VALUE              | attributes                                   |
| --legacy                   | legacy forwarding mode (V2) (default: false) |
| --forwardingMode VALUE     | forwarding mode (V2)                         |
| --status VALUE             | status                                       |
| --context VLAUE, -C VLAUE  | @context VLAUE (LD)                          |
| --pretty, -P               | pretty format (default: false)               |
| --help                     | show help (default: true)                    |

### Example for NGSI-LD

#### Request:

```console
ngsi template registration --ngsiType ld \
  --type Building \
  --id urn:ngsi-ld:Building:store001 \
  --attrs tweets \
  --provider "http://context-provider:3000/static/tweets" \
  --description "ContextSourceRegistration"
```

#### Response:

```json
{
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
  ],
  "type": "ContextSourceRegistration"
}
```

### Example for NGSIv2

#### Request:

```console
ngsi template registration \
  --ngsiType v2 \
  --description "sensor source" \
  --attrs temperature,pressure,humidity \
  --providedId urn:ngsi-ld:Device:device001 \
  --type Device \
  --provider http://raspi
```

```json
{
  "description": "sensor source",
  "dataProvided": {
    "entities": [
      {
        "id": "urn:ngsi-ld:Device:device001",
        "type": "Device"
      }
    ],
    "attrs": [
      "temperature",
      "pressure",
      "humidity"
    ]
  },
  "provider": {
    "http": {
      "url": "http://raspi"
    }
  }
}
```

```console
5fb9dcc4a723657d763c6317
```
