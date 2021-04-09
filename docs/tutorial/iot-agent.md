# NGSI Go tutorial for IoT Agent

## Get settings of alias

```console
ngsi server list --host iota
```

```json
serverType iota
serverHost http://iot-agent:4041
FIWARE-Service openiot
FIWARE-ServicePath /
```

## Print version

```console
ngsi version --host iota --pretty
```

```json
{
  "libVersion": "2.14.0",
  "port": "4041",
  "baseRoot": "/",
  "version": "1.15.0"
}
```

## Create service

```console
ngsi services create \
--apikey 4jggokgpepnvsb2uv4s40d59ov \
--cbroker http://orion:1026 \
--type Event \
--resource /iot/d
```

## List services

```console
ngsi services list --pretty
```

```json
{
  "count": 1,
  "services": [
    {
      "commands": [],
      "lazy": [],
      "attributes": [],
      "_id": "601f442fd679e88d3e8c01a3",
      "resource": "/iot/d",
      "apikey": "apikey",
      "service": "openiot",
      "subservice": "/",
      "__v": 0,
      "static_attributes": [],
      "internal_attributes": [],
      "entity_type": "Event"
    }
  ]
}
```

## Update service

```console
ngsi services update \
--resource /iot/d \
--apikey 4jggokgpepnvsb2uv4s40d59ov \
--type Device
```

## List services

```console
ngsi services list --pretty
```

```json
{
  "count": 1,
  "services": [
    {
      "commands": [],
      "lazy": [],
      "attributes": [],
      "_id": "601f44ddd679e860a98c01a4",
      "resource": "/iot/d",
      "apikey": "4jggokgpepnvsb2uv4s40d59ov",
      "service": "openiot",
      "subservice": "/",
      "__v": 0,
      "static_attributes": [],
      "internal_attributes": [],
      "entity_type": "Device"
    }
  ]
}
```

## Create device

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

## List devices

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

## Get device

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

## Update device

```console
ngsi devices update \
--id sensor001 \
--data '{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor111"}'
```

## Get device

```console
ngsi devices get --id sensor001 --pretty
```

```json
{
  "device_id": "sensor001",
  "service": "openiot",
  "service_path": "/",
  "entity_name": "urn:ngsi-ld:WeatherObserved:sensor111",
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

## Delete device

```console
ngsi devices delete --id sensor001
```

## Get device

```console
ngsi devices get --id sensor001 --pretty
```

```console
idasDevicesGet005 404 Not Found {"name":"DEVICE_NOT_FOUND","message":"No device was found with id:sensor001"}
```

## Delete service

```console
ngsi services delete --resource /iot/d
``` 

## List services

```console
ngsi services list --pretty
```

```json
{
  "count": 0,
  "services": []
}
```
