# update - NGSI command

This command will update entities, attribute(s), or subscription.

-   [Update multiple entities](#update-multiple-entities)
-   [Update an attribute](#update-an-attribute)
-   [Update multiple attributes](#update-multiple-attributes)
-   [Update a subscription](#update-a-subscription)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="update-multiple-entities"/>

## Update multiple entities

This command updates multiple entities.

```console
ngsi update [common options] entities [options]
```

### Options

| Options                   | Description                          |
| ------------------------- | ------------------------------------ |
| --keyValues, -k           | specify keyValues (default: false)   |
| --data value, -d value    | specify data                         |
| --noOverwrite, -n         | specify noOverwrite (default: false) |
| --replace, -r             | specify replace (default: false)     |
| --link value, -L value    | specify @context                     |
| --help                    | show help (default: false)           |

### Example

```console
ngsi update entities
```

<a name="update-an-attribute"/>

## Update an attribute

This command updates an attribute.

```console
ngsi update [common options] attr [options]
```

### Options

| Options                | Description                |
| ---------------------- | -------------------------- |
| --id value, -i value   | specify id                 |
| --data value, -d value | specify data               |
| --attrName value       | specify attribute name     |
| --help                 | show help (default: false) |

### Example

```console
ngsi update attr
```

<a name="update-multiple-attributes"/>

## Update multiple attributes

This command updates multiple attributes.

```console
ngsi update [common options] attrs [options]
```

### Options

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --id value, -i value   | specify id                         |
| --type value, -t value | specfiy entity type                |
| --keyValues, -k        | specfiy keyValues (default: false) |
| --data value, -d value | specify data                       |
| --attrName value       | specify attribute name             |
| --help                 | show help (default: false)         |

### Example

```console
ngsi update attrs
```

<a name="update-a-subscription"/>

## Update a subscription

This command update a subscription.

```console
ngsi update [common options] subscription [options]
```

### Options

| Options                   | Description                                                |
| ------------------------- | ---------------------------------------------------------- |
| --id value, -i value      | specify id                                                 |
| --skipInitialNotification | specify skipInitialNotification (default: false)           |
| --data value, -d value    | specify data                                               |
| --uri value               | specify URL or URI                                         |
| --url value, -u value     | specify url to be invoked when a notification is generated |
| --expires value, -e value | specify expires                                            |
| --throttling value        | specify throttling (default: 0)                            |
| --keyValues, -k           | specify keyValues (default: false)                         |
| --query value, -q value   | specify query                                              |
| --link value, -L value    | specify @context                                           |
| --nAttrs value            | specify attributes to be notified                          |
| --wAttrs value            | specify watched attributes                                 |
| --description value       | specify description                                        |
| --get                     | specify get (default: false)                               |
| --status value            | specify status                                             |
| --subjectId value         | specify subjectId                                          |
| --idPattern value         | specify idPattern                                          |
| --type value, -t value    | specify Entity Type                                        |
| --typePattern value       | specify typePatternA                                       |
| --mq value, -m value      | specify mq                                                 |
| --georel value            | specify georel                                             |
| --geometry value          | specify geometry                                           |
| --coords value            | specify coords                                             |
| --headers value           | specify headers                                            |
| --qs value                | specify qs                                                 |
| --method value            | specify method                                             |
| --payload value           | specify payload                                            |
| --metadata value          | specify metadata                                           |
| --exceptAttrs value       | specify exceptAttrs                                        |
| --attrsFormat value       | specify attrsFormat                                        |
| --help                    | show help (default: false)                                 |

### Example

```console
ngsi update subscription --id 5fa78b70627088ba9b91b1c0 --expires 1day
```
