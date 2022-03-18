[![Let's FIWARE Banner](https://raw.githubusercontent.com/lets-fiware/ngsi-go/gh-pages/img/ngsi-go-logo-non-free.png)](https://www.letsfiware.jp/)
[![NGSI v2](https://img.shields.io/badge/NGSI-v2-5dc0cf.svg)](https://fiware-ges.github.io/orion/api/v2/stable/)
[![NGSI LD](https://img.shields.io/badge/NGSI-LD-d6604d.svg)](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.05.01_60/gs_CIM009v010501p.pdf)

![FIWARE: Tools](https://nexus.lab.fiware.org/repository/raw/public/badges/chapters/deployment-tools.svg)
[![License: MIT](https://img.shields.io/github/license/lets-fiware/ngsi-go.svg)](https://opensource.org/licenses/MIT)
![GitHub all releases](https://img.shields.io/github/downloads/lets-fiware/ngsi-go/total)
[![Support badge](https://img.shields.io/badge/tag-fiware-orange.svg?logo=stackoverflow)](https://stackoverflow.com/questions/tagged/fiware+ngsi-go)
<br/>
![GitHub top language](https://img.shields.io/github/languages/top/lets-fiware/ngsi-go)
![Lines of code](https://img.shields.io/tokei/lines/github/lets-fiware/ngsi-go)
[![Lint](https://github.com/lets-fiware/ngsi-go/actions/workflows/lint.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lets-fiware/ngsi-go)](https://goreportcard.com/report/github.com/lets-fiware/ngsi-go)
<br/>
[![Build](https://github.com/lets-fiware/ngsi-go/actions/workflows/build.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/build.yml)
[![Unit Tests](https://github.com/lets-fiware/ngsi-go/actions/workflows/unit-test.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/unit-test.yml)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)
[![E2E tests](https://github.com/lets-fiware/ngsi-go/actions/workflows/e2e-test.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/e2e-test.yml)
<br/>
[![Docs](https://github.com/lets-fiware/ngsi-go/actions/workflows/docs.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/docs.yml)
[![Dockerfile](https://github.com/lets-fiware/ngsi-go/actions/workflows/dockerfile.yml/badge.svg)](https://github.com/lets-fiware/ngsi-go/actions/workflows/dockerfile.yml)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/4973/badge)](https://bestpractices.coreinfrastructure.org/projects/4973)

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
    -   [NGSI-LD](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.05.01_60/gs_CIM009v010501p.pdf) API
    -   [STH-Comet](https://github.com/telefonicaid/fiware-sth-comet) API
    -   [QuantumLeap](https://github.com/orchestracities/ngsi-timeseries-api) API
    -   [Cygnus](https://github.com/telefonicaid/fiware-cygnus/blob/master/doc/cygnus-common/installation_and_administration_guide/management_interface_v1.md) API
    -   [IoT Agent](https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/apiary/iotagent.apib) Provision API
    -   [Perseo](https://github.com/telefonicaid/perseo-fe/blob/master/documentation/api.md) API
    -   [Keyrock](https://github.com/ging/fiware-idm/blob/master/apiary.apib) API
    -   [WireCloud](https://github.com/Wirecloud/wirecloud/blob/develop/docs/restapi/applicationmashup.apib) API
-   Various convenience commands
    -   NGSI commands to manage NGSI Entity, subscription, registration and so on
    -   Time series commands to manage historical data
    -   IoT Agent commands to manage IoT Agent Provision API
    -   Rules commands to manage Context-Aware CEP
    -   WireCloud command to manage Application Mashup RESTful API
    -   Convenience commands
        -   Print version, health-check status and API lists of FIWARE GEs
        -   Admin command for FIWARE Orion
        -   Copy and remove entities at once
        -   Create template of subscription or registration
        -   Notification receiver
        -   Registration proxy
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
  "version" : "3.6.0",
  "uptime" : "0 d, 0 h, 0 m, 1 s",
  "git_hash" : "973850279e63d58cb93dff751648af5ec6e05777",
  "compile_time" : "Wed Mar 2 10:34:48 UTC 2022",
  "compiled_by" : "root",
  "compiled_in" : "5e6b6f1167f7",
  "release_date" : "Wed Mar 2 10:34:48 UTC 2022",
  "machine" : "x86_64",
  "doc" : "https://fiware-orion.rtfd.io/en/3.6.0/",
  "libversions": {
     "boost": "1_66",
     "libcurl": "libcurl/7.61.1 OpenSSL/1.1.1k zlib/1.2.11 nghttp2/1.33.0",
     "libmosquitto": "2.0.12",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.1",
     "rapidjson": "1.1.0",
     "mongoc": "1.17.4",
     "bson": "1.17.4"
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
   ngsi - command-line tool for FIWARE Open APIs

USAGE:
   ngsi [global options] command [options] [arguments...]

VERSION:
   ngsi version 0.12.0 (git_hash:06a13ec2347c05c9fae96106577c06371b7c6bf5)

COMMANDS:
   help, h  Shows a list of commands or help for one command
   APPLICATION MASHUP:
     preferences  manage preferences for WireCloud
     macs         manage mashable application components for WireCloud
     workspaces   manage workspaces for WireCloud
     tabs         manage tabs for WireCloud
   Context-Aware CEP:
     rules  rules command for PERSEO
   CONVENIENCE:
     admin       admin command for FIWARE Orion, Cygnus, Perseo, Scorpio
     apis        print endpoints of API
     cp          copy entities
     wc          print number of entities, subscriptions, registrations or types
     man         print urls of document
     health      print health status
     ls          list entities
     queryproxy  query proxy
     rm          remove entities
     receiver    notification receiver
     regproxy    registration proxy
     template    create template of subscription or registration
     tokenproxy  token proxy
     version     print the version
   IoT Agent:
     devices   manage devices for IoT Agent
     services  manage services for IoT Agent
   Keyrock:
     applications   manage applications for Keyrock
     users          manage users for Keyrock
     organizations  manage organizations for Keyrock
     providers      print service providers for Keyrock
   NGSI:
     append   append attributes
     create   create entity(ies), subscription, registration or ldContext
     delete   delete entity(ies), attribute, subscription, registration or ldContext
     get      get entity(ies), attribute(s), subscription, registration type or ldContext
     list     list types, attributes, entities, tentities, subscriptions or registrations
     replace  replace entities or attributes
     update   update entities, attribute(s) or subscription
     upsert   upsert entity or entities
   PERSISTING CONTEXT DATA:
     namemappings   manage namemappings for Cygnus
     groupingrules  manage groupingrules for Cygnus
   TIME SERIES:
     hdelete  delete historical raw and aggregated time series context information
     hget     get historical raw and aggregated time series context information
   MANAGEMENT:
     broker    manage config for broker
     context   manage @context
     settings  manage settings
     server    manage config for server
     token     manage token
     license   print OSS license information

GLOBAL OPTIONS:
   --syslog LEVEL        syslog logging LEVEL (off, err, info, debug)
   --stderr LEVEL        stderr logging LEVEL (err, info, debug)
   --configDir DIR       configuration DIR name
   --config FILE         configuration FILE name
   --cache FILE          cache FILE name
   --batch, -B           don't use previous args (batch) (default: false)
   --insecureSkipVerify  TLS/SSL skip certificate verification (default: false)
   --help                show help (default: false)
   --version, -v         print the version (default: false)

PREVIOUS ARGS:
   None
```

## Tutorial

You can try [the tutorial](docs/tutorial/index.md) to understand how to use the NGSI Go.
You need a environment running Docker engine and docker-compose.

## Install

### Install NGSI Go binary

The NGSI Go binary is installed in `/usr/local/bin`.

#### Installation on Linux

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.12.0/ngsi-v0.12.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.12.0-linux-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.12.0-linux-arm.tar.gz` and `ngsi-v0.12.0-linux-arm64.tar.gz` binaries are also available.

#### Installation on Mac

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.12.0/ngsi-v0.12.0-darwin-amd64.tar.gz
sudo tar zxvf ngsi-v0.12.0-darwin-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.12.0-darwin-arm64.tar.gz` binary is also available.

### Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```console
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

## Third party packages

The NGSI Go makes no use of third-party packages.

-   [Open Source Insights](https://deps.dev/go/github.com%2Flets-fiware%2Fngsi-go)

## Copyright and License

Copyright (c) 2020-2022 Kazuhito Suda<br>
Licensed under the [MIT License](./LICENSE).
