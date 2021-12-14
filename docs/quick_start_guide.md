# NGSI Go Quick Start Guide

## Add a broker

You register an alias to access the broker.

```console
ngsi broker add --host letsfiware --brokerHost http://localhost:1026 --ngsiType v2
```

## Get broker version by using an alias

You can get the version by using the alias `letsfiware`.

```console
ngsi version -h letsfiware
```

```json
{
"orion" : {
  "version" : "3.4.0",
  "uptime" : "0 d, 0 h, 0 m, 1 s",
  "git_hash" : "e8ed9c5f363ab503ac406a9e62a213640d9c2864",
  "compile_time" : "Tue Dec 14 09:34:09 UTC 2021",
  "compiled_by" : "root",
  "compiled_in" : "5221bc0800ef",
  "release_date" : "Tue Dec 14 09:34:09 UTC 2021",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.4.0/",
  "libversions": {
     "boost": "1_66",
     "libcurl": "libcurl/7.61.1 OpenSSL/1.1.1k zlib/1.2.11 nghttp2/1.33.0",
     "libmosquitto": "2.0.12",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.1",
     "rapidjson": "1.1.0",
     "mongoc": "1.17.4",
     "bson": "1.17.4"
  }
}
}
```

Once you access the broker, you can omit to specify the broker.

```console
ngsi version
```

If you want to check the current default settings, you can run the following command.

```console
ngsi settings list
```

## Create a entity

```console
ngsi create entity --keyValues \
--data ' {
      "id":"urn:ngsi-ld:Product:110",
      "type":"Product",
      "name": "Lemonade",
      "size": "S",
      "price": 99
}'
```

## Get a entity

```console
ngsi get entity --id urn:ngsi-ld:Product:110 --type Product
```

```json
{"id":"urn:ngsi-ld:Product:110","type":"Product","name":{"type":"Text","value":"Lemonade","metadata":{}},"price":{"type":"Number","value":99,"metadata":{}},"size":{"type":"Text","value":"S","metadata":{}}}
```

## Update attribute

```console
ngsi update attr --id urn:ngsi-ld:Product:110 --attr price --data 11
```

## Get a entity (keyValues)

```console
ngsi get entity --id urn:ngsi-ld:Product:110 --keyValues
```

```json
{"id":"urn:ngsi-ld:Product:110","name":"Lemonade","price":11,"size":"S","type":"Product"}
```

## Print number of entities

```console
ngsi wc entities --type Product
```

```text
10
```

## Delete a entity

```console
ngsi delete entity --id urn:ngsi-ld:Product:110
```

## Print number of entities again

```console
ngsi wc entities --type Product
```

```text
9
```
