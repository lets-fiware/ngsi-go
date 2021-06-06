# cp - Convenience command

This command copies multiple entities from a source to a destination.

```console
ngsi cp [options]
```

## Options

| Options                   | Description                                      |
| ------------------------- | ------------------------------------------------ |
| --host value, -h value    | specify host or alias for source (Required)      |
| --host2 value, -d value   | specify host or alias (Required) for destination |
| --token value             | specify oauth token for source                   |
| --service value, -s value | specify FIWARE Service for source                |
| --path value, -p value    | specify FIWARE ServicePath for source            |
| --type value, -t value    | specify Entity Type (Required)                   |
| --token2 value            | specify oauth token for destination              |
| --service2 value          | specify FIWARE Service for destination           |
| --path2 value             | specify FIWARE ServicePath for destination       |
| --run                     | actually run to copy entities (default: false)   |
| --context2 value          | specify @context for destination                 |
| --ngsiV1                  | NGSI v1 mode (default: false)                    |
| --help                    | show help (default: false)                       |

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
