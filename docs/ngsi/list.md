# list - NGSI command

This command lists types, entities, subscriptions or registrations.

-   [List multiple types](#list-multiple-types)
-   [List multiple entities](#list-multiple-entities)
-   [List multiple subscriptions](#list-multiple-subscriptions)
-   [List multiple registrations](#list-multiple-registrations)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="list-multiple-types"/>

## List multiple types

This command lists types.

```bash
ngsi list [common options] types [options]
```

### Options

| Options       | Description                  |
| ------------- | ---------------------------- |
| --verbose, -v | verbose (default: false)     |
| --json, -j    | JSON format (default: false) |
| --help        | show help (default: false)   |

### Examples

#### Request:

```bash
$ ngsi list types
InventoryItem
Product
Shelf
Store
```

#### Request:

```bash
$ ngsi list types --json
["InventoryItem","Product","Shelf","Store"]
```

<a name="list-multiple-entities"/>

## List multiple entities

This command lists multiple entities.

```bash
ngsi list [common options] entities [options]
```

### Options

| Options                   | Description                                 |
| ------------------------- | ------------------------------------------- |
| --host value, -h value    | specify host or alias for source (Required) |
| --token value             | specify oauth token                         |
| --service value, -s value | specify FIWARE Service                      |
| --path value, -p value    | specify FIWARE ServicePath                  |
| --type value, -t value    | specify Entity Type                         |
| --idPattern value         | specify idPattern                           |
| --typePattern value       | specify typePattern                         |
| --query value, -q value   | specify query                               |
| --mq value, -m value      | specify mq                                  |
| --georel value            | specify georel                              |
| --geometry value          | specify geometry                            |
| --coords value            | specify coords                              |
| --attrs value             | specify attrs                               |
| --metadata value          | specify metadata                            |
| --orderBy value           | specify orderBy                             |
| --count, -C               | specify count (default: false)              |
| --keyValues, -k           | specify keyValues (default: false)          |
| --values, -V              | specify values (default: false)             |
| --unique, -u              | specify unique (default: false)             |
| --id value, -i value      | specify id                                  |
| --link value, -L value    | specify @context                            |
| --verbose, -v             | specify verbose (default: false)            |
| --lines, -1               | specify lines (default: false)              |
| --safeString value        | use safe string (value: on/off)             |
| --help                    | show help (default: false)                  |

### Example

#### Request:

```bash
$ $ ngsi list entities --type Product
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

```bash
$ $ ngsi list entities --type Product --count
14
```

#### Request:

```bash
$ ngsi list entities --type Product --idPattern '0{2}'
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

```bash
$ ngsi list entities --type Product --idPattern '1{2}'
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112

$ ngsi list entities --type Product --idPattern '1{2}' --count
3
```

#### Request:

```bash
$ ngsi list entities --type Product --idPattern '1{2}' --verbose | jq .
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

```bash
$ ngsi list entities --type Product --idPattern '1{2}' --verbose --keyValues | jq .
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

```bash
$ ngsi list entities --type Product --idPattern '1{2}' --count
3

$ ngsi list entities --type Product --idPattern '1{2}' | xargs -L 1 ngsi delete entity --id

$ ngsi list entities --type Product --idPattern '1{2}' --count
0
```

#### Request:

```bash
ngsi list entities -q "refProduct%==urn:ngsi-ld:Product:001" --attrs type
```

<a name="list-multiple-subscriptions"/>

## List multiple subscriptions

This command lists multiple subscriptions.

```bash
ngsi list [common options] subscriptions [options]
```

### Options

| Options                 | Description                        |
| ----------------------- | ---------------------------------- |
| --verbose, -v           | verbose (default: false)           |
| --json, -j              | JSON format (default: false)       |
| --status value          | specify status                     |
| --localTime             | specify localTime (default: false) |
| --query value, -q value | specify query                      |
| --items value, -i value | specify itmes                      |
| --safeString value      | use safe string (value: on/off)    |
| --help                  | show help (default: false)         |

### Examples for NGSI-LD

#### Request:

```bash
$ ngsi list subscriptions
urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce
urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf
```

#### Request:

```bash
$ ngsi list subscriptions --verbose
urn:ngsi-ld:Subscription:5f67fd65ef40bb66fe006dce Notify me of low stock in Store 001
urn:ngsi-ld:Subscription:5f680822ef40bb66fe006dcf LD Notify me of low stock in Store 002
```

#### Request:

```bash
$ ngsi list subscriptions --json | jq .
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

<a name="list-multiple-registrations"/>

## List multiple registrations

This command lists multiple registrations.

```bash
ngsi list [common options] registrations [options]
```

### Options

| Options            | Description                     |
| ------------------ | ------------------------------- |
| --verbose, -v      | verbose (default: false)        |
| --json, -j         | JSON format (default: false)    |
| --safeString value | use safe string (value: on/off) |
| --help             | show help (default: false)      |


### Examples for NGSI-LD

#### Request:

```bash
$ ngsi list registrations
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0
```

#### Request:

```bash
$ ngsi list registrations -v
urn:ngsi-ld:ContextSourceRegistration:5f6840e6ef40bb66fe006dd0 ContextSourceRegistration
```

#### Request:

```bash
$ ngsi list registrations -j | jq .
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
