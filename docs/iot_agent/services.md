# services - IoT Agent command

This command allows you to list, create, update and delete service entry for IoT Agent.

-   [List configuration groups](#list-configuration-group)
-   [Create a configuration group](#create-a-configuration-group)
-   [Update a configuration group](#update-a-configuration-group)
-   [Delete a configuration group](#delete-a-configuration-group)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="list-configuration-group"></a>

## List configuration groups

This command lists configuration groups.

```console
ngsi services [command options] list [options]
```

### Options

| Options          | Description                                                             |
| ---------------- | ----------------------------------------------------------------------- |
| --limit value    | maximum number of services (default: 0)                                 |
| --offset value   | offset to skip a given number of elements at the beginning (default: 0) |
| --resource value | uri for the iotagent                                                    |
| --pretty, -P     | pretty format (default: false)                                          |
| --help           | show help (default: false)                                              |

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

| Options                | Description                                      |
| ---------------------- | ------------------------------------------------ |
| --data value, -d value | data body (payload)                              |
| --apikey value         | a key used for devices belonging to this service |
| --token value          | token obtained from the authentication system    |
| --cbroker value        | url of context broker or broker alias            |
| --type value, -t value | Entity Type                                      |
| --resource value       | uri for the iotagent                             |
| --help                 | show help (default: false)                       |

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

| Options                | Description                                      |
| ---------------------- | ------------------------------------------------ |
| --data value, -d value | data body (payload)                              |
| --apikey value         | a key used for devices belonging to this service |
| --token value          | token obtained from the authentication system    |
| --cbroker value        | url of context broker or broker alias            |
| --type value, -t value | Entity Type                                      |
| --resource value       | uri for the iotagent                             |
| --help                 | show help (default: false)                       |

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

| Options          | Description                                           |
| ---------------- | ----------------------------------------------------- |
| --apikey value   | a key used for devices belonging to this service      |
| --resource value | uri for the iotagent                                  |
| --device         | remove devices in service/subservice (default: false) |
| --help           | show help (default: false)                            |

### Examples

#### Request:

```console
ngsi services delete --resource /iot/d
```

```console
ngsi services delete --resource /iot/d --device on
```
