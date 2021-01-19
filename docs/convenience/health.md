# health - Convenience command

This command prints the health status of FIWARE GE specified by the `--host` option.

```console
ngsi health [options]
```

## Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

### Example

```console
ngsi health --host quantumleap
```

```json
{
  "status": "pass"
}
```

```console
ngsi health --host quantumleap
```

```json
healthCheck004 error 503 SERVICE UNAVAILABLE {
  "details": {
    "crateDB": "cannot reach crate"
  },
  "status": "fail"
}
```
