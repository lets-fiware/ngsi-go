# replace - NGSI command

This command replaces multiple entities or attributes

-   [Replace multiple entities](#replace-multiple-entities)
-   [Replace multiple attributes](#replace-multiple-attributes)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="replace-multiple-entities"/>

## Replace multiple entities

This command replaces multiple entities.

```bash
ngsi replace [common options] entities [options]
```

### Options

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --keyValues, -k        | specify keyValues (default: false) |
| --data value, -d value | specify data                       |
| --help                 | show help (default: false)         |

### Example

```bash
$ ngsi replace entities \
--data '[
{
    "id":"urn:ngsi-ld:Product:010", "type":"Product",
    "price":{"type":"Integer", "value": 1199}
  }
]'
```

<a name="replace-multiple-attributes"/>

## Replace multiple attributes

This command replaces multiple attributes.

```bash
ngsi replace [common options] attrs [options]
```

### Options

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --id value, -i value   | specify id                         |
| --keyValues, -k        | specify keyValues (default: false) |
| --data value, -d value | specify data                       |
| --help                 | show help (default: false)         |

### Example

```bash
$ ngsi get entity --id urn:ngsi-ld:Sensor:001
{"id":"urn:ngsi-ld:Sensor:001","type":"Sensor","Temperature":{"type":"Text","value":"30","metadata":{}}}

$ ngsi replace attrs --id urn:ngsi-ld:Sensor:001 --keyValues --data '{"Temperature":30}'

$ ngsi get entity --id urn:ngsi-ld:Sensor:001
{"id":"urn:ngsi-ld:Sensor:001","type":"Sensor","Temperature":{"type":"Number","value":30,"metadata":{}}}
```
