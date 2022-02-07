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
  "version" : "3.5.1",
  "uptime" : "0 d, 0 h, 0 m, 1 s",
  "git_hash" : "56593100c692ae943fdfbc14be5f27d0e6908adb",
  "compile_time" : "Sun Feb 6 21:02:25 UTC 2022",
  "compiled_by" : "root",
  "compiled_in" : "bf3518e8dd0b",
  "release_date" : "Sun Feb 6 21:02:25 UTC 2022",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.5.1/",
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
