# applications / roles - Keyrock command

This command allows you to manage roles for Keyrock.

-   [List roles](#list-roles)
-   [Get a role](#get-a-role)
-   [Create a role](#create-a-role)
-   [Update a role](#update-a-role)
-   [Delete a role](#delete-a-role)
-   [List permissions associated to a role](#list-permissions-associated-to-a-role)
-   [Assign a permission to a role](#assign-a-permission-to-a-role)
-   [Delete a permission from a role](#delete-a-permission-to-a-role)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --help                 | show help (default: true)              |

<a name="list-roles"></a>

## List roles

This command lists all roles.

```console
ngsi applications [command options] roles --aid {id} list [options]
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
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "roles": [
    {
      "id": "purchaser",
      "name": "Purchaser"
    },
    {
      "id": "provider",
      "name": "Provider"
    }
  ]
}
```

<a name="get-a-role"></a>

## Get a role.

This command gets a role.

```console
ngsi application [command options] roles --aid {id} get [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b get --rid purchaser --pretty
```

```console
{
  "role": {
    "id": "purchaser",
    "name": "Purchaser",
    "is_internal": true,
    "oauth_client_id": "idm_admin_app"
  }
}
```

<a name="create-a-role"></a>

## Create a role

This command creates a role.

```console
ngsi application [command options] roles --aid {id} create [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --data VALUE, -d VALUE | role data                              |
| --name VALUE, -n VALUE | role name                              |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b create --name role1
```

```console
dd214cb3-c445-4ae1-88bf-65db88226b51
```

<a name="update-a-role"></a>

## Update a role 

This command updates a role.

```console
ngsi application [command options] roles --aid {id} update [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --data VALUE, -d VALUE | role data                              |
| --name VALUE, -n VALUE | role name                              |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  update --rid dd214cb3-c445-4ae1-88bf-65db88226b51 --name "role99"
```

```console
{"values_updated":{"name":"role99"}}
```

<a name="delete-a-role"></a>

## Delete a role 

This command deletes a role.

```console
ngsi application [command options] roles delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  delete --rid dd214cb3-c445-4ae1-88bf-65db88226b51
```

<a name="list-permissions-associated-to-a-role"></a>

## List permissions associated to a role

This command list permissions associated to a role.

```console
ngsi application [command options] roles permissions [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  permissions --rid 7423d744-8682-4d5d-b338-2f4efcfa824e --pretty
```

```console
{
  "role_permission_assignments": [
    {
      "id": "5ed4075e-cc31-4830-8b9a-f7a04eb25a36",
      "is_internal": false,
      "name": "permission1",
      "description": "test",
      "action": "GET",
      "resource": "login",
      "xml": null
    }
  ]
}
```

<a name="assign-a-permission-to-a-role"></a>

## Assign a permission to a role

This command assigns a permission to a role.

```console
ngsi application [command options] roles assign [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --pid VALUE, -p VALUE  | permission id (required)               |
| --verbose, -v          | verbose (default: false)               |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  assign --rid 7423d744-8682-4d5d-b338-2f4efcfa824e --pid 5ed4075e-cc31-4830-8b9a-f7a04eb25a36
```

```console
{
  "role_permission_assignments": {
    "role_id": "7423d744-8682-4d5d-b338-2f4efcfa824e",
    "permission_id": "5ed4075e-cc31-4830-8b9a-f7a04eb25a36"
  }
}
```

<a name="delete-a-permission-to-a-role"></a>

## Delete a permission from a role

This command deletes a permission from a role

```console
ngsi application [command options] roles unassign [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --rid VALUE, -r VALUE  | role id (required)                     |
| --pid VALUE, -p VALUE  | permission id (required)               |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications roles --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  unassign --rid 7423d744-8682-4d5d-b338-2f4efcfa824e --pid 5ed4075e-cc31-4830-8b9a-f7a04eb25a36
```
