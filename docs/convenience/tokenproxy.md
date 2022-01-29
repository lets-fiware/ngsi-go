# regproxy - Convenience command

This command allows you to manage tokenproxy server that provides auth token from Keyrock.

-   [Server](#server)
-   [Sanity check](#sanity-check)

<a name="server"></a>

## Server

This command allows you to start up a tokenproxy server.

```console
ngsi tokenproxy server [options]
```

## Options

| Options                        | Description                                 |
| ------------------------------ | ------------------------------------------- |
| --host VALUE                   | host for tokenproxy                         |
| --port VALUE, -p VALUE         | port for tokenproxy                         |
| --https, -s                    | start in https (default: false)             |
| --key VALUE, -k VALUE          | key file (only needed if https is enabled)  |
| --cert VALUE, -c VALUE         | cert file (only needed if https is enabled) |
| --idmHost VALUE                | host for Keyrock                            |
| --clientId VALUE, -I VALUE     | client id for Keyrock                       |
| --clientSecret VALUE, -S VALUE | client secret for Keyrock                   |
| --verbose, -v                  | verbose (default: false)                    |
| --help                         | show help (default: true)                   |

### Example

```console
ngsi --stderr info tokenproxy server \
 --idmHost http://keyrock:3000 \
 --clientId a1a6048b-df1d-4d4f-9a08-5cf836041d14" \
 --clientSecret e4cc0147-e38f-4211-b8ad-8ae5e6a107f9 \
 --verbose
```

<a name="sanity-check"></a>

## Sanity check

This command allows you to check a tokenproxy server healthy.

```console
ngsi tokenproxy health [options]
```

## Options

| Options                | Description                       |
| ---------------------- | --------------------------------- |
| --host VALUE, -h VALUE | tokenproxy server host (required) |
| --pretty, -P           | pretty format (default: false)    |
| --help                 | show help (default: true)         |

### Example

```
ngsi server add --host tokenproxy --serverType tokenproxy --serverHost http://0.0.0.0:1029
```

```
ngsi tokenproxy health --host tokenproxy --pretty
```

```
{
  "ngsi-go": "tokenproxy",
  "version": "0.11.0 (git_hash:a7da56ae829c3204e31aa0c82ed1d5cca2a37ef9)",
  "health": "OK",
  "idm": "http://keyrock:3000/oauth2/token",
  "clientId": "a1a6048b-df1d-4d4f-9a08-5cf836041d14",
  "clientSecret": "e4cc0147-e38f-4211-b8ad-8ae5e6a107f9",
  "verbose": true,
  "uptime": "0 d, 0 h, 0 m, 1 s",
  "timesent": 0,
  "success": 0,
  "revoke": 0,
  "failure": 0
}
```

<a name="example"></a>

## Example

### How to get a token

```
curl http://localhost:1029/token \
  --data "username=admin@test.com" \
  --data "password=1234"
```

```
{
  "access_token": "41cc7caa5f2ba2da9b250273b2445c0c5c3cb3d6",
  "token_type": "bearer",
  "expires_in": 3599,
  "refresh_token": "6fd8edc885580f50a9c37a1e43f505a68f40b6a2",
  "scope": [
    "bearer"
  ]
}
```

### How to revoke a token

```
curl http://localhost:1029/revoke \
  --data "token=6fd8edc885580f50a9c37a1e43f505a68f40b6a2"
```

