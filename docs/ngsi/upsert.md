# upsert - NGSI command

This command upserts entities.

-   [Upsert an entity](#upsert-an-entity)
-   [Upsert multiple entities](#upsert-multiple-entities)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="upsert-an-entity"></a>

## Upset an entity

This command upserts an entity.

```console
ngsi upsert [common options] entity [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | entity data (required)                 |
| --keyValues, -K           | keyValues (default: false)             |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

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

<a name="upsert-multiple-entities"></a>

## Upsert multiple entities

This command upserts multiple entities.

```console
ngsi upsert [common options] entities [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | entities data (required)               |
| --replace, -r             | replace (default: false)               |
| --update, -u              | update (default: false)                |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --help                    | show help (default: true)              |

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
