# Global Options

| Options	 | Description                                      |
| -------------- | ------------------------------------------------ |
| --syslog LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --stderr LEVEL | specify logging LEVEL (off, err, info, debug)    |
| --config FILE  | specify configuration FILE                       |
| --cache FILE   | specify cache FILE                               |
| --batch, -B    | don't use previous args (batch) (default: false) |
| --help         | show help (default: false)                       |
| --version, -v  | print the version (default: false)               |

## syslog

This option specifies the logging LEVEL of messages to be output to syslog.

## stderr

This option specifies the logging LEVEL of messages to be output to stderr.

## cache

This option specifies a cache file.

## batch

This option doesn't use previous args.

## help

This option prints the usage of NGSI Go.

```console
ngsi --help
```

```text
NAME:
   ngsi - unix-like command-line tool for FIWARE NGSI and NGSI-LD

USAGE:
   ngsi [global options] command [command options] [arguments...]

VERSION:
   0.3.0 (git_hash:d466356676201c8b237eaa4e768b9742d1b91120)

COMMANDS:
   help, h  Shows a list of commands or help for one command
   CONVENIENCE:
     admin     admin command for FIWARE Orion
     cp        copy entities
     wc        print number of entities, subscriptions, registrations, or types
     man       print urls of document
     ls        list entities
     rm        remove entities
     template  create template of subscription or registration
     version   print the version of Context Broker
   MANAGEMENT:
     broker    manage config for broker
     context   manage @context
     settings  manage settings
     token     manage token
   NGSI:
     append   append attributes
     create   create entity(ies), subscription or registration
     delete   delete entity(ies), attribute, subscription, or registration
     get      get entity(ies), attribute(s), subscription, registration, or type
     list     list types, entities, subscriptions, or registrations
     replace  replace entities or attributes
     update   update entities, attribute(s), or subscription
     upsert   upsert entities

GLOBAL OPTIONS:
   --syslog LEVEL  specify logging LEVEL (off, err, info, debug)
   --stderr LEVEL  specify logging LEVEL (off, err, info, debug)
   --config FILE   specify configuration FILE
   --cache FILE    specify cache FILE
   --batch, -B     don't use previous args (batch) (default: false)
   --help          show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   (c) 2020 Kazuhito Suda
```

## version

This option prints the version of NGSI Go.

```console
ngsi --version
```

```text
ngsi version 0.3.0 (git_hash:d466356676201c8b237eaa4e768b9742d1b91120)
```
