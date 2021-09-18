# regproxy - Convenience command

This command allows you to start up a registration proxy server. It's put between a context broker and
a Csource/CPr with a protected API endpoint.

-   [Server](#server)
-   [Sanity check](#sanity-check)
-   [Configration](#config)

<a name="server"></a>

## Server

This command allows you to start up a registration proxy server.

```console
ngsi regproxy server [options]
```

## Options

| Options                | Description                                 |
| ---------------------- | ------------------------------------------- |
| --host VALUE, -h VALUE | context broker or csource host (required)   |
| --rhost VALUE          | host for registration proxy                 |
| --port VALUE, -p VALUE | port for registration proxy                 |
| --url VALUE, -u VALUE  | url for registration proxy                  |
| --replaceService VALUE | replace FIWARE-Serivce                      |
| --replacePath VALUE    | replace FIWARE-SerivcePath                  |
| --addPath VALUE        | add path to FIWARE-SerivcePath              |
| --replaceURL VALUE     | replace URL of forwarding destination       |
| --https, -s            | start in https (default: false)             |
| --key VALUE, -k VALUE  | key file (only needed if https is enabled)  |
| --cert VALUE, -c VALUE | cert file (only needed if https is enabled) |
| --verbose, -v          | verbose (default: false)                    |
| --help                 | show help (default: true)                   |

### Example

```console
ngsi --stderr info regproxy server --host orion-with-keyrock --verbose
```

<a name="sanity-check"></a>

## Sanity check

This command allows you to check a regproxy server healthy.

```console
ngsi regproxy health [options]
```

## Options

| Options                | Description                    |
| ---------------------- | ------------------------------ |
| --host VALUE, -h VALUE | regproxy host (required)       |
| --pretty, -P           | pretty format (default: false) |
| --help                 | show help (default: true)      |

### Example

```
ngsi server add --host regproxy --serverType regproxy --serverHost http://localhost:1028/
```

```
ngsi regproxy health --host regproxy --pretty
```

```
{
  "ngsi-go": "regproxy",
  "version": "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)",
  "health": "OK",
  "csource": "https://orion.letfiware.jp",
  "verbose": false,
  "uptime": "0 d, 1 h, 32 m, 44 s",
  "timesent": 5,
  "success": 4,
  "failure": 1
}
```

<a name="config"></a>

## Configration

This command allows you to change configration for a regproxy server.

```console
ngsi regproxy config [options]
```

## Options

| Options                   | Description                           |
| ------------------------- | ------------------------------------- |
| --host VALUE, -h VALUE    | regproxy host (required)              |
| --verbose VALUE, -v VALUE | verbose log (on/off)                  |
| --replaceService VALUE    | replace FIWARE-Serivce                |
| --replacePath VALUE       | replace FIWARE-SerivcePath            |
| --addPath VALUE           | add path to FIWARE-SerivcePath        |
| --replaceURL VALUE        | replace URL of forwarding destination |
| --pretty, -P              | pretty format (default: false)        |
| --help                    | show help (default: true)             |

### Example

```
ngsi regproxy config --host regproxy --verbose on --replacePath "/fiware" --replaceURL "" --pretty
```

```
{
  "verbose": true,
  "path": "/fiware"
}
```
