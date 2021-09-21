# receiver - Convenience command

This command can receive notifications related with subscriptions that a context broker sends.

```console
ngsi receiver [options]
```

## Options

| Options                | Description                                 |
| ---------------------- | ------------------------------------------- |
| --host VALUE, -h VALUE | host for receiver                           |
| --port VALUE, -p VALUE | port for receiver                           |
| --url VALUE, -u VALUE  | url for receiver                            |
| --pretty, -P           | pretty format (default: false)              |
| --https, -s            | start in https (default: false)             |
| --key VALUE, -k VALUE  | key file (only needed if https is enabled)  |
| --cert VALUE, -c VALUE | cert file (only needed if https is enabled) |
| --verbose, -v          | verbose (default: false)                    |
| --header               | print receive header (default: false)       |
| --help                 | show help (default: true)                   |

### Example

```console
ngsi receiver --verbose
```

```json
{
  "subscriptionId": "5fd412e8ecb082767349b975",
  "data": [
    {
      "id": "device001",
      "type": "device",
      "temperature": {
        "type": "Number",
        "value": 21,
        "metadata": {}
      }
    }
  ]
}
```

### Example - https mode

Make a key file and a cert file.

```console
openssl genrsa 2048 > myself.key
openssl req -new -key myself.key > myself.csr
openssl x509 -days 3650 -req -signkey myself.key < myself.csr > myself.crt
```

Start up a receiver in https mode.

```console
ngsi receiver --https --key myself.key --cert myself.crt
```

### Use case

#### Start up a receiver

Run `ngsi receiver` command with --pretty option.

```console
ngsi receiver --pretty
```

Open another terminal and run the following commands on it.

#### Create an entity

```console
ngsi create --host orion entity --keyValues \
--data '{"type": "device", "id": "device001", "temperature": 26}'
```

#### Create a subscription

```console
ngsi create --host orion subscription --idPattern ".*" --url http://192.168.1.1:1028/
```

```console
5fd412e8ecb082767349b975
```

#### Update an attribute value

```console
ngsi update --host orion attr --id device001 --attr temperature --data 22
```

#### Notification message

You will find the following message on the terminal that you ran `ngsi receiver` command.

```json
{
  "subscriptionId": "5fd412e8ecb082767349b975",
  "data": [
    {
      "id": "device001",
      "type": "device",
      "temperature": {
        "type": "Number",
        "value": 21,
        "metadata": {}
      }
    }
  ]
}
```

#### Print the subscription

```console
ngsi get subscription --id 5fd412e8ecb082767349b975 --pretty
```

```json
{
  "id": "5fd412e8ecb082767349b975",
  "subject": {
    "entities": [
      {
        "idPattern": ".*"
      }
    ],
    "condition": {}
  },
  "notification": {
    "timesSent": 2,
    "lastNotification": "2020-12-12T01:12:13.000Z",
    "lastSuccess": "2020-12-12T01:12:13.000Z",
    "lastSuccessCode": 204,
    "onlyChangedAttrs": false,
    "http": {
      "url": "http://192.168.1.1:1028/"
    },
    "attrsFormat": "normalized"
  },
  "status": "active"
}
```
