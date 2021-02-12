# applications / organizations - Keyrock command

This command allows you to manage organizations in an application for Keyrock.

-   [list organizations in an application](#list-organizations-in-an-application)
-   [Get roles of an organization in an application](#get-roles-of-an-organization-in-an-application)
-   [Assign a role to an organization](#assign-a-role-to-an-organization)
-   [Delete a role assignment from an organization](#delete-a-role-assignment-from-an-organization)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-organizations-in-an-application"></a>

## List organizations

This command lists all organizations.

```console
ngsi applications [command options] organizations --aid {id} list [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --aid value, -i value  | application id                 |
| --verbose, -v  verbose | (default: false)               |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications organizations --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "role_organization_assignments": [
    {
      "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
      "role_organization": "member",
      "role_id": "fb3b897b-5484-4c5b-93d9-4669a46422ad"
    }
  ]
}
```

<a name="get-roles-of-an-organization-in-an-application"></a>

## Get roles of an organization in an application

This command gets roles of an organization in an application.

```console
ngsi application [command options] organizations --aid {id} get [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --oid value, -o value | organization id                |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications organizations --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  get --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca --pretty
```

```console
{
  "role_organization_assignments": [
    {
      "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
      "role_id": "fb3b897b-5484-4c5b-93d9-4669a46422ad"
    }
  ]
}
```

<a name="assign-a-role-to-an-organization"></a>

## Assign a role to an organization

This command assigns a role to an organization.

```console
ngsi application [command options] organizations --aid {id} assigns [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --aid value, -i value  | application id                 |
| --oid value, -o value  | organization id                |
| --rid value, -r value  | role id                        |
| --orid value, -c value | organization role id           |
| --verbose, -v          | verbose (default: false)       |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications organizations --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  assign --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca \
         --rid fb3b897b-5484-4c5b-93d9-4669a46422ad \
         --orid member --pretty
```

```console
{
  "role_organization_assignments": {
    "role_id": "fb3b897b-5484-4c5b-93d9-4669a46422ad",
    "organization_id": "f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca",
    "oauth_client_id": "8b58ecff-fb45-4811-945c-6f42339db06b",
    "role_organization": "member"
  }
}
```

<a name="delete-a-role-assignment-from-an-organization"></a>

## Delete a role assignment from an organization

This command deletes a role assignment from an organization.

```console
ngsi application [command options] organizations --aid {id} unassign [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --aid value, -i value  | application id                 |
| --oid value, -o value  | organization id                |
| --rid value, -r value  | role id                        |
| --orid value, -c value | organization role id           |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications organizations --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  unassign --oid f1f2fd72-12ee-4ced-bbe8-1d99803fa0ca \
           --rid fb3b897b-5484-4c5b-93d9-4669a46422ad \
           --orid member --pretty
```
