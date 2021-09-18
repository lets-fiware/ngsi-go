# workspaces - Application mashup command

This command manages workspaces for WireCloud.

-   [List workspaces](#list-workspaces)
-   [Get workspace](#get-workspace)

## Common Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --help                 | show help (default: true)              |

<a name="list-workspaces"></a>

## List workspaces

This command lists workspaces.

```console
ngsi workspaces [options] list
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --json, -j             | JSON format (default: false)           |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi workspaces --host wirecloud list
```

```json
1 fiware FIWARE 2021/04/10 10:42:29
2 maplibre maplibre 2021/05/11 20:33:19
3 debug debug 2021/04/25 08:16:19
```

<a name="get-workspace"></a>

## Get a workspace

This command gets a workspace.

```console
ngsi workspaces [options] get
```

## Options

| Options                | Description                            |
| ---------------------- | -------------------------------------- |
| --host VALUE, -h VALUE | broker or server host VALUE (required) |
| --wid VALUE, -w VALUE  | workspace id (required)                |
| --users, -u            | list users (default: false)            |
| --tabs, -t             | list tabs (default: false)             |
| --widgets, -W          | list widgets (default: false)          |
| --operators, -o        | list operators (default: false)        |
| --json, -j             | JSON format (default: false)           |
| --pretty, -P           | pretty format (default: false)         |
| --help                 | show help (default: true)              |

### Example

```console
ngsi workspaces --host wc get --wid 1 --pretty
```

```json
{
  "id": "1",
  "name": "fiware",
  "title": "FIWARE",
  "public": false,
  "shared": false,
  "requireauth": false,
  "owner": "admin",
  "removable": true,
  "lastmodified": 1625019840620,
  "description": "",
  "longdescription": "",
  "preferences": {
    "public": {
      "inherit": false,
      "value": "false"
    },
    "requireauth": {
      "inherit": false,
      "value": "false"
    },
    "sharelist": {
      "inherit": false,
      "value": "[]"
    },
    "initiallayout": {
      "inherit": false,
      "value": "Fixed"
    },
    "baselayout": {
      "inherit": false,
      "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"
    }
  },
  "users": [
    {
      "fullname": "",
      "username": "admin",
      "organization": false,
      "accesslevel": "owner"
    }
  ],
  "groups": [],
  "empty_params": [],
  "extra_prefs": [],
  "tabs": [
    {
      "id": "40",
      "name": "tab",
      "title": "Tab",
      "visible": true,
      "preferences": {
        "requireauth": {
          "inherit": true,
          "value": "false"
        },
        "sharelist": {
          "inherit": true,
          "value": "[]"
        },
        "initiallayout": {
          "inherit": true,
          "value": "Fixed"
        },
        "baselayout": {
          "inherit": true,
          "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"
        }
      },
      "iwidgets": []
    }
  ],
  "wiring": {
    "version": "2.0",
    "connections": [],
    "operators": {},
    "visualdescription": {
      "behaviours": [],
      "components": {
        "operator": {},
        "widget": {}
      },
      "connections": []
    }
  }
}
```

