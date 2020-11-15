# wc - Convenience command

This command prints number of entities, subscriptions, registrations, or types.

## Commands

-   [Print number of entities](#print-number-of-entities)
-   [Print number of subscription](#print-number-of-subscriptions)
-   [Print number of registrations](#print-number-of-registrations)
-   [Print number of types](#print-number-of-types)

### Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --service value, -s value | specify FIWARE Service     |
| --path value, -p value    | specify FIWARE ServicePath |
| --help                    | show help (default: false) |

<a name="print-number-of-entities"/>

## Print number of entities

This command prints number of entities.

```bash
ngsi wc [common options] entities [options]
```

### Options

| Options                | Description                        |
| ---------------------- | ---------------------------------- |
| --type value, -t value | specify Entity Type                |
| --link value, -L value | specify @context                   |
| --help                 | specify show help (default: false) |

### Exapmles

#### Example 1

```bash
$ ngsi wc --host orion entities
3606
```

#### Example 2

```
$ ngsi wc --host orion entities --type EvacuationSpace
231
```

<a name="print-number-of-subscriptions"/>

## Print number of subscriptions

This command prints number of subscriptions.

```bash
ngsi wc [common options] subscriptions
```

### Exapmles

#### Example 1

```bash
$ ngsi wc --host orion subscriptions
2
```

<a name="print-number-of-registrations"/>

## Print number of registrations

This command prints number of registrations.

```bash
ngsi wc [common options] registrations
```

### Exapmles

#### Example 1

```bash
$ ngsi wc --host orion registrations
1
```

<a name="print-number-of-types"/>

## Print number of types

This command will print number of types.

```bash
ngsi wc [common options] types
```

### Exapmles

#### Example 1

```bash
$ ngsi wc --host orion types
16
```
