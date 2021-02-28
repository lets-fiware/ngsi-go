# Global Options

| Options        | Description                                      |
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
   ngsi - command-line tool for FIWARE NGSI and NGSI-LD

USAGE:
   ngsi [global options] command [command options] [arguments...]

VERSION:
   0.7.0 (git_hash:a5885d223bdac4c5b3aba9c430eaa10c8584e161)

COMMANDS:
   help, h  Shows a list of commands or help for one command
   CONVENIENCE:
     admin     admin command for FIWARE Orion, Cygnus, Perseo, Scorpio
     apis      print endpoints of API
     cp        copy entities
     wc        print number of entities, subscriptions, registrations or types
     man       print urls of document
     health    print health status
     ls        list entities
     rm        remove entities
     receiver  notification receiver
     template  create template of subscription or registration
     version   print the version
   Context-Aware CEP:
     rules  rules command for PERSEO
   IoT Agent:
     devices   manage devices for IoT Agent
     services  manage services for IoT Agent
   Keyrock:
     applications   manage applications for Keyrock
     users          manage users for Keyrock
     organizations  manage organizations for Keyrock
     providers      print service providers for Keyrock
   MANAGEMENT:
     broker    manage config for broker
     context   manage @context
     settings  manage settings
     server    manage config for server
     token     manage token
   NGSI:
     append   append attributes
     create   create entity(ies), subscription or registration
     delete   delete entity(ies), attribute, subscription or registration
     get      get entity(ies), attribute(s), subscription, registration or type
     list     list types, entities, subscriptions or registrations
     replace  replace entities or attributes
     update   update entities, attribute(s) or subscription
     upsert   upsert entity or entities
   PERSISTING CONTEXT DATA:
     namemappings   manage namemappings for Cygnus
     groupingrules  manage groupingrules for Cygnus
   TIME SERIES:
     hdelete  delete historical raw and aggregated time series context information
     hget     get historical raw and aggregated time series context information

GLOBAL OPTIONS:
   --syslog LEVEL  specify logging LEVEL (off, err, info, debug)
   --stderr LEVEL  specify logging LEVEL (off, err, info, debug)
   --config FILE   specify configuration FILE
   --cache FILE    specify cache FILE
   --batch, -B     don't use previous args (batch) (default: false)
   --help          show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   (c) 2020-2021 Kazuhito Suda
```

## version

This option prints the version of NGSI Go.

```console
ngsi --version
```

```text
ngsi version 0.7.0 (git_hash:a5885d223bdac4c5b3aba9c430eaa10c8584e161)
```
