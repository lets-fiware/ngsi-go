# rm - Convenience command

This command removes all entities of an entity type specified by the `--type` option.

```console
ngsi rm [options]
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)        |
| --service VALUE, -s VALUE | FIWARE Service VALUE                          |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                      |
| --type VALUE, -t VALUE    | Entity Type (required)                        |
| --link VALUE, -L VALUE    | @context VALUE (LD)                           |
| --ngsiV1                  | NGSI v1 mode (default: false)                 |
| --skipForwarding          | skip forwarding to CPrs (v2) (default: false) |
| --run                     | run command (default: false)                  |
| --help                    | show help (default: true)                     |

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
