# rm - Convenience command

This command removes all entities of an entity type specified by the `--type` option.


```console
ngsi rm [options]
```

### Options

| Options                       | Description                                    |
| ----------------------------- | ---------------------------------------------- |
| --host value, -h value        | specify host or alias (Required)               |
| --token value                 | specify oauth token                            |
| --service value, -s value     | specify FIWARE Service                         |
| --path value, -p value        | specify FIWARE ServicePath                     |
| --type value, -t value        | specify Entity Type (Required)                 |
| --run                         | actually run to copy entities (default: false) |
| --help                        | show help (default: false)                     |

#### Example

```console
ngsi rm --host orion --type EvacuationSpace --run
