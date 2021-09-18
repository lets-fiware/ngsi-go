# tabs - Application mashup command

This command manages tabs of a workspace for WireCloud.

-   [List tabs](#list-tabs)
-   [Get tab](#get-tab)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-tabs"></a>

## List tabs

This command lists tabs of a workspace.

```console
ngsi tabs [options] list
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --wid VALUE, -w VALUE  | workspace id (required)                |
| --json, -j             | JSON format (default: false)           |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi tabs --host wirecloud list --wid 1
```

```json
40 tab Tab
41 tab-2 Tab 2
42 tab-3 Tab 3
```

<a name="get-tab"></a>

## Get a tab

This command gets a tab of workspace.

```console
ngsi tabs [options] get
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --wid VALUE, -w VALUE  | workspace id (required)                |
| --tid VALUE, -t VALUE  | tab id (required)                      |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi tabs get --wid 1 --tid 40 --pretty
```

```json
{
  "id": "40",
  "iwidgets": [],
  "name": "tab",
  "preferences": {
    "baselayout": {
      "inherit": true,
      "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"
    },
    "initiallayout": {
      "inherit": true,
      "value": "Fixed"
    },
    "requireauth": {
      "inherit": true,
      "value": "false"
    },
    "sharelist": {
      "inherit": true,
      "value": "[]"
    }
  },
  "title": "Tab",
  "visible": true
}
```
