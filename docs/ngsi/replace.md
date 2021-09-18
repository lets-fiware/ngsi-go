# replace - NGSI command

This command replaces multiple entities or attributes

-   [Replace multiple entities](#replace-multiple-entities)
-   [Replace multiple attributes](#replace-multiple-attributes)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --help                    | show help (default: true)              |

<a name="replace-multiple-entities"></a>

## Replace multiple entities

This command replaces multiple entities.

```console
ngsi replace [common options] entities [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --data VALUE, -d VALUE    | entities data (required)               |
| --keyValues, -K           | keyValues (default: false)             |
| --help                    | show help (default: true)              |

### Example

```console
ngsi replace entities \
--data '[
{
    "id":"urn:ngsi-ld:Product:010", "type":"Product",
    "price":{"type":"Integer", "value": 1199}
  }
]'
```

<a name="replace-multiple-attributes"></a>

## Replace multiple attributes

This command replaces multiple attributes.

```console
ngsi replace [common options] attrs [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --type VALUE, -t VALUE    | entity type                            |
| --data VALUE, -d VALUE    | attributes data                        |
| --keyValues, -K           | keyValues (default: false)             |
| --help                    | show help (default: true)              |

### Example

#### Request

```console
ngsi get entity --id urn:ngsi-ld:Sensor:001
```

```json
{"id":"urn:ngsi-ld:Sensor:001","type":"Sensor","Temperature":{"type":"Text","value":"30","metadata":{}}}
```

#### Request

```console
ngsi replace attrs --id urn:ngsi-ld:Sensor:001 --keyValues --data '{"Temperature":30}'
```

#### Request

```console
ngsi get entity --id urn:ngsi-ld:Sensor:001
```

```json
{"id":"urn:ngsi-ld:Sensor:001","type":"Sensor","Temperature":{"type":"Number","value":30,"metadata":{}}}
```
