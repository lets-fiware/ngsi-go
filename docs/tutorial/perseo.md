# NGSI Go tutorial for Perseo

## Get settings of alias

```console
ngsi server get --host perseo
```

```json
serverType perseo
serverHost http://perseo:9090
```

## Print version

```console
ngsi version --host perseo --pretty
```

```json
{
  "error": null,
  "data": {
    "name": "perseo",
    "description": "IOT CEP front End",
    "version": "1.12.1"
  }
}
```

## List rules

```console
ngsi rules list --verbose --pretty

```json
[]
```

## Create rules

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
}'
```

## List rules

```console
ngsi rules list
```

```console
blood_rule_update
```

## Get rule

```console
ngsi rules get --name blood_rule_update --pretty
```

```json
{
  "error": null,
  "data": {
    "_id": "6024c00a8e2bfc0012c77486",
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

## Delete rule

```console
ngsi rules delete --name blood_rule_update
```

## Print number of rules

```console
ngsi rules list --count
```

```console
0
```
