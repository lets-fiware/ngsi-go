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
   "version" : "2.5.2",
   "uptime" : "0 d, 13 h, 54 m, 48 s",
   "git_hash" : "11e4cbfef30d28347162e5c4ef4de3a5d2797c69",
   "compile_time" : "Thu Dec 17 08:43:46 UTC 2020",
   "compiled_by" : "root",
   "compiled_in" : "5a4a8800b1fa",
   "release_date" : "Thu Dec 17 08:43:46 UTC 2020",
   "doc" : "https://fiware-orion.rtfd.io/en/2.5.2/",
   "libversions": {
      "boost": "1_53",
      "libcurl": "libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
      "libmicrohttpd": "0.9.70",
      "openssl": "1.0.2k",
      "rapidjson": "1.1.0",
      "mongodriver": "legacy-1.1.2"
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
