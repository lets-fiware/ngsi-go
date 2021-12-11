# update - NGSI command

This command will update entities, attribute(s), or subscription.

-   [Update multiple entities](#update-multiple-entities)
-   [Update an attribute](#update-an-attribute)
-   [Update an attribute instance of temporal entity](#update-an-attribute-instance-of-temporal-entity)
-   [Update multiple attributes](#update-multiple-attributes)
-   [Update a subscription](#update-a-subscription)

## Common Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --help                    | show help (default: true)              |

<a name="update-multiple-entities"></a>

## Update multiple entities

This command updates multiple entities.

```console
ngsi update [common options] entities [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --data VALUE, -d VALUE    | entities data (required)               |
| --keyValues, -K           | keyValues (default: false)             |
| --noOverwrite, -n         | noOverwrite (default: false)           |
| --replace, -r             | replace (default: false)               |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --help                    | show help (default: true)              |

### Example

```console
ngsi update entities
```

<a name="update-an-attribute"></a>

## Update an attribute

This command updates an attribute.

```console
ngsi update [common options] attr [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --attr VALUE              | attribute name (required)              |
| --data VALUE, -d VALUE    | attribute data                         |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Example

```console
ngsi update attr
```

[Update attribute instance of temporal entity](#update-attribute-instance-of-temporal-entity)

<a name="update-an-attribute-instance-of-temporal-entity"></a>

## Update an attribute instance of temporal entity

This command updates an attribute instance of temporal entity.

```console
ngsi update [common options] tattr [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | temporal entity id (required)          |
| --attr VALUE              | attribute name (required)              |
| --instanceId VALUE        | attribute instance id (LD) (required)  |
| --data VALUE, -d VALUE    | attribute instance of temporal entity  |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

<a name="update-multiple-attributes"></a>

## Update multiple attributes

This command updates multiple attributes.

```console
ngsi update [common options] attrs [options]
```

### Options

| Options                   | Description                            |
| ------------------------- | -------------------------------------- |
| --host VALUE, -h VALUE    | broker or server host VALUE (required) |
| --service VALUE, -s VALUE | FIWARE Service VALUE                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE               |
| --id VALUE, -i VALUE      | entity id (required)                   |
| --type VALUE, -t VALUE    | entity type                            |
| --keyValues, -K           | keyValues (default: false)             |
| --data VALUE, -d VALUE    | attributes data                        |
| --link VALUE, -L VALUE    | @context VALUE (LD)                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                    |
| --safeString VALUE        | use safe string (VALUE: on/off)        |
| --help                    | show help (default: true)              |

### Example

```console
ngsi update attrs
```

<a name="update-a-subscription"></a>

## Update a subscription

This command update a subscription.

```console
ngsi update [common options] subscription [options]
```

### Options

| Options                   | Description                                            |
| ------------------------- | ------------------------------------------------------ |
| --host VALUE, -h VALUE    | broker or server host VALUE (required)                 |
| --service VALUE, -s VALUE | FIWARE Service VALUE                                   |
| --path VALUE, -p VALUE    | FIWARE ServicePath VALUE                               |
| --id VALUE, -i VALUE      | subscription id (required)                             |
| --data VALUE, -d VALUE    | subscription data                                      |
| --skipInitialNotification | skipInitialNotification (default: false)               |
| --subscriptionId VALUE    | subscription id (LD)                                   |
| --name VALUE              | subscription name (LD)                                 |
| --description VALUE       | description                                            |
| --entityId VALUE          | entity id                                              |
| --idPattern VALUE         | idPattern                                              |
| --type VALUE, -t VALUE    | entity type                                            |
| --typePattern VALUE       | typePattern (v2)                                       |
| --wAttrs VALUE            | watched attributes                                     |
| --timeInterval VALUE      | time interval (LD)                                     |
| --query VALUE, -q VALUE   | filtering by attribute value                           |
| --mq VALUE, -m VALUE      | filtering by metadata (v2)                             |
| --geometry VALUE          | geometry                                               |
| --coords VALUE            | coords                                                 |
| --georel VALUE            | georel                                                 |
| --geoproperty VALUE       | geoproperty (LD)                                       |
| --csf VALUE               | context source filter (LD)                             |
| --active                  | active (LD) (default: false)                           |
| --inactive                | inactive (LD) (default: false)                         |
| --nAttrs VALUE            | attributes to be notified                              |
| --keyValues, -K           | keyValues (default: false)                             |
| --uri VALUE, -u VALUE     | uri/url to be invoked when a notification is generated |
| --accept VALUE            | accept header (json or ld+json)                        |
| --expires VALUE, -e VALUE | expires                                                |
| --throttling VALUE        | throttling                                             |
| --timeRel VALUE           | temporal relationship (LD)                             |
| --timeAt VALUE            | timeAt (LD)                                            |
| --endTimeAt VALUE         | endTimeAt (LD)                                         |
| --timeProperty VALUE      | timeProperty (LD)                                      |
| --link VALUE, -L VALUE    | @context VALUE (LD)                                    |
| --context VLAUE, -C VLAUE | @context VLAUE (LD)                                    |
| --status VALUE            | status                                                 |
| --headers VALUE           | headers (v2)                                           |
| --qs VALUE                | qs (v2)                                                |
| --method VALUE            | method (v2)                                            |
| --payload VALUE           | payload (v2)                                           |
| --metadata VALUE          | metadata (v2)                                          |
| --exceptAttrs VALUE       | exceptAttrs (v2)                                       |
| --attrsFormat VALUE       | attrsFormat (v2)                                       |
| --safeString VALUE        | use safe string (VALUE: on/off)                        |
| --raw                     | handle raw data (default: false)                       |
| --help                    | show help (default: true)                              |

### Example

```console
ngsi update subscription --id 5fa78b70627088ba9b91b1c0 --expires 1day
```
