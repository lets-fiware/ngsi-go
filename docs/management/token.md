# Token - Management command

This command gets an oauth token and prints a token information.

```console
ngsi token [options]
```

## Options

| Options                | Description                                 |
| ---------------------- | ------------------------------------------- |
| --host value, -h value | specify host or alias (Required)            |
| --verbose, -v          | print detailed information (default: false) |
| --pretty, -P           | pretty format (default: false)              |
| --expires, -e          | print expires (default: false)              |
| --revoke, -r           | revoke token (default: false)               |
| --help                 | show help (default: false)                  |

### Example 1

```console
ngsi token -h orion
```

```text
8385a04bd4e3d1da323843f32a18c9e0d5ad42e1
```

### Example 2

```console
export TOKEN=`ngsi token -h orion`
echo $TOKEN
63dcaf3e87d9775578b46a7bb046be74365e9b96
```

### Example 3

Get detailed information about a token

```console
ngsi token -h orion --verbose
```

```json
{"AccessToken":"7385a04bd4e3d1da723843f32a18c9e0d5ad42c9","ExpiresIn":3599,"RefreshToken":"59580f9a024ad8a28464e8be024b5c753dea2b9c","Scope":["bearer"],"TokenType":"Bearer"}
```

### Example 4

Get the number of seconds until a token expires

```console
ngsi token -h orion --expires
```

```text
2045
```
