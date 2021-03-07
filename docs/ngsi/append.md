# append - NGSI command

This command appends multiple attributes.

-   [Append multiple attributes](#append-multiple-attributes)
-   [append attribute instance of temporal entity](#append-attribute-instance-of-temporal-entity)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="append-multiple-attributes"></a>

## Append multiple attributes

This command appneds multiple attributes.

```console
ngsi append [common options] attrs [options]
```

### Options

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --id value, -i value      | specify id                         |
| --type value, -t value    | specify entity Type                |
| --keyValues, -k           | specify keyValues (default: false) |
| --append, -a              | specify append (default: false)    |
| --data value, -d value    | specify data                       |
| --link value, -L value    | specify @context                   |
| --context value, -C value | specify @context (LD)              |
| --help                    | show help (default: false)         |

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

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --id value, -i value      | specify id                         |
| --data value, -d value    | data                               |
| --link value, -L value    | @context (LD)                      |
| --context value, -C value | @context (LD)                      |
| --help                    | show help (default: false)         |
