# receiver - notification receiver

This command can receive notifications related with subscriptions that context broker send.

```console
ngsi receiver [options]
```

### Options

| Options                | Description                                 |
| ---------------------- | ------------------------------------------- |
| --port value, -p value | specify port for receiver (default: "1028") |
| --pretty, -P           | pretty format (default: false)              |
| --verbose, -v          | verbose (default: false)                    |
| --help                 | show help (default: false)                  |

### Example

```console
ngsi rm --host orion --type EvacuationSpace --run
```

```
{
  "data": [
    {
      "id": "device001",
      "temperature": {
        "metadata": {},
        "type": "Number",
        "value": 24
      },
      "type": "device"
    }
  ],
  "subscriptionId": "5fd412e8ecb082767349b975"
}
```

### Use case

#### Start up a receiver

Run `ngsi receiver` command.

```console
ngsi receiver --pretty
```

Open another terminal and run the following commands on it.

#### Create an entity

```console
ngsi create --host orion entity --data '{"type": "device", "id": "device001", "temperature": 26}'
```

#### Create a subscription

```console
ngsi create --host orin subscription --idPattern ".*" --url http://192.168.1.1:1028/
```

```console
5fd412e8ecb082767349b975
```

#### Update an attribute value

```console
ngsi update --host orion attr --id device001 --attrName temperature --data 22
```

#### Notification message

You will find the following message on the terminal you run `ngsi receiver` command.

```json
{
  "data": [
    {
      "id": "device001",
      "temperature": {
        "metadata": {},
        "type": "Number",
        "value": 22
      },
      "type": "device"
    }
  ],
  "subscriptionId": "5fd412e8ecb082767349b975"
}
```
