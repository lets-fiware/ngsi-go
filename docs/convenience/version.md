# version - Convenience command

This command prints the version of Context Broker host specified by the `--host` option.

```console
ngsi version [options]
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi version -h orion
```

```json
{
"orion" : {
  "version" : "3.5.0",
  "uptime" : "0 d, 0 h, 0 m, 1 s",
  "git_hash" : "e86026dc05e0fec718f7d7fea6fcce4d58fc8d5e",
  "compile_time" : "Fri Jan 28 09:14:42 UTC 2022",
  "compiled_by" : "root",
  "compiled_in" : "999c711180e7",
  "release_date" : "Fri Jan 28 09:14:42 UTC 2022",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.5.0/",
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
