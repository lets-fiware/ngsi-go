# version - Convenience command

This command prints the version of Context Broker host specified by the `--host` option.

```console
ngsi version [options]
```

### Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

#### Example

```console
ngsi version -h orion
```

```json
{
 "orion" : {
   "version" : "2.5.0",
   "uptime" : "0 d, 5 h, 7 m, 50 s",
   "git_hash" : "63cc107657ae10aa03f1c83bdea0be869d8e26a1",
   "compile_time" : "Fri Oct 30 09:02:37 UTC 2020",
   "compiled_by" : "root",
   "compiled_in" : "320890801dd4",
   "release_date" : "Fri Oct 30 09:02:37 UTC 2020",
   "doc" : "https://fiware-orion.rtfd.io/en/2.5.0/",
   "libversions": {
      "boost": "1_53",
      "libcurl": "libcurl/7.29.0 NSS/3.44 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
      "libmicrohttpd": "0.9.70",
      "openssl": "1.0.2k",
      "rapidjson": "1.1.0",
      "mongodriver": "legacy-1.1.2"
   }
 }
}
```
