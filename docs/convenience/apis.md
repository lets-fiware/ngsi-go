# apis- Convenience command

This command prints the version of FIWARE GE specified by the `--host` option.

```console
ngsi version [options]
```

## Options

| Options                | Description                      |
| ---------------------- | -------------------------------- |
| --host value, -h value | specify host or alias (Required) |
| --token value          | specify oauth token              |
| --pretty, -P           | pretty format (default: false)   |
| --help                 | show help (default: false)       |

### Example

```console
ngsi version -h orion
```

```json
{
"orion" : {
  "version" : "3.0.0",
  "uptime" : "0 d, 0 h, 17 m, 19 s",
  "git_hash" : "d6f8f4c6c766a9093527027f0a4b3f906e7f04c4",
  "compile_time" : "Mon Apr 12 14:48:44 UTC 2021",
  "compiled_by" : "root",
  "compiled_in" : "f307ca0746f5",
  "release_date" : "Mon Apr 12 14:48:44 UTC 2021",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.0.0/",
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

```console
ngsi version --host comet
```

```json
{"version":"2.8.0-next"}
```

```console
ngsi version --host quantumleap
```

```json
{
  "version": "0.7.6"
}
```
