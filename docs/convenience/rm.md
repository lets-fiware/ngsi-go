# rm - Convenience command

This command removes all entities of an entity type specified by the `--type` option.

```console
ngsi rm [options]
```

## Options

| Options                   | Description                                         |
| ------------------------- | --------------------------------------------------- |
| --host value, -h value    | specify host or alias (Required)                    |
| --token value             | specify oauth token                                 |
| --service value, -s value | specify FIWARE Service                              |
| --path value, -p value    | specify FIWARE ServicePath                          |
| --type value, -t value    | specify Entity Type (Required)                      |
| --link value, -L value    | specify @context (LD)                               |
| --ngsiV1                  | NGSI v1 mode (default: false)                       |
| --skipForwarding          | skip forwarding to CPrs (v2) (Orion 3.1.0 or later) |
| --run                     | actually run to copy entities (default: false)      |
| --help                    | show help (default: false)                          |

### Example

```console
ngsi rm --host orion --type EvacuationSpace --run
```

```console
ngsi rm --type Device,Event,Thing --run 
```

```console
ngsi rm --type AEDFacilities --ngsiV1 --run 
```
