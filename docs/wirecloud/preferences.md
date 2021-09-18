# preferences - Application mashup command

This command manages preferences for WireCloud.

-   [Get preferences](#get-preferences)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="get-preferences"></a>

## Get preferences

This command gets preferences.

```console
ngsi preferences [options] get
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi preferences --host wirecloud get --pretty
```

```json
{
  "language": {
    "inherit": false,
    "value": "default"
  }
}
```
