# hget - NGSI command

This command gets an entity, an attribute, multiple attributes, a subscription or a registration.

-   [Get hstory of an attribute](#get-hstory-of-an-attribute)
-   [Get history of attributes](#get-history-of-attributes)
-   [List of all the entity id](#list-of-all-the-entity-id)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="get-hstory-of-an-attribute"></a>

## Get hstory of an attribute

This command gets hstory of an attribute.

```console
ngsi hget [command options] attr [options]
```

### Options

| Options                | Description                                                                    |
| ---------------------- | ------------------------------------------------------------------------------ |
| --type value, -t value | Entity type                                                                    |
| --id value, -i value   | Entity id                                                                      |
| --attr value           | attribute name                                                                 |
| --sameType             | same type (default: false)                                                     |
| --nTypes               | nTypes (default: false)                                                        |
| --aggrMethod value     | aggregation method (max, min, sum, sum, occur)                                 |
| --aggrPeriod value     | aggregation period or resolution of the aggregated data to be retrieved        |
| --fromDate value       | starting date from which data should be retrieved                              |
| --toDate value         | final date until which data should be retrieved                                |
| --lastN value          | number of data entries to retrieve since the final date backwards (default: 0) |
| --hLimit value         | maximum number of data entries to retrieve (default: 0)                        |
| --hOffset value        | offset to be applied to data entries to be retrieved (default: 0)              |
| --georel value         | georel                                                                         |
| --geometry value       | geometry                                                                       |
| --coords value         | coords                                                                         |
| --value                | values only (default: false)                                                   |
| --pretty, -P           | pretty format (default: false)                                                 |
| --help                 | show help (default: false)                                                     |

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

| Options                | Description                                                                    |
| ---------------------- | ------------------------------------------------------------------------------ |
| --type value, -t value | Entity Type                                                                    |
| --id value, -i value   | Entity id                                                                      |
| --attrs value          | attributes                                                                     |
| --sameType             | same type (default: false)                                                     |
| --aggrMethod value     | aggregation method (max, min, sum, sum, occur)                                 |
| --aggrPeriod value     | aggregation period or resolution of the aggregated data to be retrieved        |
| --fromDate value       | starting date from which data should be retrieved                              |
| --toDate value         | final date until which data should be retrieved                                |
| --lastN value          | number of data entries to retrieve since the final date backwards (default: 0) |
| --hLimit value         | maximum number of data entries to retrieve (default: 0)                        |
| --hOffset value        | offset to be applied to data entries to be retrieved (default: 0)              |
| --georel value         | georel                                                                         |
| --geometry value       | geometry                                                                       |
| --coords value         | coords                                                                         |
| --value                | values only (default: false)                                                   |
| --pretty, -P           | pretty format (default: false)                                                 |
| --help                 | show help (default: false)                                                     |

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

| Options                | Description                                                       |
| ---------------------- | ----------------------------------------------------------------- |
| --type value, -t value | Entity Type                                                       |
| --fromDate value       | starting date from which data should be retrieved                 |
| --toDate value         | final date until which data should be retrieved                   |
| --hLimit value         | maximum number of data entries to retrieve (default: 0)           |
| --hOffset value        | offset to be applied to data entries to be retrieved (default: 0) |
| --pretty, -P           | pretty format (default: false)                                    |
| --help                 | show help (default: false)                                        |

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
