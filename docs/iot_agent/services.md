# services - IoT Agent command

This command allows you to list, create, update and delete service entry for IoT Agent.

-   [List configuration groups](#list-configuration-group)
-   [Create a configuration group](#create-a-configuration-group)
-   [Update a configuration group](#update-a-configuration-group)
-   [Delete a configuration group](#delete-a-configuration-group)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="list-configuration-group"></a>

## List configuration groups

This command lists configuration groups.

```console
ngsi services [command options] list [options]
```

### Options

| Options                   | Description                                                |
| ------------------------- | ---------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                     |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                       |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                   |
| --limit VALUE             | maximum number of services                                 |
| --offset VALUE            | offset to skip a given number of elements at the beginning |
| --resource VALUE          | uri for the iotagent                                       |
| --pretty, -P              | pretty format (default: false)                             |
| --help                    | show help (default: true)                                  |

### Examples

#### Request:

```console
ngsi services list --pretty

```json
{
  "count": 1,
  "services": [
    {
      "commands": [],
      "lazy": [],
      "attributes": [],
      "_id": "601e25597d7b3d691be82d23",
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
````

<a name="create-a-configuration-group"></a>

## Create a configuration group 

This command creates a configuration group.

```console
ngsi services [command options] create [options]
```

### Options

| Options                   | Description                                      |
| ------------------------- | ------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)           |
| --service VALUE, -s VALUE | FIWARE Service VALUE                             |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                         |
| --data VALUE              | data body (payload)                              |
| --apikey VALUE            | a key used for devices belonging to this service |
| --token VALUE             | token obtained from the authentication system    |
| --cbroker VALUE           | url of context broker or broker alias            |
| --type VALUE, -t VALUE    | Entity Type                                      |
| --resource VALUE          | uri for the iotagent                             |
| --help                    | show help (default: true)                        |

### Examples

#### Request:

```console
ngsi services create \
--apikey apikey \
--cbroker http://orion:1026 \
--type Event \
--resource /iot/d
```

#### Request:

```console
ngsi services --host iota create --data \
'{
  "services": [
    {
      "apikey": "apikey",
      "cbroker": "http://orion:1026",
      "entity_type": "Thing",
      "resource": "/iot/d"
    }
  ]
}'
```

<a name="update-a-configuration-group"></a>

## Update a configuration group

This command updates a configuration group.

```console
ngsi services [command options] update [options]
```

### Options

| Options                   | Description                                      |
| ------------------------- | ------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)           |
| --service VALUE, -s VALUE | FIWARE Service VALUE                             |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                         |
| --resource VALUE          | uri for the iotagent (required)                  |
| --data VALUE              | data body (payload)                              |
| --apikey VALUE            | a key used for devices belonging to this service |
| --token VALUE             | token obtained from the authentication system    |
| --cbroker VALUE           | url of context broker or broker alias            |
| --type VALUE, -t VALUE    | Entity Type                                      |
| --help                    | show help (default: true)                        |

### Examples

#### Request:

```console
ngsi services update \
--resource /iot/d \
--apikey 4jggokgpepnvsb2uv4s40d59ov \
--type Event
```

<a name="delete-a-configuration-group"></a>

## Delete a configuration group

This command deletes a configuration group.

```console
ngsi services [command options] delete [options]
```

### Options

| Options                   | Description                                           |
| ------------------------- | ----------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                  |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                              |
| --resource VALUE          | uri for the iotagent (required)                       |
| --apikey VALUE            | a key used for devices belonging to this service      |
| --device                  | remove devices in service/subservice (default: false) |
| --help                    | show help (default: true)                             |

### Examples

#### Request:

```console
ngsi services delete --resource /iot/d
```

```console
ngsi services delete --resource /iot/d --device on
```
