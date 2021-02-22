# namemappings - Cygnus command

This command allows you to manage namemappings for Cygnus.

-   [List name mappings](#list-namemappings)
-   [Create a name mapping](#create-a-namemapping)
-   [Delete a name mapping](#delete-a-namemapping)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-namemappings"></a>

## List name mappings

This command lists all name mappings.

```console
ngsi namemappings [command options] list [options]
```

### Options

| Options      | Description                    |
| ------------ | ------------------------------ |
| --pretty, -P | pretty format (default: false) |
| --help       | show help (default: false)     |

### Examples

#### Request:

```console
ngsi namemappings list --pretty
```

```console
{
  "success": "true",
  "result": {
    "serviceMappings": [
      {
        "originalService": "^(.*)",
        "newService": "null",
        "servicePathMappings": [
          {
            "originalServicePath": "/myservicepath1",
            "newServicePath": "/new_myservicepath1",
            "entityMappings": [
              {
                "originalEntityId": "myentityid1",
                "originalEntityType": "myentitytype1",
                "newEntityId": "new_myentityid1",
                "newEntityType": "new_myentitytype1",
                "attributeMappings": [
                  {
                    "originalAttributeName": "myattributename1",
                    "originalAttributeType": "myattributetype1",
                    "newAttributeName": "new_myattributename1",
                    "newAttributeType": "new_myattributetype1"
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  }
}
```

<a name="create-a-namemapping"></a>

## Create a name mapping

This command creates a name mapping.

```console
ngsi namemapping [command options] create [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --data value, -d value | name mapping data              |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi namemappings create -pretty \
--data '{
  "serviceMappings": [
    {
      "servicePathMappings": [
        {
          "originalServicePath": "/myservicepath1",
          "newServicePath": "/new_myservicepath1",
          "entityMappings": [
            {
              "originalEntityId": "myentityid1",
              "originalEntityType": "myentitytype1",
              "newEntityId": "new_myentityid1",
              "newEntityType": "new_myentitytype1",
              "attributeMappings": [
                {
                  "originalAttributeName": "myattributename1",
                  "originalAttributeType": "myattributetype1",
                  "newAttributeName": "new_myattributename1",
                  "newAttributeType": "new_myattributetype1"
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}'
```

```console
{
  "success": "true",
  "result": {
    "serviceMappings": [
      {
        "originalService": "^(.*)",
        "newService": "null",
        "servicePathMappings": [
          {
            "originalServicePath": "/myservicepath1",
            "newServicePath": "/new_myservicepath1",
            "entityMappings": [
              {
                "originalEntityId": "myentityid1",
                "originalEntityType": "myentitytype1",
                "newEntityId": "new_myentityid1",
                "newEntityType": "new_myentitytype1",
                "attributeMappings": [
                  {
                    "originalAttributeName": "myattributename1",
                    "originalAttributeType": "myattributetype1",
                    "newAttributeName": "new_myattributename1",
                    "newAttributeType": "new_myattributetype1"
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  }
}
```

<a name="delete-a-namemapping"></a>

## Delete a name mapping

This command deletes a name mapping.

```console
ngsi namemapping [command options] delete [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --data value, -d value | name mapping data              |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi namemappings delete --pretty --data \
'{
  "serviceMappings": [
    {
      "originalService": "^(.*)",
      "newService": "null",
      "servicePathMappings": [
        {
          "originalServicePath": "/myservicepath1",
          "newServicePath": "/new_myservicepath1",
          "entityMappings": [
            {
              "originalEntityId": "myentityid1",
              "originalEntityType": "myentitytype1",
              "newEntityId": "new_myentityid1",
              "newEntityType": "new_myentitytype1",
              "attributeMappings": [
                {
                  "originalAttributeName": "myattributename1",
                  "originalAttributeType": "myattributetype1",
                  "newAttributeName": "new_myattributename1",
                  "newAttributeType": "new_myattributetype1"
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}'
```

```json
{
  "success": "true",
  "result": {
    "serviceMappings": [
      {
        "originalService": "^(.*)",
        "newService": "null",
        "servicePathMappings": [
          {
            "originalServicePath": "/myservicepath1",
            "newServicePath": "/new_myservicepath1",
            "entityMappings": [
              {
                "originalEntityId": "myentityid1",
                "originalEntityType": "myentitytype1",
                "newEntityId": "new_myentityid1",
                "newEntityType": "new_myentitytype1",
                "attributeMappings": []
              }
            ]
          }
        ]
      }
    ]
  }
}
```
