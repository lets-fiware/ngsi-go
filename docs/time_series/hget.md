# hget - NGSI command

This command gets an entity, an attribute, multiple attributes, a subscription or a registration.

-   [Get hstory of an attribute](#get-hstory-of-an-attribute)
-   [Get history of attributes](#get-history-of-attributes)
-   [List of all the entity id](#list-of-all-the-entity-id)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="get-hstory-of-an-attribute"></a>

## Get hstory of an attribute

This command gets hstory of an attribute.

```console
ngsi hget [command options] attr [options]
```

### Options

| Options                   | Description                                                             |
| ------------------------- | ----------------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                                  |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                                    |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                                |
| --type VALUE, -t VALUE    | Entity Type                                                             |
| --id VALUE, -i VALUE      | id                                                                      |
| --attr VALUE              | attribute name                                                          |
| --sameType                | same type (default: false)                                              |
| --nTypes                  | nTypes (default: false)                                                 |
| --aggrMethod VALUE        | aggregation method (max, min, sum, sum, occur)                          |
| --aggrPeriod VALUE        | aggregation period or resolution of the aggregated data to be retrieved |
| --fromDate VALUE          | starting date from which data should be retrieved                       |
| --toDate VALUE            | final date until which data should be retrieved                         |
| --lastN VALUE             | number of data entries to retrieve since the final date backwards       |
| --hLimit VALUE            | maximum number of data entries to retrieve                              |
| --hOffset VALUE           | offset to be applied to data entries to be retrieved                    |
| --georel VALUE            | georel                                                                  |
| --geometry VALUE          | geometry                                                                |
| --coords VALUE            | coords                                                                  |
| --value                   | values only (default: false)                                            |
| --pretty, -P              | pretty format (default: false)                                          |
| --safeString VALUE        | use safe string (VALUE: on/off)                                         |
| --help                    | show help (default: true)                                               |

### Examples

#### Request:

```console
ngsi hget --host comet --service openiot --path / \
attr --type Thing --id device001 --attr A1 \
--hLimit 3 --hOffset 0 --pretty
```

```json
{
  "type": "StructuredValue",
  "value": [
    {
      "recvTime": "2016-09-13T00:00:00.000Z",
      "attrType": "Number",
      "attrValue": 1
    },
    {
      "recvTime": "2016-09-13T00:01:00.000Z",
      "attrType": "Number",
      "attrValue": 2
    },
    {
      "recvTime": "2016-09-13T00:02:00.000Z",
      "attrType": "Number",
      "attrValue": 3
    }
  ]
}
```

#### Request:

```console
ngsi hget --host quantumleap \
attr --id device001 --attr A1 --fromDate -5years -hLimit 3
```

```json
{
  "attr": "A1",
  "entityId": "device001",
  "index": [
    "2016-09-13T00:00:00.000+00:00",
    "2016-09-13T00:01:00.000+00:00",
    "2016-09-13T00:02:00.000+00:00"
  ],
  "values": [
    1.0,
    2.0,
    3.0
  ]
}
```

<a name="get-history-of-attributes"></a>

## Get history of attributes

This command gets history of attributes.

```console
ngsi hget [command options] attributes [options]
```

### Options

| Options                   | Description                                                             |
| ------------------------- | ----------------------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                                  |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                                    |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                                |
| --type VALUE, -t VALUE    | Entity Type                                                             |
| --id VALUE, -i VALUE      | id                                                                      |
| --attrs VALUE             | attributes                                                              |
| --sameType                | same type (default: false)                                              |
| --nTypes                  | nTypes (default: false)                                                 |
| --aggrMethod VALUE        | aggregation method (max, min, sum, sum, occur)                          |
| --aggrPeriod VALUE        | aggregation period or resolution of the aggregated data to be retrieved |
| --fromDate VALUE          | starting date from which data should be retrieved                       |
| --toDate VALUE            | final date until which data should be retrieved                         |
| --lastN VALUE             | number of data entries to retrieve since the final date backwards       |
| --hLimit VALUE            | maximum number of data entries to retrieve                              |
| --hOffset VALUE           | offset to be applied to data entries to be retrieved                    |
| --georel VALUE            | georel                                                                  |
| --geometry VALUE          | geometry                                                                |
| --coords VALUE            | coords                                                                  |
| --value                   | values only (default: false)                                            |
| --pretty, -P              | pretty format (default: false)                                          |
| --safeString VALUE        | use safe string (VALUE: on/off)                                         |
| --help                    | show help (default: true)                                               |

### Examples

#### Request:

```console
ngsi hget --host quantumleap \
attrs --id device001 --attrs A1,A2 --fromDate -5years -hLimit 3
```

```json
{
  "attributes": [
    {
      "attr": "A1",
      "values": [
        1.0,
        2.0,
        3.0
      ]
    },
    {
      "attr": "A2",
      "values": [
        2.0,
        3.0,
        4.0
      ]
    }
  ],
  "entityId": "device001",
  "index": [
    "2016-09-13T00:00:00.000+00:00",
    "2016-09-13T00:01:00.000+00:00",
    "2016-09-13T00:02:00.000+00:00"
  ]
}
```

<a name="list-of-all-the-entity-id"></a>

## List of all the entity id

This command lists of all the entity id.

```console
ngsi hget [common options] entities [options]
```

### Options

| Options                   | Description                                          |
| ------------------------- | ---------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)               |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                 |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                             |
| --type VALUE, -t VALUE    | Entity Type                                          |
| --fromDate VALUE          | starting date from which data should be retrieved    |
| --toDate VALUE            | final date until which data should be retrieved      |
| --hLimit VALUE            | maximum number of data entries to retrieve           |
| --hOffset VALUE           | offset to be applied to data entries to be retrieved |
| --pretty, -P              | pretty format (default: false)                       |
| --safeString VALUE        | use safe string (VALUE: on/off)                      |
| --help                    | show help (default: true)                            |

### Examples

#### Request:

```console
hget --host quantumleap entities
```

```json
[
  {
    "id": "Event001",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  },
  {
    "id": "Event002",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  },
  {
    "id": "device001",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Thing"
  },
  {
    "id": "device002",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Thing"
  }
]
```
