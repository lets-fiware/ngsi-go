# template - Convenience command

This command generates a json-style query text for subscription or registration.

-   [Subscription](#subscription)
-   [Registration](#registration)

## Subscription

This command generates a json-style query text to create a subscription and print it to stdout.

```console
ngsi template subscription [options]
```

### Options for NGSI-LD

| Options                   | Description                                                    |
| ------------------------- | -------------------------------------------------------------- |
| --ngsiType value          | specify NGSI type: ld                                          |
| --data value, -d value    | specify data                                                   |
| --subscriptionId value    | specify subscription id (LD)                                   |
| --name value              | specify subscription name (LD)                                 |
| --description value       | specify description                                            |
| --entityId value          | specify entityId (LD)                                          |
| --idPattern value         | specify idPattern                                              |
| --type value, -t value    | specify Entity Type                                            |
| --typePattern value       | specify typePattern (v2)                                       |
| --wAttrs value            | specify watched attributes                                     |
| --timeInterval value      | specify time interval (LD) (default: 0)                        |
| --query value, -q value   | filtering by attribute value                                   |
| --mq value, -m value      | filtering by metadata (v2)                                     |
| --geometry value          | specify geometry                                               |
| --coords value            | specify coords                                                 |
| --georel value            | specify georel                                                 |
| --geoproperty value       | sprcify geoproperty (LD)                                       |
| --csf value               | specify context source filter (LD)                             |
| --active                  | specify active (LD) (default: false)                           |
| --inactive                | specify inactive (LD) (default: false)                         |
| --nAttrs value            | specify attributes to be notified                              |
| --keyValues, -k           | specify keyValues (default: false)                             |
| --uri value               | specify uri/url to be invoked when a notification is generated |
| --accept value            | specify accept header (json or ld+json)                        |
| --expires value, -e value | specify expires                                                |
| --throttling value        | specify throttling (default: 0)                                |
| --timeRel value           | specify temporal relationship (LD)                             |
| --timeAt value            | specify timeAt (LD)                                            |
| --endTimeAt value         | specify endTimeAt (LD)                                         |
| --timeProperty value      | specify timeProperty (LD)                                      |
| --link value, -L value    | specify @context (LD)                                          |
| --status value            | specify status                                                 |
| --headers value           | specify headers (v2)                                           |
| --qs value                | specify qs (v2)                                                |
| --method value            | specify method (v2)                                            |
| --payload value           | specify payload (v2)                                           |
| --metadata value          | specify metadata (v2)                                          |
| --exceptAttrs value       | specify exceptAttrs (v2)                                       |
| --attrsFormat value       | specify attrsFormat (v2)                                       |
| --help                    | show help (default: false)                                     |

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

## Registration

This command will generate a json-style query to create a registration and print it to stdout.

```console
ngsi template registration
```

### Options

| Options                    | Description                                     |
| -------------------------- | ----------------------------------------------- |
| --ngsiType value           | specify NGSI type: v2 or ld                     |
| --providedId value         | specify id                                      |
| --idPattern value          | specify idPattern                               |
| --type value, -t value     | specify Entity Type                             |
| --attrs value              | specify attrs                                   |
| --provider value, -p value | specify URL of context provider/source          |
| --description value        | specify description                             |
| --legacy                   | specify legacy forwarding mode (default: false) |
| --forwardingMode value     | specify forwarding mode                         |
| --expires value, -e value  | specify expires                                 |
| --status value             | specify status                                  |
| --help                     | show help (default: false)                      |

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

```
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

```
5fb9dcc4a723657d763c6317
```
