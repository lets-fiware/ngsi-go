# upsert - NGSI command

This command upserts entities.

-   [Upsert an entity](#upsert-an-entity)
-   [Upsert multiple entities](#upsert-multiple-entities)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="upsert-an-entity"/>

## Upset an entity

This command upserts an entity.

```console
ngsi upsert [common options] entity [options]
```

### Options

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --data value, -d value    | specify data                       |
| --keyValues, -k           | specify keyValues (default: false) |
| --link value, -L value    | specify @context                   |
| --help                    | show help (default: false)         |

### Example

#### Request:

```console
ngsi upsert entity \
--data ' {
      "id":"urn:ngsi-ld:Product:010",
      "type":"Product",
      "name":{"type":"Text", "value":"Lemonade"},
      "size":{"type":"Text", "value": "S"},
      "price":{"type":"Integer", "value": 99}
}'
```

<a name="upsert-multiple-entities"/>

## Upsert multiple entities

This command upserts multiple entities.

```console
ngsi upsert [common options] entities [options]
```

### Options

| Options                   | Description                      |
| ------------------------- | -------------------------------- |
| --data value, -d value    | specify data                     |
| --replace, -r             | specfiy replace (default: false) |
| --update, -u              | specify update (default: false)  |
| --link value, -L value    | specify @context                 |
| --help                    | show help (default: false)       |

### Example

```console
ngsi upsert entities \
--data '[
  {
    "id":"urn:ngsi-ld:Product:011", "type":"Product",
    "name":{"type":"Text", "value":"Brandy"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1199}
  },
  {
    "id":"urn:ngsi-ld:Product:012", "type":"Product",
    "name":{"type":"Text", "value":"Port"},
    "size":{"type":"Text", "value": "M"},
    "price":{"type":"Integer", "value": 1099}
  }
]'
```
