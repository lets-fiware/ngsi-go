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
