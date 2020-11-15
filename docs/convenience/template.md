# template - Convenience command

This command generates a json-style query text for subscription or registration.

-   [Subscription](#subscription)
-   [Registration](#registration)

## Subscription

This command generates a json-style query text to create a subscription and print it to stdout.

```
ngsi template subscription [options]
```

### Options

| Options                   | Description                                                |
| ------------------------- | ---------------------------------------------------------- |
| --ngsiType value          | specify NGSI type: v2 or ld                                |
| --data value, -d value    | specify data                                               |
| --uri value               | specify url or uri                                         |
| --url value, -u value     | specify url to be invoked when a notification is generated |
| --expires value, -e value | specify expires                                            |
| --throttling value        | specify throttling (default: 0)                            |
| --keyValues, -k           | specify keyValues (default: false)                         |
| --query value, -q value   | specify query                                              |
| --link value, -L value    | specify @context                                           |
| --nAttrs value            | specify attributes to be notified                          |
| --wAttrs value            | specify watched attributes                                 |
| --description value       | specify description                                        |
| --get                     | specify get (default: false)                               |
| --status value            | specify status                                             |
| --subjectId value         | specify subjectId                                          |
| --idPattern value         | specify idPattern                                          |
| --type value, -t value    | specify Entity Type                                        |
| --typePattern value       | specify typePattern                                        |
| --mq value, -m value      | specify mq                                                 |
| --georel value            | specify georel                                             |
| --geometry value          | specify geometry                                           |
| --coords value            | specify coords                                             |
| --headers value           | specify headers                                            |
| --qs value                | specify qs                                                 |
| --method value            | specify method                                             |
| --payload value           | specify payload                                            |
| --metadata value          | specify metadata                                           |
| --exceptAttrs value       | specify exceptAttrs                                        |
| --attrsFormat value       | specify attrsFormat                                        |
| --help                    | show help (default: false)                                 |

### Example for NGSI-LD

#### Reuqest:

```
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

```
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

```
ngsi template subscription --ngsiType ld \
  --description "LD Notify me of low stock in Store 002" \
  --type Shelf \
  --wAttrs "numberOfItems" \
  --query "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store002" \
  --nAttrs "numberOfItems,stocks,locatedIn" \
  --uri "http://tutorial:3000/subscription/low-stock-store002"
```

#### Respose:

```
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

```
$ ngsi template subscription --ngsiType v2 --idPattern ".*" --type Sensor --wAttrs temperature --nAttrs temperature --url http://192.168.0.1/ --expires 1day
```

#### Response:

```
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

```
ngsi template registration
```

### Options

| Options                    | Description                            |
| -------------------------- | -------------------------------------- |
| --id value, -i value       | specify id                             |
| --type value, -t value     | specify Entity Type                    |
| --attrs value              | specify attrs                          |
| --provider value, -p value | specify URL of context provider/source |
| --description value        | specify description                    | 
| --help                     | show help (default: false)             |

### Example for NGSI-LD

#### Request:

```
ngsi template registration --ngsiType ld \
  --type Building \
  --id urn:ngsi-ld:Building:store001 \
  --attrs tweets \
  --provider "http://context-provider:3000/static/tweets" \
  --description "ContextSourceRegistration"
```

#### Response:

```
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

#### Reuqest:

```
ngsi template subscription --ngsiType ld \
  --description "LD Notify me of low stock in Store 002" \
  --type Shelf \
  --wAttrs "numberOfItems" \
  --query "numberOfItems<10;locatedIn==urn:ngsi-ld:Building:store002" \
  --nAttrs "numberOfItems,stocks,locatedIn" \
  --uri "http://tutorial:3000/subscription/low-stock-store002" > subscription.json
```

#### Respose:

```
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

#### Response:

