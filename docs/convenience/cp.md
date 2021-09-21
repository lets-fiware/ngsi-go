# cp - Convenience command

This command copies multiple entities from a source to a destination.

```console
ngsi cp [options]
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)        |
| --host2 VALUE, -d VALUE   | host or alias (required)                      |
| --service VALUE, -s VALUE | FIWARE Service VALUE                          |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                      |
| --link VALUE, -L VALUE    | @context VALUE (LD)                           |
| --type VALUE, -t VALUE    | Entity Type (required)                        |
| --service2 VALUE          | FIWARE Service for destination                |
| --path2 VALUE             | FIWARE ServicePath for destination            |
| --context2 VALUE          | @context for destination                      |
| --ngsiV1                  | NGSI v1 mode (default: false)                 |
| --skipForwarding          | skip forwarding to CPrs (v2) (default: false) |
| --run                     | run command (default: false)                  |
| --help                    | show help (default: true)                     |

### Example

#### Request:

```console
ngsi cp --host orion1 --host2 orion2 --type EvacuationSpace --run
```

#### Request:

```
ngsi cp --run --host orion-ld --host2 orion-ld --service2 openiot --type TemperatureSensor --link ctx
```

#### Request:

```
ngsi cp --host orion --type TemperatureSensor --host2 orion-ld --context2 ctx --run
```
