# organizations - Keyrock command

This command allows you to manage organizations for Keyrock.

-   [List organizations](#list-organizations)
-   [Get an organizations](#get-an-organization)
-   [Create an organizations](#create-an-organization)
-   [Get an organizations](#update-an-get-organization)
-   [Delete an organizations](#delete-an-organization)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-organization"></a>

## List organizations

This command lists all organizations.

```console
ngsi organizations [command options] list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations list --pretty
```

```console
{
  "organizations": [
    {
      "role": "owner",
      "Organization": {
        "id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
        "name": "test",
        "description": "test organizations",
        "image": "default",
        "website": null
      }
    },
    {
      "role": "owner",
      "Organization": {
        "id": "f672e00e-9f19-430b-8e0b-06b3ac7a8e4d",
        "name": "test",
        "description": "test organizations",
        "image": "default",
        "website": null
      }
    }
  ]
}
```

<a name="get-an-organization"></a>

## Get an organization

This command gets an organization.

```console
ngsi organization [command options] get [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations get --oid f672e00e-9f19-430b-8e0b-06b3ac7a8e4d --pretty
```

```console
{
  "organization": {
    "id": "f672e00e-9f19-430b-8e0b-06b3ac7a8e4d",
    "name": "test",
    "description": "test organizations",
    "website": null,
    "image": "default"
  }
}
```

<a name="create-an-organization"></a>

## Create an organization

This command creates an organization.

```console
ngsi organization [command options] create [options]
```

### Options

| Options                       | Description                            |
| ----------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE        | broker or server host VALUE (required) |
| --name VALUE, -n VALUE        | organization name (required)           |
| --description VALUE, -d VALUE | description                            |
| --website VALUE, -w VALUE     | website                                |
| --verbose, -v                 | verbose (default: false)               |
| --pretty, -P                  | pretty format (default: false)         |
| --help                        | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations create --name test --description "test organizations"
```

```console
f672e00e-9f19-430b-8e0b-06b3ac7a8e4d
```

<a name="update-an-organization"></a>

## Update an organization

This command updates an organization.

```console
ngsi organization [command options] update [options]
```

### Options

| Options                       | Description                            |
| ----------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE        | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE         | organization id (required)             |
| --name VALUE, -n VALUE        | organization name                      |
| --description VALUE, -d VALUE | description                            |
| --website VALUE, -w VALUE     | website                                |
| --verbose, -v                 | verbose (default: false)               |
| --pretty, -P                  | pretty format (default: false)         |
| --help                        | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations update --oid f672e00e-9f19-430b-8e0b-06b3ac7a8e4d --website https://www.letsfiware.jp/ --pretty
```

```console
{
  "values_updated": {
    "website": "https://www.letsfiware.jp/"
  }
}
```

<a name="delete-an-organization"></a>

## Delete an organization

This command deletes an organization.

```console
ngsi organization [command options] delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations delete --oid f672e00e-9f19-430b-8e0b-06b3ac7a8e4d
```
