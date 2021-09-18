# append - NGSI command

This command appends multiple attributes.

-   [Append multiple attributes](#append-multiple-attributes)
-   [append attribute instance of temporal entity](#append-attribute-instance-of-temporal-entity)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="append-multiple-attributes"></a>

## Append multiple attributes

This command appneds multiple attributes.

```console
ngsi append [common options] attrs [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --type VALUE, -t VALUE    | entity type                            |
| --keyValues, -K           | keyValues (default: false)             |
| --append, -a              | append (default: false)                |
| --data VALUE, -d VALUE    | attributes data                        |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Example

#### Request:

```console
ngsi append attrs --id urn:ngsi-ld:Product:001 \
--data '{
      "specialOffer":{"type": "Boolean", "value": true}
}'
```

#### Request:

```console
ngsi append attrs --id urn:ngsi-ld:Product:001 \
--keyValues --data '{"specialOffer":false}'
```

<a name="append-attribute-instance-of-temporal-entity"></a>

## Append attribute instance of temporal entity

This command appneds attribute instance of temporal entity.

```console
ngsi append [common options] tattrs [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | temporal entity id (required)          |
| --data VALUE, -d VALUE    | attributes instance of temporal entity |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |
