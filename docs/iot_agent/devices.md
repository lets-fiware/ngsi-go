# devices - IoT Agent command

This command allows you to list, create, update and delete device entry for IoT Agent.

-   [List all devices](#list-all-devices)
-   [Create a device](#create-a-device)
-   [Get a device](#create-a-get-device)
-   [Update a device](#update-a-device)
-   [Delete a device](#delete-a-device)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="list-all-device"></a>

## List all devices

This command lists all devices.

```console
ngsi devices [command options] list [options]
```

### Options

| Options          | Description                                                             |
| ---------------- | ----------------------------------------------------------------------- |
| --limit value    | maximum number of devices (default: 0)                                  |
| --offset value   | offset to skip a given number of elements at the beginning (default: 0) |
| --detailed value | all device information (on) or only name (off) (default: "off")         |
| --entity value   | get a device from entity name                                           |
| --protocol value | get devices with this protocol                                          |
| --pretty, -P     | pretty format (default: false)                                          |
| --help           | show help (default: false)                                              |

### Examples

#### Request:

```console
ngsi devices list --pretty
```

```json
{
  "count": 1,
  "devices": [
    {
      "device_id": "sensor001",
      "service": "openiot",
      "service_path": "/",
      "entity_name": "urn:ngsi-ld:WeatherObserved:sensor001",
      "entity_type": "Sensor",
      "transport": "HTTP",
      "attributes": [
        {
          "object_id": "d",
          "name": "dateObserved",
          "type": "DateTime"
        },
        {
          "object_id": "t",
          "name": "temperature",
          "type": "Number"
        },
        {
          "object_id": "h",
          "name": "relativeHumidity",
          "type": "Number"
        },
        {
          "object_id": "p",
          "name": "atmosphericPressure",
          "type": "Number"
        }
      ],
      "lazy": [],
      "commands": [],
      "static_attributes": [
        {
          "name": "location",
          "type": "geo:json",
          "value": {
            "type": "Point",
            "coordinates": [
              139.7671,
              35.68117
            ]
          }
        }
      ],
      "explicitAttrs": false
    }
  ]
}
```

<a name="create-a-device"></a>

## Create a device

This command creates a device.

```console
ngsi devices [command options] create [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --data value, -d value | data body (payload)        |
| --help                 | show help (default: false) |

### Examples

#### Request:

```console
ngsi devices create --data \
'{
 "devices": [
   {
     "device_id":   "sensor001",
     "entity_name": "urn:ngsi-ld:WeatherObserved:sensor001",
     "entity_type": "Sensor",
     "timezone":    "Asia/Tokyo",
     "attributes": [
       { "object_id": "d", "name": "dateObserved", "type": "DateTime" },
       { "object_id": "t", "name": "temperature", "type": "Number" },
       { "object_id": "h", "name": "relativeHumidity", "type": "Number" },
       { "object_id": "p", "name": "atmosphericPressure", "type": "Number" }
     ],
     "static_attributes": [
       { "name":"location", "type": "geo:json", "value" : { "type": "Point", "coordinates" : [ 139.7671, 35.68117 ] } }
     ]
   }
 ]
}'
```

<a name="get-a-device"></a>

## Get a device

This command gets a device.

```console
ngsi devices [command options] get [options]
```

### Options

| Options              | Description                    |
| -------------------- | ------------------------------ |
| --id value, -i value | device id                      |
| --pretty, -P         | pretty format (default: false) |
| --help               | show help (default: false)     |

### Examples

#### Request:

```console
ngsi devices get --id sensor001 --pretty
```

```json
{
  "device_id": "sensor001",
  "service": "openiot",
  "service_path": "/",
  "entity_name": "urn:ngsi-ld:WeatherObserved:sensor001",
  "entity_type": "Sensor",
  "transport": "HTTP",
  "attributes": [
    {
      "object_id": "d",
      "name": "dateObserved",
      "type": "DateTime"
    },
    {
      "object_id": "t",
      "name": "temperature",
      "type": "Number"
    },
    {
      "object_id": "h",
      "name": "relativeHumidity",
      "type": "Number"
    },
    {
      "object_id": "p",
      "name": "atmosphericPressure",
      "type": "Number"
    }
  ],
  "lazy": [],
  "commands": [],
  "static_attributes": [
    {
      "name": "location",
      "type": "geo:json",
      "value": {
        "type": "Point",
        "coordinates": [
          139.7671,
          35.68117
        ]
      }
    }
  ],
  "explicitAttrs": false
}
```

<a name="update-a-device"></a>

## Update a device

This command updates a device.

```console
ngsi devices [command options] update [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | device id                  |
| --data value, -d value | data body (payload)        |
| --help                 | show help (default: false) |

### Examples

#### Request:

```console
ngsi devices update \
--id sensor003 \
--data '{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}'
```

<a name="delete-a-device"></a>

## Delete a device

This command deletes a device.

```console
ngsi devices [command options] delete [options]
```

### Options

| Options              | Description                |
| -------------------- | -------------------------- |
| --id value, -i value | device id                  |
| --help               | show help (default: false) |

### Examples

#### Request:

```console
ngsi devices delete --id sensor001
```
