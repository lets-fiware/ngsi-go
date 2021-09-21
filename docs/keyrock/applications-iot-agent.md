# applications / iota - Keyrock command

This command allows you to manage IoT Agents in an application for Keyrock.

-   [List IoT Agents](#list-iot-agents)
-   [Get an IoT Agent](#get-an-iot-agent)
-   [Create an IoT Agent](#create-an-iot-agent)
-   [Reset an IoT Agent](#reset-an-iot-agent)
-   [Delete an IoT Agent](#delete-an-iot-agent)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --help                 | show help (default: true)              |

<a name="list-iot-agents"></a>

## List IoT Agents

This command lists IoT Agents.

```console
ngsi applications [command options] iota --aid {id} list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications iota --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "iot_agents": [
    {
      "id": "iot_sensor_add39bbd-db1b-42b6-b669-f882c00ee01c"
    },
    {
      "id": "iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615"
    }
  ]
}
```

<a name="get-iot-agent"></a>

## Get an IoT Agent

This command gets an IoT Agent.

```console
ngsi application [command options] iota --aid {id} add [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --iid VALUE, -i VALUE  | IoT Agent id (required)                |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications iota --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  get --iid iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615 --pretty
```

```console
{
  "iot_agent": {
    "id": "iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615",
    "oauth_client_id": "8b58ecff-fb45-4811-945c-6f42339db06b"
  }
}
```

<a name="create-iot-agent"></a>

## Create IoT Agent

This command creates an IoT Agent.

```console
ngsi application [command options] iota --aid {id} create [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications iota --aid 8b58ecff-fb45-4811-945c-6f42339db06b create --pretty
```

```console
{
  "iot_agent": {
    "id": "iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615",
    "password": "iot_sensor_25812a47-58bf-4970-b1a3-ea4ff017d360"
  }
}
```

<a name="reset-an-iot-agent"></a>

## Reset an IoT Agent

This command resets an IoT Agent.

```console
ngsi application [command options] iota --aid {id} reset [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --iid VALUE, -i VALUE  | IoT Agent id (required)                |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications iota --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  reset --iid iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615 --pretty
```

```console
{
  "new_password": "iot_sensor_7039c3dd-2761-4bce-b0e7-e9c814239844"
}
```

<a name="delete-an-iot-agent"></a>

## Delete an IoT Agent

This command deletes an IoT Agent.

```console
ngsi application [command options] iota --aid {id} delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --iid VALUE, -i VALUE  | IoT Agent id (required)                |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications iota --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  delete --iid iot_sensor_e4fff656-cc64-44be-8f47-757b1bc95615
```
