# organizations / users - Keyrock command

This command allows you to manage users of organization for Keyrock.

-   [List users of an organization](#list-users-of-an-organization)
-   [Get info of user organization relationship](#get-info-of-user-organization-relationship)
-   [Create organization relationship](#add-a-user-to-an-organization)
-   [Delete organization relationship](#remove-a-user-from-an-organization)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --help                 | show help (default: true)              |

<a name="list-organization"></a>

## List users of an organization

This command lists users of an organization.

```console
ngsi organizations [command options] users --oid {id} list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations users --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca list --pretty
```

```console
{
  "organization_users": [
    {
      "user_id": "admin",
      "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
      "role": "owner"
    },
    {
      "user_id": "b97f26a5-c8da-4fa4-9af1-c26013538a7f",
      "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
      "role": "member"
    }
  ]
}
```

<a name="get-info-of-user-organization-relationship"></a>

## Get info of user organization relationship

This command gets info of user organization relationship.

```console
ngsi organization [command options] users --oid {id} get [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations users --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca \
  get --uid b97f26a5-c8da-4fa4-9af1-c26013538a7f --pretty
```

```console
{
  "organization_user": {
    "user_id": "b97f26a5-c8da-4fa4-9af1-c26013538a7f",
    "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
    "role": "member"
  }
}
```

<a name="add-a-user-to-an-organization"></a>

## Add a user to an organization

This command adds a user to an organization.

```console
ngsi organization [command options] users --oid {id} add [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --orid VALUE, -c VALUE | organization role id (required)        |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations users --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca \
  add --uid b97f26a5-c8da-4fa4-9af1-c26013538a7f --orid member --pretty
```

```console
{
  "user_organization_assignments": {
    "role": "member",
    "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
    "user_id": "b97f26a5-c8da-4fa4-9af1-c26013538a7f"
  }
}
```

<a name="remove-a-user-from-an-organization"></a>

## Remove a user from an organization

This command removes a user from an organization.

```console
ngsi organization [command options] users --oid {id} remove [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --oid VALUE, -o VALUE  | organization id (required)             |
| --uid VALUE, -i VALUE  | user id (required)                     |
| --orid VALUE, -c VALUE | organization role id (required)        |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi organizations users --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca \
  remove --uid b97f26a5-c8da-4fa4-9af1-c26013538a7f --orid member
```
