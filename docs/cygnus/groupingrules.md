# groupingrules - Cygnus command

This command allows you to manage groupingrules for Cygnus.

-   [List grouping rules](#list-groupingrules)
-   [Create a grouping rule](#create-a-groupingrule)
-   [Update a grouping rule](#update-a-groupingrule)
-   [Delete a grouping rule](#delete-a-groupingrule)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-groupingrules"></a>

## List grouping rules

This command lists all grouping rules.

```console
ngsi groupingrules [command options] list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi groupingrules list --pretty
```

```console
{
  "success": "true",
  "grouping_rules": [
    {
      "regex": "Room",
      "fiware_service_path": "\/rooms",
      "destination": "allrooms",
      "id": 1,
      "fields": [
        "entityType"
      ]
    }
  ]
}
```

<a name="create-a-groupingrule"></a>

## Create a grouping rule

This command creates a grouping rule.

```console
ngsi groupingrule [command options] create [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --data VALUE, -d VALUE | grouping rule data (required)          |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi groupingrules --host cygnus create --pretty --data \
'{
 "regex": "Car",
 "destination": "allcars",
 "fiware_service_path": "/cars",
 "fields": ["entityType"]
}'
```

```console
{
  "success": "true"
}
```

<a name="update-a-groupingrule"></a>

## Update a grouping rule

This command updates a grouping rule.

```console
ngsi groupingrule [command options] update [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --id VALUE, -i VALUE   | grouping rule id (required)            |
| --data VALUE, -d VALUE | grouping rule data (required)          |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi groupingrules --host cygnus update --id 2 --data \
'{
  "regex": "Device",
  "destination": "alldevices",
  "fiware_service_path": "/devices",
  "fields": [
    "entityType"
  ]
}'
```

```console
{"success":"true"}
```

<a name="delete-a-groupingrule"></a>

## Delete a grouping rule

This command deletes a grouping rule.

```console
ngsi groupingrule [command options] delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --id VALUE, -i VALUE   | grouping rule id (required)            |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi groupingrules delete --id 1
```

```json
{"success":"true"}
```
