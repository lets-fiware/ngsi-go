# providers - Keyrock command

This command prints service providers for Keyrock

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

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
