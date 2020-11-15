# version - Convenience command

This command prints the version of Context Broker host specified by the `--host` option.

```
ngsi version [options]
```

### Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

#### Example

```
$ ngsi version -h orion
{
"orion" : {
  "version" : "2.4.0-next",
  "uptime" : "0 d, 15 h, 1 m, 6 s",
  "git_hash" : "4f26834ca928e468b091729d93dabd22108a2690",
  "compile_time" : "Tue Mar 31 15:41:02 UTC 2020",
  "compiled_by" : "root",
  "compiled_in" : "d99d1dbb4d9e",
  "release_date" : "Tue Mar 31 15:41:02 UTC 2020",
  "doc" : "https://fiware-orion.rtfd.io/"
}
}
```
