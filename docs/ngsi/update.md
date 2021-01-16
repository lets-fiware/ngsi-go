# update - NGSI command

This command will update entities, attribute(s), or subscription.

-   [Update multiple entities](#update-multiple-entities)
-   [Update an attribute](#update-an-attribute)
-   [Update multiple attributes](#update-multiple-attributes)
-   [Update a subscription](#update-a-subscription)

## Common Options

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
| --context value, -C value | specify @context (LD)                |
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

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --id value, -i value      | specify id                 |
| --data value, -d value    | specify data               |
| --attrName value          | specify attribute name     |
| --link value, -L value    | specify @context           |
| --context value, -C value | specify @context (LD)      |
| --help                    | show help (default: false) |

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

| Options                   | Description                        |
| ------------------------- | ---------------------------------- |
| --id value, -i value      | specify id                         |
| --type value, -t value    | specfiy entity type                |
| --keyValues, -k           | specfiy keyValues (default: false) |
| --data value, -d value    | specify data                       |
| --link value, -L value    | specify @context                   |
| --context value, -C value | specify @context (LD)              |
| --help                    | show help (default: false)         |

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

| Options                   | Description                                                    |
| ------------------------- | -------------------------------------------------------------- |
| --id value, -i value      | speciy id                                                      |
| --data value, -d value    | specify data                                                   |
| --skipInitialNotification | specify skipInitialNotification (default: false)               |
| --subscriptionId value    | specify subscription id (LD)                                   |
| --name value              | specify subscription name (LD)                                 |
| --description value       | specify description                                            |
| --entityId value          | specify entityId (LD)                                          |
| --idPattern value         | specify idPattern                                              |
| --type value, -t value    | specify Entity Type                                            |
| --typePattern value       | specify typePattern (v2)                                       |
| --wAttrs value            | specify watched attributes                                     |
| --timeInterval value      | specify time interval (LD) (default: 0)                        |
| --query value, -q value   | specify query                                                  |
| --mq value, -m value      | specify mq (v2)                                                |
| --geometry value          | specify geometry                                               |
| --coords value            | specify coords                                                 |
| --georel value            | specify georel                                                 |
| --geoproperty value       | sprcify geoproperty (LD)                                       |
| --csf value               | specify context source filter (LD)                             |
| --active                  | specify active (LD) (default: false)                           |
| --inactive                | specify inactive (LD) (default: false)                         |
| --nAttrs value            | specify attributes to be notified                              |
| --keyValues, -k           | specify keyValues (default: false)                             |
| --uri value               | specify uri/url to be invoked when a notification is generated |
| --accept value            | specify accept header (json or ld+json)                        |
| --expires value, -e value | specify expires                                                |
| --throttling value        | specify throttling (default: 0)                                |
| --timeRel value           | specify temporal relationship (LD)                             |
| --timeAt value            | specify timeAt (LD)                                            |
| --endTimeAt value         | specify endTimeAt (LD)                                         |
| --timeProperty value      | specify timeProperty (LD)                                      |
| --link value, -L value    | specify @context (LD)                                          |
| --context value, -C value | specify @context (LD)                                          |
| --status value            | specify status                                                 |
| --headers value           | specify headers (v2)                                           |
| --qs value                | specify qs (v2)                                                |
| --method value            | specify method (v2)                                            |
| --payload value           | specify payload (v2)                                           |
| --metadata value          | specify metadata (v2)                                          |
| --exceptAttrs value       | specify exceptAttrs (v2)                                       |
| --attrsFormat value       | specify attrsFormat (v2)                                       |
| --safeString value        | use safe string (value: on/off)                                |
| --help                    | show help (default: false)                                     |

### Example

```console
ngsi update subscription --id 5fa78b70627088ba9b91b1c0 --expires 1day
```
