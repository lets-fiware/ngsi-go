# providers - Keyrock command

This command prints service providers for Keyrock

## Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --host value, -h value | specify host or alias          |
| --token value          | specify oauth token            |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

### Examples

#### Request:

```console
ngsi providers --host keyrock --pretty
```

```console
{
  "information": {
    "total_users": 3,
    "total_organizations": 1,
    "total_applications": 5
  }
}
```
