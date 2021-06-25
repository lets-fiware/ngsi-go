# tabs - Application mashup command

This command manages tabs of a workspace for WireCloud.

-   [List tabs](#list-tabs)
-   [Get tab](#get-tab)

## Common Options

| Options                         | Description                      |
| ------------------------------- | -------------------------------- |
| --host value, -h value          | specify host or alias (Required) |
| --token value                   | specify oauth token              |
| --help                          | show help (default: false)       |

<a name="list-tabs"></a>

## List tabs

This command lists tabs of a workspace.

```console
ngsi tabs [options] list
```

## Options

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --wid value, -w value | workspace id                   |
| --json, -j            | JSON format (default: false)   |
| --pretty, -P          | pretty format (default: false) |

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

| Options               | Description                    |
| --------------------- | ------------------------------ |
| --wid value, -w value | workspace id                   | 
| --tid value, -t value | tab id                         |
| --json, -j            | JSON format (default: false)   |
| --pretty, -P          | pretty format (default: false) |

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
