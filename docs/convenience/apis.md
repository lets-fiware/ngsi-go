# apis- Convenience command

This command prints the version of FIWARE GE specified by the `--host` option.

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
