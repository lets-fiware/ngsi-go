# applications / trustedi - Keyrock command

This command allows you to manage trusted applications for Keyrock.

-   [List trusted applications associated to an application](#list-trusted-applications-associated-to-an-application)
-   [Add trusted application](#add-a-trusted-application)
-   [Delete trusted application](#delete-a-trusted-application)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --help                 | show help (default: true)              |

<a name="list-trusted-applications-associated-to-an-application"></a>

## List trusted applications associated to an application

This command lists trusted applications associated to an application.

```console
ngsi applications [command options] trusted --aid {id} list [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications trusted --aid 8b58ecff-fb45-4811-945c-6f42339db06b list --pretty
```

```console
{
  "trusted_applications": [
    "97235ddb-e690-42ff-a4a9-5488bffa4b3b"
  ]
}
```

<a name="add-a-trusted-application"></a>

## Add a trusted application

This command adds a trusted application

```console
ngsi application [command options] trusted --aid {id} add [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --tid VALUE, -t VALUE  | trusted application id (required)      |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications trusted --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  add --tid 97235ddb-e690-42ff-a4a9-5488bffa4b3b
```

```console
{
  "oauth_client_id": "8b58ecff-fb45-4811-945c-6f42339db06b",
  "trusted_oauth_client_id": "97235ddb-e690-42ff-a4a9-5488bffa4b3b"
}
```

<a name="delete-a-trusted-application"></a>

## Delete a trusted application

This command deletes a trusted application.

```console
ngsi application [command options] trusted --aid {id} delete [options]
```

### Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --aid VALUE, -i VALUE  | application id (required)              |
| --tid VALUE, -t VALUE  | trusted application id (required)      |
| --help                 | show help (default: true)              |

### Examples

#### Request:

```console
ngsi applications trusted --aid 8b58ecff-fb45-4811-945c-6f42339db06b \
  delete --tid 97235ddb-e690-42ff-a4a9-5488bffa4b3b
```
