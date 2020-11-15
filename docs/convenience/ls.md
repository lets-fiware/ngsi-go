# ls - Convenience command

This command lists multiple entities

```
ngsi ls [options]
```

### Options

| Options                       | Description                                 |
| ----------------------------- | ------------------------------------------- |
| --host value, -h value        | specify host or alias for source (Required) |
| --token value                 | specify oauth token                         |
| --service value, -s value     | specify FIWARE Service                      |
| --path value, -p value        | specify FIWARE ServicePath                  |
| --type value, -t value        | specify Entity Type                         |
| --idPattern value             | specify idPattern                           |
| --typePattern value           | specify typePattern                         |
| --query value, -q value       | specify query                               |
| --mq value, -m value          | specify mq                                  |
| --georel value                | specify georel                              |
| --geometry value              | specify geometry                            |
| --coords value                | specify coords                              |
| --attrs value                 | specify attrs                               |
| --metadata value              | specify metadata                            |
| --orderBy value               | specify orderBy                             |
| --count, -C                   | specify count (default: false)              |
| --keyValues, -k               | specify keyValues (default: false)          |
| --values, -V                  | specify values (default: false)             |
| --unique, -u                  | specify unique (default: false)             |
| --id value, -i value          | specify id                                  |
| --link value, -L value        | specify @context                            |
| --verbose, -v                 | specify verbose (default: false)            |
| --lines, -1                   | specify lines (default: false)              |
| --help                        | show help (default: false)                  |

### Example

#### Request:

```
$ ngsi ls --type Product
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

```
$ ngsi ls --type Product --count
14
```

#### Request:

```
$ ngsi ls --type Product --idPattern '0{2}'
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

```
$ ngsi ls --type Product --idPattern '1{2}'
urn:ngsi-ld:Product:110
urn:ngsi-ld:Product:111
urn:ngsi-ld:Product:112

$ ngsi ls --type Product --idPattern '1{2}' --count
3
```

#### Request:

```
$ ngsi ls --type Product --idPattern '1{2}' --verbose | jq .
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

```
$ ngsi ls --type Product --idPattern '1{2}' --verbose --keyValues | jq .
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

```
$ ngsi ls --type Product --idPattern '1{2}' --count
3

$ ngsi ls --type Product --idPattern '1{2}' | xargs -L 1 ngsi delete entity --id

$ ngsi ls --type Product --idPattern '1{2}' --count
0
```
