# create - NGSI command

-   [Create an entity](#create-an-entity)
-   [Create multiple entities](#create-multiple-entities)
-   [Create a subscription](#create-a-subscription)
-   [Create a registration](#create-a-registration)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="create-an-entity"/>

## Create an entity

This command create an entity.

```console
ngsi create [command options] entity [options]
```

### Options

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --data value, -d value    | specify data                       |
| --keyValues, -k           | specify keyValues (default: false) |
| --upsert                  | specify upsert (default: false)    |
| --link value, -L value    | specify @context                   |
| --help                    | show help (default: false)         |

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

<a name="create-multiple-entities"/>

## Create multiple entities

This command create entities.

```console
ngsi create [common options] entities [options]
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

<a name="create-a-subscription"/>

## Create a subscription

This command reads a query that the template command generated and creates a subscription.

```console
ngsi create [command options] subscription [options]
```

### Options

| Options                   | Description                                                |
| ------------------------- | ---------------------------------------------------------- |
| --data value, -d value    | specify data                                               |
| --skipInitialNotification | specify skipInitialNotification (default: false)           |
| --ngsiType value          | specify NGSI type: v2 or ld                                |
| --uri value               | specify URL or URI                                         |
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
| --typePattern value       | specify typePatternA                                       |
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
ngsi get subscription --id 5fa7988a627088ba9b91b1c1 | jq .
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

<a name="create-a-registration"/>

## Create a registration

This command reads a query that the template command generated and creates a registration.

```console
ngsi create [command options] registration [options]
```

### Options

| Options                    | Description                                     |
| -------------------------- | ----------------------------------------------- |
| --data value, -d value     | specify data                                    |
| --link value, -L value     | specify @context                                |
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
| --safeString value         | use safe string (value: on/off)                 |
| --help                     | show help (default: false)                      |

### Example for NGSI-LD

#### Request:

```console
ngsi create registration --data @registration.json
```

```text
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0
```
