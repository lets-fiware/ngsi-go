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
[![Build Status](https://travis-ci.com/lets-fiware/ngsi-go.svg?branch=main)](https://travis-ci.com/lets-fiware/ngsi-go)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)

# What is NGSI Go?

The NGSI Go is a UNIX command-line tool FIWARE supporting both [NGSI v2](https://fiware-ges.github.io/orion/api/v2/stable/) and [NGSI-LD](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf), which simplifies syntax.

## Contents

-   [Usage](usage.md)
-   [Quick Start Guide](quick_start_guide.md)
-   [Tutorial](tutorial/index.md)
-   [Install](install.md)
-   [Build from source](build_source.md)
-   [FIWARE Open APIs mapping table](apis_mapping_table.md)

## Tutorial

-   [NGSI-LD CRUD](tutorial/ngsi-ld-crud.md)
-   [NGSIv2 CRUD](tutorial/ngsi-v2-crud.md)
-   [STH-Comet](tutorial/comet.md)
-   [QuantumLeap](tutorial/quantumleap.md)
-   [IoT Agent](tutorial/iot-agent.md)

## Command reference

### NGSI

-   [append](ngsi/append.md): append attributes
-   [create](ngsi/create.md): create entity(ies), subscription or registration
-   [delete](ngsi/delete.md): delete entity(ies), attribute, subscription or registration
-   [get](ngsi/get.md): get entity(ies) or attribute(s)
-   [list](ngsi/list.md): list types, entities, subscriptions or registrations
-   [replace](ngsi/replace.md): replace entities or attributes
-   [update](ngsi/update.md): update entities, attribute(s) or subscription
-   [upsert](ngsi/upsert.md): upsert entities

### Time series

-   [hdelete](time_series/hdelete.md): delete historical data
-   [hget](time_series/hget.md): get historical data

### Convenience

-   [admin](convenience/admin.md): administrative command for FIWARE Orion
-   [apis](convenience/apis.md): print endpoints of FWARE Open APIs
-   [cp](convenience/cp.md): copy entities
-   [wc](convenience/wc.md): print number of entities, subscriptions or registrations
-   [man](convenience/man.md): print  URLs of the documents related to the NGSI Go
-   [health](convenience/health.md): print health status of FIWARE GEs
-   [ls](convenience/ls.md): list entities
-   [rm](convenience/rm.md): remove entities
-   [receiver](convenience/receiver.md): notification receiver
-   [template](convenience/template.md): create template of subscription or registration
-   [version](convenience/version.md): print the version of Context Broker

### Management

-   [broker](management/broker.md): manage config for broker
-   [context](management/context.md): manage @context
-   [settings](management/settings.md):  manage settings
-   [server](management/server.md): manage config for server
-   [token](management/token.md): manage token

### Global Options

-   [Global Options](global.md)

### Files

-   [Files](files.md)

## GitHub repository

-   [lets-fiware/ngsi-go](https://github.com/lets-fiware/ngsi-go/)

## Copyright and License

Copyright (c) 2020-2021 Kazuhito Suda<br>
Licensed under the [MIT license](https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/LICENSE).
