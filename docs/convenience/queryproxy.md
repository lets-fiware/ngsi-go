# georoxy - Convenience command

This command allows you to manage queryproxy server that provides POST-based Query API Endpoint.

If you request too long URL, then it may give “414 Request URI too large” error message. The reason for those
"really long URLs" are the URI parameters such as `coords`, `q`, `attrs` for GET /v2/entities. The command
solves this problem by POST-based Query `POST /v2/ex/entities`.

-   [Server](#server)
-   [Sanity check](#sanity-check)

<a name="server"></a>

## Server

This command allows you to start up a queryproxy server.

```console
ngsi queryproxy server [options]
```

### Options

| Options                      | Description                                 |
| ---------------------------- | ------------------------------------------- |
| --host value, -h value       | context broker                              |
| --replaceURL value, -u value | replace URL (default: "/v2/ex/entities")    |
| --qhost value                | host for queryproxy (default: "0.0.0.0")    |
| --port value, -p value       | port for queryproxy (default: "1030")       |
| --https, -s                  | start in https (default: false)             |
| --key value, -k value        | key file (only needed if https is enabled)  |
| --cert value, -c value       | cert file (only needed if https is enabled) |
| --verbose, -v                | verbose (default: false)                    |
| --help                       | show help (default: false)                  |

### Example

```console
ngsi --stderr info queryproxy server \
 --host orion \
 --verbose
```

```
curl http://localhost:1030/v2/ex/entities --data "type=Device"
```

<a name="sanity-check"></a>

## Sanity check

This command allows you to check a queryproxy server healthy.

```console
ngsi queryproxy health [options]
```

### Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --host value, -h value | regproxy host                  |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |


### Example

```
ngsi server add --host queryproxy --serverType queryproxy --serverHost http://localhost:1030
```

```
ngsi queryproxy health --host queryproxy
```

```
{
  "ngsi-go": "queryproxy",
  "version": "0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)",
  "health": "OK",
  "orion": "http://orion:1026/v2/entities",
  "verbose": true,
  "uptime": "0 d, 0 h, 0 m, 1 s",
  "timesent": 0,
  "success": 0,
  "failure": 0
}
```
