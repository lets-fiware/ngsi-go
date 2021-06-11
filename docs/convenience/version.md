# version - Convenience command

This command prints the version of Context Broker host specified by the `--host` option.

```console
ngsi version [options]
```

## Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

### Example

```console
ngsi version -h orion
```

```json
{
"orion" : {
  "version" : "3.1.0",
  "uptime" : "0 d, 0 h, 0 m, 10 s",
  "git_hash" : "260505c911ecf204ebcf0bd31788013c225da6dd",
  "compile_time" : "Wed Jun 9 12:59:59 UTC 2021",
  "compiled_by" : "root",
  "compiled_in" : "e11ec65d5407",
  "release_date" : "Wed Jun 9 12:59:59 UTC 2021",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.1.0/",
  "libversions": {
     "boost": "1_66",
     "libcurl": "libcurl/7.61.1 OpenSSL/1.1.1g zlib/1.2.11 nghttp2/1.33.0",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.1",
     "rapidjson": "1.1.0",
     "mongoc": "1.17.4",
     "bson": "1.17.4"
  }
}
}
```
