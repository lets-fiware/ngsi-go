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
| --host value                   | host for tokenproxy (default: "0.0.0.0")    |
| --port value, -p value         | port for tokenproxy (default: "1029")       |
| --https, -s                    | start in https (default: false)             |
| --key value, -k value          | key file (only needed if https is enabled)  |
| --cert value, -c value         | cert file (only needed if https is enabled) |
| --idmHost value                | host for Keyrock                            |
| --clientId value, -I value     | specify client id for Keyrock               |
| --clientSecret value, -S value |  specify client secret for Keyrock          |
| --verbose, -v                  | verbose (default: false)                    |
| --help                         | show help (default: false)                  |

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

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --host value, -h value | regproxy host                  |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: false)     |

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
  "version": "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)",
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

