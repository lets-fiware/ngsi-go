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
  "version" : "3.3.1",
  "uptime" : "0 d, 0 h, 0 m, 1 s",
  "git_hash" : "a9ff9652c7b93240f48d2b497783407a80861370",
  "compile_time" : "Thu Nov 11 10:08:31 UTC 2021",
  "compiled_by" : "root",
  "compiled_in" : "831b4bc01053",
  "release_date" : "Thu Nov 11 10:08:31 UTC 2021",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.3.1/",
  "libversions": {
     "boost": "1_66",
     "libcurl": "libcurl/7.61.1 OpenSSL/1.1.1g zlib/1.2.11 nghttp2/1.33.0",
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
