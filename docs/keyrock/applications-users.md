# applications / users - Keyrock command

This command allows you to manage users for Keyrock.

-   [List users](#list-users)
-   [Get a user](#get-a-user)
-   [Assign a user](#assign-a-user)
-   [Unassign a user](#unassign-a-user)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --help                 | show help (default: true)              |

<a name="list-users"></a>

## List users

This command lists all users.

```console
ngsi applications [command options] users --aid {id} list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

<a name="get-a-user"></a>

## Get a user.

This command gets a user.

```console
ngsi application [command options] users --aid {id} get [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

<a name="assign-a-user"></a>

## assign a user

This command assigns a user.

```console
ngsi application [command options] users --aid {id} assign [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

<a name="unassign-a-user"></a>

## Unassign a user 

This command unassigns a user.

```console
ngsi application [command options] users --aid {id} unassign [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --help                 | show help (default: true)              |
