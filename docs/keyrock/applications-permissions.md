# applications / permissions - Keyrock command

This command allows you to manage permissions for Keyrock.

-   [List permissions](#list-permissions)
-   [Get a permission](#get-a-permission)
-   [Create a permission](#create-a-permission)
-   [Update a permission](#update-a-permission)
-   [Delete a permission](#delete-a-permission)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-permissions"></a>

## List permissions

This command lists all permissions.

```console
ngsi applications [command options] permissions --aid {id} list [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --pid value, -p value | permission id                  |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications permissions --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "permissions": [
    {
      "id": "6",
      "name": "Get and assign only public owned roles",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    },
    {
      "id": "5",
      "name": "Get and assign all public application roles",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    },
    {
      "id": "4",
      "name": "Manage authorizations",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    },
    {
      "id": "3",
      "name": "Manage roles",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    },
    {
      "id": "2",
      "name": "Manage the application",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    },
    {
      "id": "1",
      "name": "Get and assign all internal application roles",
      "description": null,
      "action": null,
      "resource": null,
      "xml": null
    }
  ]
}
```

<a name="get-a-permission"></a>

## Get a permission.

This command gets a permission.

```console
ngsi application [command options] permissions --aid {id} get [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --pid value, -p value | permission id                  |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications permissions --aid 8b58ecff-fb45-4811-945c-6f42339db06b get --pid 1 --pretty
```

```console
{
  "permission": {
    "id": "1",
    "name": "Get and assign all internal application roles",
    "description": null,
    "is_internal": true,
    "action": null,
    "resource": null,
    "is_regex": 0,
    "xml": null,
    "oauth_client_id": "idm_admin_app"
  }
}
```

<a name="create-a-permission"></a>

## Create a permission

This command creates a permission.

```console
ngsi application [command options] permissions --aid {id} create [options]
```

### Options

| Options                       | Description                    |
| ----------------------------- | ------------------------------ |
| --pid value, -p value         | permission id                  |
| --data value, -d value        | permission data                |
| --name value, -n value        | permission name                |
| --description value, -D value | description                    |
| --action value, -a value      | action                         |
| --resource value, -r value    | resoruce                       |
| --verbose, -v                 | verbose (default: false)       |
| --pretty, -P                  | pretty format (default: false) |
| --help                        | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications permissions --aid $aid \
  create --name "permission1" \
  --description "test" \
  --action "GET" \
  --resource "login"
```

```console
ab781799-d2bb-4022-b4b9-5101cbc98e12
```

<a name="update-a-permission"></a>

## Update a permission 

This command updates a permission.

```console
ngsi application [command options] permissions --aid {id} update [options]
```

### Options

| Options                       | Description                    |
| ----------------------------- | ------------------------------ |
| --aid value, -i value         | application id                 |
| --pid value, -p value         | permission id                  |
| --data value, -d value        | permission data                |
| --name value, -n value        | permission name                |
| --description value, -D value | description                    |
| --action value, -a value      | action                         |
| --resource value, -r value    | resoruce                       |
| --verbose, -v                 | verbose (default: false)       |
| --pretty, -P                  | pretty format (default: false) |
| --help                        | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications permissions --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  update --pid ab781799-d2bb-4022-b4b9-5101cbc98e12 --name "perm1"
```

```console
{"values_updated":{"name":"perm1"}}
```

<a name="delete-a-permission"></a>

## Delete a permission 

This command deletes a permission.

```console
ngsi application [command options] permissions --aid {id} delete [options]
```

### Options

| Options                       | Description                    |
| ----------------------------- | ------------------------------ |
| --aid value, -i value         | application id                 |
| --pid value, -p value         | permission id                  |
| --help                        | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications permissions --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  delete --pid ab781799-d2bb-4022-b4b9-5101cbc98e12
```
