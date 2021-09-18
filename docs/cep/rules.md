# rules - Perseo command

This command allows you to list, create, get and delete rule entry for Perseo.

-   [List all rules](#list-all-rules)
-   [Create a rule](#create-a-rule)
-   [Get a rule](#create-a-get-rule)
-   [Delete a rule](#delete-a-rule)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="list-all-rule"></a>

## List all rules

This command lists all rules.

```console
ngsi rules [command options] list [options]
```

### Options

| Options                   | Description                                             |
| ------------------------- | ------------------------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                  |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                    |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                                |
| --limit VALUE             | maximum number of rules                                 |
| --offset VALUE            | offset to skip a given number of rules at the beginning |
| --count                   | print number of rules (default: false)                  |
| --raw                     | print raw data (default: false)                         |
| --verbose, -v             | verbose (default: false)                                |
| --pretty, -P              | pretty format (default: false)                          |
| --help                    | show help (default: true)                               |

### Examples

#### Request:

```console
ngsi rules list
```

```console
blood_rule_update
```

<a name="create-a-rule"></a>

## Create a rule

This command creates a rule.

```console
ngsi rules [command options] create [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | rule data (required)                   |
| --verbose, -v             | verbose (default: false)               |
| --pretty, -P              | pretty format (default: false)         |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi rules --host perseo create \
--data '{
    "name": "blood_rule_update",
    "text": "select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]",
    "action": {
        "type": "update",
        "parameters": {
            "attributes": [
                {
                    "name": "abnormal",
                    "value": "true",
                    "type": "boolean"
                }
            ]
        }
    }
}
```

<a name="get-a-rule"></a>

## Get a rule

This command gets a rule.

```console
ngsi rules [command options] get [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --name VALUE, -n VALUE    | rule name (required)                   |
| --pretty, -P              | pretty format (default: false)         |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi rules get --name blood_rule_update --pretty
```

```json
{
  "error": null,
  "data": {
    "_id": "6024c5af8e2bfc0012c77487",
    "name": "blood_rule_update",
    "text": "select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]",
    "action": {
      "type": "update",
      "parameters": {
        "attributes": [
          {
            "name": "abnormal",
            "value": "true",
            "type": "boolean"
          }
        ]
      }
    },
    "subservice": "/",
    "service": "unknownt"
  }
}
```

<a name="delete-a-rule"></a>

## Delete a rule

This command deletes a rule.

```console
ngsi rules [command options] delete [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --name VALUE, -n VALUE    | rule name (required)                   |
| --help                    | show help (default: true)              |

### Examples

#### Request:

```console
ngsi rules delete --name blood_rule_update
```
