# users - Keyrock command

This command allows you to manage users for Keyrock.

-   [List users](#list-users)
-   [Get a users](#get-a-user)
-   [Create a users](#create-a-user)
-   [Get a users](#update-an-get-user)
-   [Delete a users](#delete-a-user)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-user"></a>

## List users

This command lists all users.

```console
ngsi users [command options] list [options]
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
ngsi users list --pretty
```

```console
{
  "users": [
    {
      "id": "31ea0ac1-595f-479e-9854-f911a26a3d51",
      "username": "alice",
      "email": "alice@test.com",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-20T20:42:23.000Z",
      "description": null,
      "website": null
    },
    {
      "id": "admin",
      "username": "admin",
      "email": "keyrock@letsfiware.jp",
      "enabled": true,
      "gravatar": false,
      "date_password": "2021-02-20T20:40:33.000Z",
      "description": null,
      "website": null
    }
  ]
}
```

<a name="get-a-user"></a>

## Get a user

This command gets a user.

```console
ngsi user [command options] get [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi users get --uid 31ea0ac1-595f-479e-9854-f911a26a3d51 --pretty
```

```console
{
  "user": {
    "id": "31ea0ac1-595f-479e-9854-f911a26a3d51",
    "username": "alice",
    "email": "alice@test.com",
    "enabled": true,
    "admin": false,
    "image": "default",
    "gravatar": false,
    "date_password": "2021-02-20T20:42:23.000Z",
    "description": null,
    "website": null,
    "extra": null
  }
}
```

<a name="create-a-user"></a>

## Create a user

This command creates a user.

```console
ngsi user [command options] create [options]
```

### Options

| Options                    | Description                            |
| -------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE     | broker or server host VALUE (required) |
| --username VALUE, -u VALUE | user name (required)                   |
| --password VALUE, -p VALUE | password (required)                    |
| --email VALUE, -e VALUE    | email (required)                       |
| --verbose, -v              | verbose (default: false)               |
| --pretty, -P               | pretty format (default: false)         |
| --help                     | show help (default: true)              |

### Examples

#### Request:

```console
ngsi users create --username alice --email alice@test.com --password test
```

```console
31ea0ac1-595f-479e-9854-f911a26a3d51
```

<a name="update-a-user"></a>

## Update a user

This command updates a user.

```console
ngsi user [command options] update [options]
```

### Options

| Options                       | Description                            |
| ----------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE        | broker or server host VALUE (required) |
| --uid VALUE, -i VALUE         | user id (required)                     |
| --username VALUE, -u VALUE    | user name                              |
| --password VALUE, -p VALUE    | password                               |
| --email VALUE, -e VALUE       | email                                  |
| --gravatar, -g                | gravatar (default: false)              |
| --description VALUE, -d VALUE | description                            |
| --website VALUE, -w VALUE     | website                                |
| --extra VALUE, -E VALUE       | extra information                      |
| --pretty, -P                  | pretty format (default: false)         |
| --help                        | show help (default: true)              |

### Examples

#### Request:

```console
ngsi users update --uid 31ea0ac1-595f-479e-9854-f911a26a3d51 --description "test user" --pretty
```

```console
{
  "values_updated": {
    "description": "test user"
  }
}
```

<a name="delete-a-user"></a>

## Delete a user

This command deletes a user.

```console
ngsi user [command options] delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi users delete --uid 31ea0ac1-595f-479e-9854-f911a26a3d51
```
