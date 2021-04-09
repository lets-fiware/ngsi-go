# NGSI Go tutorial for STH-Comet

## Get settings of alias

```console
ngsi server get --host comet
```

```json
serverType comet
serverHost http://sth-comet:8666
FIWARE-Service openiot
FIWARE-ServicePath /
```

## Print version

```
ngsi version --host comet
```

```json
{"version":"2.8.0-next"}
```

## Raw data consumption

### Filtering by offset and limit

```console
ngsi hget --host comet \
attr --hLimit 3 --hOffset 0 --type Thing --id device001 --attr A1 --pretty
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

### Filtering by number of last entries

```console
ngsi hget --host comet \
attr -lastN 3 --type Thing --id device001 --attr A1 --pretty
```

```json
{
  "type": "StructuredValue",
  "value": [
    {
      "recvTime": "2016-09-15T01:57:00.000Z",
      "attrType": "Number",
      "attrValue": 2998
    },
    {
      "recvTime": "2016-09-15T01:58:00.000Z",
      "attrType": "Number",
      "attrValue": 2999
    },
    {
      "recvTime": "2016-09-15T01:59:00.000Z",
      "attrType": "Number",
      "attrValue": 3000
    }
  ]
}
```
## Aggregated data consumption by aggregation method and resolution

### Filtering by aggrMethod, aggrPeriod

```console
ngsi hget --host comet attr \
--aggrMethod sum --aggrPeriod day --type Thing --id device001 --attr A1 --pretty
```

```json
{
  "type": "StructuredValue",
  "value": [
    {
      "_id": {
        "attrName": "A1",
        "origin": "2016-09-01T00:00:00.000Z",
        "resolution": "day"
      },
      "points": [
        {
          "offset": 13,
          "samples": 1440,
          "sum": 1037520
        },
        {
          "offset": 14,
          "samples": 1440,
          "sum": 3111120
        },
        {
          "offset": 15,
          "samples": 120,
          "sum": 352860
        }
      ]
    }
  ]
}
```

### Filtering by max, day

```console
ngsi hget --host comet attr \
--attr A1 --aggrMethod max --aggrPeriod day --type Thing --id device001 --pretty
```

```json
{
  "type": "StructuredValue",
  "value": [
    {
      "_id": {
        "attrName": "A1",
        "origin": "2016-09-01T00:00:00.000Z",
        "resolution": "day"
      },
      "points": [
        {
          "offset": 13,
          "samples": 1440,
          "max": 1440
        },
        {
          "offset": 14,
          "samples": 1440,
          "max": 2880
        },
        {
          "offset": 15,
          "samples": 120,
          "max": 3000
        }
      ]
    }
  ]
}
```

### Filtering by min, day

```console
ngsi hget --host comet attr --attr A1 --aggrMethod min --aggrPeriod day --type Thing --id device001 --pretty
```

```json
{
  "type": "StructuredValue",
  "value": [
    {
      "_id": {
        "attrName": "A1",
        "origin": "2016-09-01T00:00:00.000Z",
        "resolution": "day"
      },
      "points": [
        {
          "offset": 13,
          "samples": 1440,
          "min": 1
        },
        {
          "offset": 14,
          "samples": 1440,
          "min": 1441
        },
        {
          "offset": 15,
          "samples": 120,
          "min": 2881
        }
      ]
    }
  ]
}
```

## Deleting historical raw and aggregated time series context information

### Deleting all the data associated to certain service and service path

```
ngsi hdelete --host comet entities--run
```

### Deleting all the data associated to certain entity, service and service path

```
ngsi hdelete --host comet entity --type Thing --id device001 --run
```

### Deleting all the data associated to certain attribute of certain entity, service and service path

```
ngsi hdelete --host comet attr --type Thing --id device001 --attr A1 --run
```
