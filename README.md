[![Let's FIWARE Banner](https://raw.githubusercontent.com/lets-fiware/ngsi-go/gh-pages/img/lets-fiware-logo-non-free.png)](https://www.letsfiware.jp/)
[![NGSI v2](https://img.shields.io/badge/NGSI-v2-5dc0cf.svg)](https://fiware-ges.github.io/orion/api/v2/stable/)
[![NGSI LD](https://img.shields.io/badge/NGSI-LD-d6604d.svg)](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf)

![FIWARE: Tools](https://nexus.lab.fiware.org/repository/raw/public/badges/chapters/deployment-tools.svg)
[![License: MIT](https://img.shields.io/github/license/lets-fiware/ngsi-go.svg)](https://opensource.org/licenses/MIT)
![GitHub all releases](https://img.shields.io/github/downloads/lets-fiware/ngsi-go/total)
[![Support badge](https://img.shields.io/badge/tag-fiware-orange.svg?logo=stackoverflow)](https://stackoverflow.com/questions/tagged/fiware+ngsi-go)
<br/>
![GitHub top language](https://img.shields.io/github/languages/top/lets-fiware/ngsi-go)
![Lines of code](https://img.shields.io/tokei/lines/github/lets-fiware/ngsi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/lets-fiware/ngsi-go)](https://goreportcard.com/report/github.com/lets-fiware/ngsi-go)
![CI](https://github.com/lets-fiware/ngsi-go/workflows/CI/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)

The NGSI Go is a command-line interface supporting FIWARE Open APIs for FIWARE developers.

| :books: [Documentation](https://ngsi-go.letsfiware.jp/) | :dart: [Roadmap](./roadmap.md) |
|---------------------------------------------------------|--------------------------------|

## Contents

<details>
<summary><strong>Details</strong></summary>

-   [Getting Started with NGSI Go](#getting-started-with-ngsi-go)
-   [Usage](#usage)
-   [Tutorial](#tutorial)
-   [Install](#install)
-   [Third party packages](#third-party-packages)
-   [Copyright and License](#copyright-and-license)

</details>

# What is NGSI Go?

> "Brave (hero), bearer of the blood of Erdrick, hero of legend! Know that your weapon will not
> serve to vanquish the Dragonload."
>
> — DRAGON WARRIOR (DRAGON QUEST)

The NGSI Go is a command-line interface supporting FIWARE Open APIs, which simplifies syntax.
It's a powerful tool and easy to use. It has various features as shown:

-   Supported FIWARE Open APIs
    -   FIWARE [NGSI v2](https://fiware-ges.github.io/orion/api/v2/stable/) API
    -   [NGSI-LD](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf) API
    -   [STH-Comet](https://github.com/telefonicaid/fiware-sth-comet) API
    -   [QuantumLeap](https://github.com/orchestracities/ngsi-timeseries-api) API
    -   [Cygnus](https://github.com/telefonicaid/fiware-cygnus/blob/master/doc/cygnus-common/installation_and_administration_guide/management_interface_v1.md) API
    -   [IoT Agent](https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/apiary/iotagent.apib) Provision API
    -   [Perseo](https://github.com/telefonicaid/perseo-fe/blob/master/documentation/api.md) API
    -   [Keyrock](https://github.com/ging/fiware-idm/blob/master/apiary.apib) API
-   Various convenience commands
    -   NGSI commands to manage NGSI Entity, subscription, registration and so on
    -   Time series commands to manage historical data
    -   IoT Agent commands to manage IoT Agent Provision APIs
    -   Rules commands to manage Context-Aware CEP
    -   Convenience commands
        -   Print version, health-check status and API lists of FIWARE GEs
        -   Admin command for FIWARE Orion
        -   Copy and remove entities at once
        -   Create template of subscription or registration
        -   Notification receiver
    -   Management commands
        -   Broker alias or server alias with API endpoint URL, FIWARE Service and FIWARE ServicePath
        -   Manage @context
        -   Integrated oauth token management
-   Compatible with a number of traditional UNIX commands for filtering text
-   A single binary program written in Golang

## Getting Started with NGSI Go

You register an alias to access the broker.

```console
ngsi broker add --host letsfiware --brokerHost http://localhost:1026 --ngsiType v2
```

You can get the version by using the alias `letsfiware`.

```console
ngsi version -h letsfiware
```

```json
{
 "orion" : {
   "version" : "2.5.2",
   "uptime" : "0 d, 13 h, 54 m, 48 s",
   "git_hash" : "11e4cbfef30d28347162e5c4ef4de3a5d2797c69",
   "compile_time" : "Thu Dec 17 08:43:46 UTC 2020",
   "compiled_by" : "root",
   "compiled_in" : "5a4a8800b1fa",
   "release_date" : "Thu Dec 17 08:43:46 UTC 2020",
   "doc" : "https://fiware-orion.rtfd.io/en/2.5.2/",
   "libversions": {
      "boost": "1_53",
      "libcurl": "libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
      "libmicrohttpd": "0.9.70",
      "openssl": "1.0.2k",
      "rapidjson": "1.1.0",
      "mongodriver": "legacy-1.1.2"
   }
 }
}
```

Once you access the broker, you can omit to specify the broker.

```console
ngsi version
```

If you want to check the current settings, you can run the following command.

```console
ngsi settings list
```

## Usage

```text
NAME:
   ngsi - command-line tool for FIWARE NGSI and NGSI-LD

USAGE:
   ngsi [global options] command [command options] [arguments...]

VERSION:
   0.8.0 (git_hash:995c1cec43797999a40ed6b686985dbdd0b2afcc)

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

## Tutorial

You can try [the tutorial](docs/tutorial/index.md) to understand how to use the NGSI Go.
You need a environment running Docker engine and docker-compose.

## Install

### Install NGSI Go binary

The NGSI Go binary is installed in `/usr/local/bin`.

#### Installation on Linux

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.8.0/ngsi-v0.8.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.8.0-linux-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.8.0-linux-arm.tar.gz` and `ngsi-v0.8.0-linux-arm64.tar.gz` binaries are also available.

#### Installation on Mac

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.8.0/ngsi-v0.8.0-darwin-amd64.tar.gz
sudo tar zxvf ngsi-v0.8.0-darwin-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.8.0-darwin-arm64.tar.gz` binary is also available.

### Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```console
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

## Third party packages

The NGSI Go makes use of the following package:

| Package                                         | OSS License        |
| ----------------------------------------------- | ------------------ |
| [urfave/cli](https://github.com/urfave/cli)     | MIT License        |

The dependencies of dependencies have been omitted from the list.

## Copyright and License

Copyright (c) 2020-2021 Kazuhito Suda<br>
Licensed under the [MIT License](./LICENSE).
