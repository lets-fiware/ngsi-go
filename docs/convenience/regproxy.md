# regproxy - Convenience command

This command allows you to start up a registration proxy server. It's put between a context broker and
a Csource/CPr with a protected API endpoint.

```console
ngsi regproxy [options]
```

## Options

| Options                          | Description                                      |
| -------------------------------- | ------------------------------------------------ |
| --host value, -h value           | context broker or csource host                   |
| --rhost value                    | host for registration proxy (default: "0.0.0.0") |
| --port value, -p value           | port for registration proxy (default: "1028")    |
| --url value, -u value            | url for registration proxy (default: "/")        |
| --replaceService value, -S value | replace FIWARE-Serivce                           |
| --replacePath value, -P value    | replace FIWARE-SerivcePath                       |
| --replaceURL value, -U value     | replace URL of forwarding destination            |
| --https, -s                      | start in https (default: false)                  |
| --key value, -k value            | key file (only needed if https is enabled)       |
| --cert value, -c value           | cert file (only needed if https is enabled)      |
| --verbose, -v                    | verbose (default: false)                         |
| --help                           | show help (default: false)                       |

### Example

```console
ngsi --stderr info regproxy --host orion-with-keyrock --verbose
```
