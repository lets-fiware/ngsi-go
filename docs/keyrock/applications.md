# applications - Keyrock command

This command allows you to manage applications for Keyrock.

-   [List applications](#list-applications)
-   [Get an applications](#get-an-application)
-   [Create an applications](#create-an-application)
-   [Update an applications](#update-an-get-application)
-   [Delete an applications](#delete-an-application)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-application"></a>

## List applications

This command lists all applications.

```console
ngsi applications [command options] list [options]
```

### Options

| Options       | Description                    |
| ------------- | ------------------------------ |
| --verbose, -v | verbose (default: false)       |
| --pretty, -P  | pretty format (default: false) |
| --help        | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications list --pretty
```

```console
{
  "applications": [
    {
      "id": "58de156f-0fec-400b-bc7c-86265a885bee",
      "name": "Test_application 1",
      "description": "test app",
      "image": "default",
      "url": "http://localhost",
      "redirect_uri": "http://localhost/login",
      "redirect_sign_out_uri": null,
      "grant_type": "password,authorization_code,implicit",
      "response_type": "code,token",
      "token_types": "jwt,permanent,bearer",
      "jwt_secret": "f2be71188564ba0a",
      "client_type": null,
      "urls": {
        "permissions_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/permissions",
        "roles_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/roles",
        "users_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/users",
        "pep_proxies_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/pep_proxies",
        "iot_agents_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/iot_agents",
        "trusted_applications_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/trusted_applications"
      }
    }
  ]
}
```

<a name="get-an-application"></a>

## Get an application.

This command gets an application.

```console
ngsi application [command options] get [options]
```

### Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --aid value, -i value | application id                 |
| --pretty, -P          | pretty format (default: false) |
| --help                | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications get --aid 58de156f-0fec-400b-bc7c-86265a885bee --pretty
```

```console
{
  "application": {
    "id": "58de156f-0fec-400b-bc7c-86265a885bee",
    "name": "Test_application 1",
    "description": "test app",
    "secret": "921cf732-b39c-4e7c-815c-a91277e52b70",
    "url": "http://localhost",
    "redirect_uri": "http://localhost/login",
    "redirect_sign_out_uri": null,
    "image": "default",
    "grant_type": "password,authorization_code,implicit",
    "response_type": "code,token",
    "token_types": "jwt,permanent,bearer",
    "jwt_secret": "f2be71188564ba0a",
    "client_type": null,
    "scope": null,
    "extra": null,
    "urls": {
      "permissions_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/permissions",
      "roles_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/roles",
      "users_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/users",
      "pep_proxies_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/pep_proxies",
      "iot_agents_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/iot_agents",
      "trusted_applications_url": "/v1/applications/58de156f-0fec-400b-bc7c-86265a885bee/trusted_applications"
    }
  }
}
```

<a name="create-an-application"></a>

## Create an application

This command creates an application.

```console
ngsi application [command options] create [options]
```

### Options

| Options                        | Description                    |
| ------------------------------ | ------------------------------ |
| --data value, -d value         | application data               |
| --name value, -n value         | application name               |
| --description value, -D value  | description                    |
| --url value, -u value          | url                            |
| --redirectUri value, -R value  | redirect uri                   |
| --grantType value, -g value    | grant type                     |
| --tokenTypes value, -t value   | token types                    |
| --responseType value, -r value | response type                  |
| --clientType value, -c value   | client type                    |
| --verbose, -v                  | verbose (default: false)       |
| --pretty, -P                   | pretty format (default: false) |
| --help                         | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications create --name "Test_application 1" \
  --description "test app" \
  --redirectUri http://localhost/login \
  --url http://localhost \
  --grantType authorization_code,implicit,password \
  --tokenTypes jwt,permanent
```

```console
58de156f-0fec-400b-bc7c-86265a885bee
```

<a name="update-an-application"></a>

## Update an application

This command updates an application.

```console
ngsi application [command options] update [options]
```

### Options

| Options                        | Description                    |
| ------------------------------ | ------------------------------ |
| --aid value, -i value          | application id                 |
| --data value, -d value         | application data               |
| --name value, -n value         | application name               |
| --description value, -D value  | description                    |
| --url value, -u value          | url                            |
| --redirectUri value, -R value  | redirect uri                   |
| --grantType value, -g value    | grant type                     |
| --tokenTypes value, -t value   | token types                    |
| --responseType value, -r value | response type                  |
| --clientType value, -c value   | client type                    |
| --verbose, -v                  | verbose (default: false)       |
| --pretty, -P                   | pretty format (default: false) |
| --help                         | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications update --aid 58de156f-0fec-400b-bc7c-86265a885bee --url http://orion/ --pretty
```

```console
{
  "values_updated": {
    "url": "http://orion/",
    "token_types": "jwt,permanent,bearer,bearer",
    "scope": ""
  }
}
```

<a name="delete-an-application"></a>

## Delete an application

This command deletes an application.

```console
ngsi application [command options] delete [options]
```

### Options

| Options                        | Description                    |
| ------------------------------ | ------------------------------ |
| --aid value, -i value          | application id                 |
| --help                         | show help (default: false)     |

### Examples

#### Request:

```console
ngsi applications delete --aid 58de156f-0fec-400b-bc7c-86265a885bee
```
