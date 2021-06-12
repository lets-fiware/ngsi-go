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

<a name="print-number-of-entities"></a>

## Print number of entities

This command prints number of entities.

```console
ngsi wc [common options] entities [options]
```

### Options

| Options                | Description                                         |
| ---------------------- | --------------------------------------------------- |
| --type value, -t value | specify Entity Type                                 |
| --link value, -L value | specify @context                                    |
| --skipForwarding       | skip forwarding to CPrs (v2) (Orion 3.1.0 or later) |
| --help                 | specify show help (default: false)                  |

### Examples

#### Example 1

```console
ngsi wc --host orion entities
```

```text
3606
```

#### Example 2

```console
ngsi wc --host orion entities --type EvacuationSpace
```

```text
231
```

<a name="print-number-of-subscriptions"></a>

## Print number of subscriptions

This command prints number of subscriptions.

```console
ngsi wc [common options] subscriptions
```

### Examples

#### Example 1

```console
ngsi wc --host orion subscriptions
```

```text
2
```

<a name="print-number-of-registrations"></a>

## Print number of registrations

This command prints number of registrations.

```console
ngsi wc [common options] registrations
```

### Examples

#### Example 1

```console
ngsi wc --host orion registrations
```

```text
1
```

<a name="print-number-of-types"></a>

## Print number of types

This command will print number of types.

```console
ngsi wc [common options] types
```

### Examples

#### Example 1

```console
ngsi wc --host orion types
```

```text
16
```
