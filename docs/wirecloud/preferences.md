# preferences - Application mashup command

This command manages preferences for WireCloud.

-   [Get preferences](#get-preferences)

## Common Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

<a name="get-preferences"></a>

## Get preferences

This command gets preferences.

```console
ngsi preferences [options] get
```

## Options

| Options      | Description                    |
| ------------ | ------------------------------ |
| --pretty, -P | pretty format (default: false) |

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
