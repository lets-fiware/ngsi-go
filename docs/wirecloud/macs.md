# macs - Application mashup command

This command manages mashable application components for WireCloud.

-   [List mashable application components](#list-macs)
-   [Get mashable application component](#get-mac)
-   [Download mashable application component](#download-mac)
-   [Install mashable application component](#install-mac)
-   [Uninstall mashable application component](#uninstall-mac)

## Common Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

<a name="list-macs"></a>

## List mashable application components

This command lists mashable application components.

```console
ngsi macs [options] list
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --widget                  | filtering widget (default: false)             |
| --operator                | filtering operator (default: false)           |
| --mashup                  | filtering mashup (default: false)             |
| --vender value, -v value  | vender name of mashable application component |
| --name value, -n value    | name of mashable application component        |
| --version value, -V value | version of mashable application component     |
| --json, -j                | JSON format (default: false)                  |
| --pretty, -P              | pretty format (default: false)                |

### Example

```console
ngsi macs --host wirecloud list
```

```json
CoNWeT/ngsientity2poi/3.2.2
CoNWeT/ngsi-source/4.0.0
CoNWeT/spy-wiring/1.0.3
FICODES/quantumleap-source/0.2.1
FISUDA/cesiumjs/0.3.0
FISUDA/deckgl/0.1.0
FISUDA/maplibregl/0.6.0
FISUDA/ol-map-ja/1.0.0
NGSIGO/test-widget/0.1.0
```

<a name="get-mac"></a>

## Get mashable application component

This command gets infomation of a mashable application component.

```console
ngsi macs [options] get
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --vender value, -v value  | vender name of mashable application component |
| --name value, -n value    | name of mashable application component        |
| --version value, -V value | version of mashable application component     |
| --pretty, -P              | pretty format (default: false)                |

### Example

```console
ngsi macs --host wirecloud get --vender FISUDA --name maplibregl --version 0.6.0
```

```console
ngsi macs --host wirecloud get FISUDA/maplibregl/0.6.0
```

<a name="download-mac"></a>

## Download mashable application component

This command downloads a mashable application component.

```console
ngsi macs [options] downloads
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --vender value, -v value  | vender name of mashable application component |
| --name value, -n value    | name of mashable application component        |
| --version value, -V value | version of mashable application component     |
| --pretty, -P              | pretty format (default: false)                |

### Example

```console
ngsi macs --host wirecloud download --vender FISUDA --name maplibregl --version 0.6.0
```

```console
ngsi macs --host wirecloud download FISUDA/maplibregl/0.6.0
```

<a name="install-mac"></a>

## Install a mashable application component

This command installs a mashable application component.

```console
ngsi macs [options] install 
```

## Options

| Options                | Description                                                       |
| ---------------------- | ----------------------------------------------------------------- |
| --file value, -f value | mashable application component file                               |
| --public, -p           | install mashable application component as public (default: false) |
| --overwrite, -o        | overwrite mashable application component (default: false)         |
| --json, -j             | JSON format (default: false)                                      |
| --pretty, -P           | pretty format (default: false)                                    |

### Example

```console
ngsi macs --host wirecloud install --file NGSIGO_test-widget_0.2.0.wgt --overwirte
```

```console
ngsi macs --host wirecloud install --overwirte NGSIGO_test-widget_0.2.0.wgt
```

<a name="uninstall-mac"></a>

## Uninstall mashable application component

This command uninstalls mashable application component(s).

```console
ngsi macs [options] uninstall 
```

## Options

| Options                   | Description                                   |
| ------------------------- | --------------------------------------------- |
| --vender value, -v value  | vender name of mashable application component |
| --name value, -n value    | name of mashable application component        |
| --version value, -V value | version of mashable application component     |
| --run                     | run command (default: false)                  |
| --json, -j                | JSON format (default: false)                  |
| --pretty, -P              | pretty format (default: false)                |

### Example

```console
ngsi macs --host wirecloud uninstall --vender NGSIGO --name test-widget --version 0.2.0 --pretty --run
```

```console
ngsi macs --host wirecloud uninstall --pretty --run NGSIGO/test-widget/0.2.0
```

```console
ngsi macs --host wirecloud uninstall --vender NGSIGO --name test-widget --pretty --run
```

```console
ngsi macs --host wirecloud uninstall --pretty --run NGSIGO/test-widget
```

```
{
  "affectedVersions": [
    "0.2.0"
  ]
}
```

```console
 ngsi macs --host wirecloud uninstall --pretty --run NGSIGO/test-widget
```

```
{
  "affectedVersions": [
    "0.1.0",
    "0.2.0",
    "0.3.0"
  ]
}
```
