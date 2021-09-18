# ls - Convenience command

This command lists multiple entities

```console
ngsi ls [options]
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)        |
| --service VALUE, -s VALUE | FIWARE Service VALUE                          |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                      |
| --id VALUE, -i VALUE      | entity id                                     |
| --type VALUE, -t VALUE    | entity type                                   |
| --idPattern VALUE         | idPattern                                     |
| --typePattern VALUE       | typePattern (v2)                              |
| --query VALUE, -q VALUE   | filtering by attribute value                  |
| --mq VALUE, -m VALUE      | filtering by metadata (v2)                    |
| --georel VALUE            | georel                                        |
| --geometry VALUE          | geometry                                      |
| --coords VALUE            | coords                                        |
| --attrs VALUE             | attributes                                    |
| --metadata VALUE          | metadata (v2)                                 |
| --orderBy VALUE           | orderBy                                       |
| --count, -C               | count (default: false)                        |
| --keyValues, -K           | keyValues (default: false)                    |
| --values, -V              | values (default: false)                       |
| --unique, -U              | unique (default: false)                       |
| --skipForwarding          | skip forwarding to CPrs (v2) (default: false) |
| --link VALUE, -L VALUE    | @context VALUE (LD)                           |
| --verbose, -v             | verbose (default: false)                      |
| --lines, -1               | lines (default: false)                        |
| --pretty, -P              | pretty format (default: false)                |
| --safeString VALUE        | use safe string (VALUE: on/off)               |
| --help                    | show help (default: true)                     |

### Example

#### Request:

```console
ngsi ls --type Product
```

```text
urn:ngsi-ld:Product:001
urn:ngsi-ld:Product:002
urn:ngsi-ld:Product:003
urn:ngsi-ld:Product:004
urn:ngsi-ld:Product:005
urn:ngsi-ld:Product:006
urn:ngsi-ld:Product:007
urn:ngsi-ld:Product:008
urn:ngsi-ld:Product:009
urn:ngsi-ld:Product:010
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112
urn:ngsi-ld:Product:101
```

#### Request:

```console
ngsi ls --type Product --count
```

```text
14
```

#### Request:

```console
ngsi ls --type Product --idPattern '0{2}'
```

```text
urn:ngsi-ld:Product:001
urn:ngsi-ld:Product:002
urn:ngsi-ld:Product:003
urn:ngsi-ld:Product:004
urn:ngsi-ld:Product:005
urn:ngsi-ld:Product:006
urn:ngsi-ld:Product:007
urn:ngsi-ld:Product:008
urn:ngsi-ld:Product:009
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}'
```

```text
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' --count
```

```text
3
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' --verbose --pretty
[
  {
    "id": "urn:ngsi-ld:Product:110",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Lemonade"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 99
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "S"
    },
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:111",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Brandy"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 1199
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "M"
    },
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:112",
    "name": {
      "metadata": {},
      "type": "Text",
      "value": "Port"
    },
    "price": {
      "metadata": {},
      "type": "Number",
      "value": 1099
    },
    "size": {
      "metadata": {},
      "type": "Text",
      "value": "M"
    },
    "type": "Product"
  }
]
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' --verbose --keyValues --pretty
[
  {
    "id": "urn:ngsi-ld:Product:110",
    "name": "Lemonade",
    "price": 99,
    "size": "S",
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:111",
    "name": "Brandy",
    "price": 1199,
    "size": "M",
    "type": "Product"
  },
  {
    "id": "urn:ngsi-ld:Product:112",
    "name": "Port",
    "price": 1099,
    "size": "M",
    "type": "Product"
  }
]
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' --count
```

```text
3
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' | xargs -L 1 ngsi delete entity --id
```

#### Request:

```console
ngsi ls --type Product --idPattern '1{2}' --count
```

```text
0
```
