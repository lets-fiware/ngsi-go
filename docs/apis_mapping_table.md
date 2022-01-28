# FIWARE Open APIs mapping table

These tables show the mapping from FIWARE Open APIs to NGSI Go commands.

<details>
<summary><strong>Details</strong></summary>

-   [NGSI-LD API](#ngsi-ld-api-etsi-gs-cim-009-v141-2021-02)
    -   [Orion-LD API](#orion-ld-api)
    -   [Scorpio broker API](#scorpio-broker-api)
-   [FIWARE NGSI v2](#fiware-ngsi-v2)
    -   [Orion API](#orion-api)
-   [STH-Comet API](#sth-comet-api)
-   [QuantumLeap API](#quantumleap-api)
-   [Cygnus API](#cygnus-api)
-   [IoT Agent Provision API](#iot-agent-provision-api)
-   [Perseo FE](#perseo-fe)
-   [Perseo CORE](#perseo-core)
-   [Keyrock API](#keyrock-api)

</details>

## NGSI-LD API (ETSI GS CIM 009 V1.4.1 2021-02)

| NGSI-LD API                                                                 | NGSI Go commands                                                           | 
| --------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| POST /ngsi-ld/v1/entities/                                                  | create entity                                                              |
| GET /ngsi-ld/v1/entities/                                                   | list entities                                                              |
| GET /ngsi-ld/v1/entities/{entityId}                                         | get entity --id {entityId}                                                 |
| DELETE /ngsi-ld/v1/entities/{entityId}                                      | delete entity --id {entityId}                                              |
| POST /ngsi-ld/v1/entities/{entityId}/attrs/                                 | append attrs --id {entityId}                                               |
| PATCH /ngsi-ld/v1/entities/{entityId}/attrs/                                | update attrs --id {entityId}                                               |
| PATCH /ngsi-ld/v1/entities/{entityId}/attrs/{attrId}                        | update attr --id {entityId} --attr {attrId}                                |
| DELETE /ngsi-ld/v1/entities/{entityId}/attrs/{attrId}                       | delete attr --id {entityId} --attr {attrId}                                |
| POST /ngsi-ld/v1/subscriptions/                                             | create subscription                                                        |
| GET /ngsi-ld/v1/subscriptions/                                              | list subscriptions                                                         |
| GET /ngsi-ld/v1/subscriptions/{subscriptionId}                              | get subscription --id {subscriptionId}                                     |
| PATCH /ngsi-ld/v1/subscriptions/{subscriptionId}                            | update subscription --id {subscriptionId}                                  |
| DELETE /ngsi-ld/v1/subscriptions/{subscriptionId}                           | delete subscription --id {subscriptionId}                                  |
| GET /ngsi-ld/v1/types/                                                      | list types                                                                 |
| GET /ngsi-ld/v1/types/{type}                                                | (not yet implemented)                                                      |
| GET /ngsi-ld/v1/attributes/                                                 | list attributes                                                            |
| GET /ngsi-ld/v1/attributes/{attrId}                                         | list attributes --attr {attrId}                                            |
| POST /ngsi-ld/v1/csourceRegistrations/                                      | create registration                                                        |
| GET /ngsi-ld/v1/csourceRegistrations/                                       | list registrations                                                         |
| GET /ngsi-ld/v1/csourceRegistrations/{registrationId}                       | get registration --id {registrationId}                                     |
| PATCH /ngsi-ld/v1/csourceRegistrations/{registrationId}                     | update registration --id {registrationId}                                  |
| DELETE /ngsi-ld/v1/csourceRegistrations/{registrationId}                    | delete registration --id {registrationId}                                  |
| POST /ngsi-ld/v1/csourceSubscriptions/                                      | (not yet implemented)                                                      |
| GET /ngsi-ld/v1/csourceSubscriptions/                                       | (not yet implemented)                                                      |
| GET /ngsi-ld/v1/csourceSubscriptions/{subscriptionId}                       | (not yet implemented)                                                      |
| PATCH /ngsi-ld/v1/csourceSubscriptions/{subscriptionId}                     | (not yet implemented)                                                      |
| DELETE /ngsi-ld/v1/csourceSubscriptions/{subscriptionId}                    | (not yet implemented)                                                      |
| POST /ngsi-ld/v1/entityOperations/create                                    | create entities                                                            |
| POST /ngsi-ld/v1/entityOperations/upsert                                    | upsert entities                                                            |
| POST /ngsi-ld/v1/entityOperations/update                                    | update entities                                                            |
| POST /ngsi-ld/v1/entityOperations/delete                                    | delete entities                                                            |
| POST /ngsi-ld/v1/entityOperations/query                                     | (not yet implemented)                                                      |
| POST /ngsi-ld/v1/temporal/entities/                                         | create tentity                                                             |
| GET /ngsi-ld/v1/temporal/entities/                                          | list tentities                                                             |
| GET /ngsi-ld/v1/temporal/entities/{entityId}                                | get tentity --id {entityId}                                                |
| DELETE /ngsi-ld/v1/temporal/entities/{entityId}                             | delete tentity --id {entityId}                                             |
| POST /ngsi-ld/v1/temporal/entities/{entityId}/attrs/                        | append tattrs --id {entityId}                                              |
| DELETE /ngsi-ld/v1/temporal/entities/{entityId}/attrs/{attrId}              | delete tattr --id {entityId} --attr {attrId}                               |
| PATCH /ngsi-ld/v1/temporal/entities/{entityId}/attrs/{attrId}/{instanceId}  | update tattr --id {entityId} --attr {attrId} --instanceId {instanceId}     |
| DELETE /ngsi-ld/v1/temporal/entities/{entityId}/attrs/{attrId}/{instanceId} | delete tattr --id {entityId} --attr {attrId} --instanceId {instanceId}     |
| POST /ngsi-ld/v1/temporal/entityOperations/query                            | (not yet implemented)                                                      |
| GET /ngsi-ld/v1/jsonldContexts                                              | list ldContexts                                                            |
| GET /ngsi-ld/v1/jsonldContexts/{contextId}                                  | get ldContext --id {contextId}                                             |
| POST /ngsi-ld/v1/jsonldContexts                                             | create ldContext --data {jsonldContext}                                    |
| DELETE /ngsi-ld/v1/jsonldContexts/{contextId}                               | delete ldContext --id {contextId}                                          |

### Orion-LD API

| Orion-LD API               | NGSI Go commands      |
| -------------------------- | --------------------- |
| GET /ngsi-ld/ex/v1/version | version               |
| POST /ngsi-ld/ex/v1/notify | (not yet implemented) |

### Scorpio broker API

| Scorpio API                     | NGSI Go commands         |
| ------------------------------- | ------------------------ |
| GET /scorpio/v1/info/           | admin scorpio list       |
| GET /scorpio/v1/info/types      | admin scorpio types      |
| GET /scorpio/v1/info/localtypes | admin scorpio localtypes |
| GET /scorpio/v1/info/stats      | admin scorpio stats      |
| GET /scorpio/v1/info/health     | admin scorpio health     |

## FIWARE NGSI v2

| FIWARE NGSI v2                                     | NGSI Go commands                                  | 
| -------------------------------------------------- | ------------------------------------------------- |
| GET /v2                                            | apis                                              |
| GET /v2/entities                                   | list entities                                     |
| POST /v2/entities                                  | create entity                                     |
| GET /v2/entities/{entityId}                        | get entity --id {entityId}                        |
| DELETE /v2/entities/{entityId}                     | delete entity --id {entityId}                     |
| GET /v2/entities/{entityId}/attrs                  | get attrs --id {entityId}                         |
| POST /v2/entities/{entityId}/attrs                 | append attributes --id {entityId}                 |
| PATCH /v2/entities/{entityId}/attrs                | update attributes --id {entityId}                 |
| PUT /v2/entities/{entityId}/attrs                  | replace attributes --id {entityId}                |
| GET /v2/entities/{entityId}/attrs/{attrName}       | get attr --id {entityId} --attr {attrName}        |
| PUT /v2/entities/{entityId}/attrs/{attrName}       | update attr --id {entityId} --attr {attrName}     |
| DELETE /v2/entities/{entityId}/attrs/{attrName}    | delete attr --id {entityId} --attr {attrName}     |
| GET /v2/entities/{entityId}/attrs/{attrName}/value | (not yet implemented)                             |
| PUT /v2/entities/{entityId}/attrs/{attrName}/value | (not yet implemented)                             |
| GET /v2/types/                                     | list types                                        |
| GET /v2/types/{entityType}                         | get type --type {entityType}                      |
| GET /v2/subscriptions                              | list subscriptions                                |
| POST /v2/subscriptions                             | create subscription                               |
| GET /v2/subscriptions/{subscriptionId}             | get subscription --id {subscriptionId}            |
| PATCH /v2/subscriptions/{subscriptionId}           | update subscription --id {subscriptionId}         |
| DELETE /v2/subscriptions/{subscriptionId}          | delete subscription --id {subscriptionId}         |
| GET /v2/registrations                              | list registrations                                |
| POST /v2/registrations                             | create registration                               |
| GET /v2/registrations/{registrationId}             | get registration --id {registrationId}            |
| PATCH /v2/registrations/{registrationId}           | (not yet implemented)                             |
| DELETE /v2/registrations/{registrationId}          | delete registration --id {registrationId}         |
| POST /v2/op/update actionType=append               | upsert entities                                   |
| POST /v2/op/update actionType=appendStrict         | create entities                                   |
| POST /v2/op/update actionType=update               | update entities                                   |
| POST /v2/op/update actionType=delete               | delete entities                                   |
| POST /v2/op/update actionType=replace              | replace entities                                  |
| POST /v2/op/query                                  | get entities                                      |
| POST /v2/op/notify                                 | (not yet implemented)                             |

### Orion API

| Orion API                      | NGSI Go commands                          |
| ------------------------------ | ----------------------------------------- |
| GET /version                   | version                                   |
| GET /admin/log                 | admin log                                 |
| PUT /admin/log                 | admin log --level {logLevel}              |
| GET /log/trace                 | admin trace                               |
| PUT /log/trace/{traceLevel}    | admin trace --level {traceLevel}          |
| DELETE /log/trace              | admin trace --delete                      |
| DELETE /log/trace/{traceLevel} | admin trace --delete --level {traceLevel} |
| GET /admin/sem                 | admin semaphore                           |
| GET /admin/metrics             | admin metrics                             |
| DELETE /admin/metrics          | admin metrics --reset                     |
| GET /admin statistics          | admin statistics                          |
| DELETE /admin statistics       | admin statistics --delete                 |
| GET /cache/statistics          | admin cacheStatistics                     |
| DELETE /cache/statistics       | admin cacheStatistics --delete            |

## STH-Comet API 

| STH-Comet API                                                                                              | NGSI Go commands                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| GET /version                                                                                               | version                                                                                                         |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&hLimit={n}&hOffset={n}                  | hget attr --hLimit {n} --type {entityType} --id {enttiyId} --attr {attrName}                                    |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?entityType={entityType}&lastN={n}                         | hget attr --lastN {n} --type {entityType} --id {enttiyId} --attr {attrName}                                     |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&aggrMethod={method}&aggrPeriod={period} | hget attr --arrgMethod {method} --aggrPeriod {period} --type {entityType} --id {enttiyId} --attr {attrName}     |
| DELETE /STH/v1/contextEntities                                                                             | hdelete                                                                                                         |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}                                             | hdelete --type {entityType} --id {enttiyId}                                                                     |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}/attributes/{attrName}                       | hdelete --type {entityType} --id {enttiyId} --attr {attrName}                                                   |

-   [STH-Comet API - GitHub](https://github.com/telefonicaid/fiware-sth-comet/blob/master/apiary.apib)

## QuantumLeap API 

| QuantumLeap API                                    | NGSI Go commands                                                       |
| -------------------------------------------------- | ---------------------------------------------------------------------- |
| GET /v2/                                           | apis                                                                   |
| GET /version                                       | version                                                                |
| POST /config                                       | (not yet implemented)                                                  |
| GET /health                                        | health                                                                 |
| POST /notify                                       | (not yet implemented)                                                  |
| POST /subscribe                                    | (not yet implemented)                                                  |
| GET /v2/entities                                   | hget entities                                                          |
| GET /v2/entities/{entityId}/attrs/{attrName}       | hget attr --id {entityId} --attr {attrName}                            |
| GET /v2/entities/{entityId}/attrs/{attrName}/value | hget attr --id {entityId} --attr {attrName} --value                    |
| GET /v2/entities/{entityId}                        | hget attrs --id {entityId}                                             |
| GET /v2/entities/{entityId}/value                  | hget attrs --id {entityId} --value                                     |
| GET /v2/types/{entityType}/attrs/{attrName}        | hget attr --sameType --type {entityType} --attr {attrName}             |
| GET /v2/types/{entityType}/attrs/{attrName}/value  | hget attr --sameType --type {entityType} --attr {attrName} --value     |
| GET /v2/types/{entityType}                         | hget attrs --sameType --type {entityType}                              |
| GET /v2/types/{entityType}/value                   | hget attrs --sameType --type {entityType} --value                      |
| GET /v2/attrs/{attrName}                           | hget attr --nTypes --attr {attrName}                                   |
| GET /v2/attrs/{attrName}/value                     | hget attr --nTypes --attr {attrName} --value                           |
| GET /v2/attrs                                      | hget attrs --nTypes                                                    |
| GET /v2/attrs/value                                | hget attrs --nTypes --value                                            |
| DELETE /v2/entities/{entityId}                     | hdelete entity --id {entityId}                                         |
| DELETE /v2/types/{entityType}                      | hdelete entities --type {entityType}                                   |

-   [QuantumLeap API - GitHub](https://github.com/smartsdk/ngsi-timeseries-api/blob/master/specification/quantumleap.yml)

## Cygnus API

| Cygnus API                                         | NGSI Go commands                                                       |
| -------------------------------------------------- | ---------------------------------------------------------------------- |
| GET /v1/version                                    | version                                                                |
| GET /v1/stats                                      | admin statistics                                                       |
| PUT /v1/stats                                      | admin statistics --delete                                              |
| GET /v1/namemappings                               | namemappings list                                                      |
| POST /v1/namemappings                              | namemappings create                                                    |
| PUT /v1/namemappings                               | namemappings update                                                    |
| DELETE /v1/namemappings                            | namemappings delete                                                    |
| GET /v1/groupingrules                              | groupingrules list                                                     |
| GET /v1/groupingrules                              | groupingrules get                                                      |
| POST /v1/groupingrules                             | groupingrules create                                                   |
| PUT /v1/groupingrules                              | groupingrules update                                                   |
| DELETE /v1/groupingrules                           | groupingrules delete                                                   |
| POST /notify                                       | (not yet implemented)                                                  |
| GET /v1/subscriptions                              | (not yet implemented)                                                  |
| POST /v1/subscriptions                             | (not yet implemented)                                                  |
| DELETE /v1/subscriptions                           | (not yet implemented)                                                  |
| GET /admin/log                                     | admin log                                                              |
| PUT /admin/log                                     | admin log --level {log_level}                                          |
| GET /v1/admin/metrics                              | admin metrics                                                          |
| DELETE /v1/admin/metrics                           | admin metrics --delete                                                 |
| GET /v1/admin/log/loggers                          | admin loggers list                                                     |
| GET /v1/admin/log/loggers?name={name}              | admin loggers get --name {name}                                        |
| POST /v1/admin/log/loggers                         | admin loggers                                                          |
| PUT /v1/admin/log/loggers                          | admin loggers                                                          |
| DELETE /v1/admin/log/loggers                       | admin loggers delete                                                   |
| DELETE /v1/admin/log/loggers?name={name}           | admin loggers delete --name {name}                                     | 
| GET /v1/admin/log/appenders                        | admin appenders list                                                   |
| GET /v1/admin/log/appenders?name={name}            | admin appenders get --name {name}                                      |
| POST /v1/admin/log/appenders                       | admin appenders                                                        |
| PUT /v1/admin/log/appenders                        | admin appenders                                                        |
| DELETE /v1/admin/log/appenders                     | admin appenders delete                                                 |
| DELETE /v1/admin/log/appenders?name={name}         | admin appenders delete --name {name}                                   |

-   [Cygnus API - GitHub](https://github.com/telefonicaid/fiware-cygnus/blob/master/doc/cygnus-common/installation_and_administration_guide/management_interface_v1.md)

## IoT Agent Provision API

| IoT Agent Provision API     | NGSI Go commands                |
| --------------------------- | ------------------------------- |
| GET /services               | services list                   |
| POST /services              | services create                 |
| PUT /services               | services update                 |
| DELETE /services            | services delete                 |
| GET /devices                | devices list                    |
| GET /devices/{device_id}    | devices get --id {device_id}    |
| POST /devices/{device_id}   | devices create --id {device_id} |
| PUT /devices/{device_id}    | devices update --id {device_id} |
| DELETE /devices/{device_id} | devices delete --id {device_id} |

-   [IoT Agent Provision API - GitHub](https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/apiary/iotagent.apib)

## Perseo FE

| PESEO FE API                  | NGSI Go commands             |
| ----------------------------- | ---------------------------- |
| POST /notices                 | (not yet implemented)        |
| GET /rules                    | rules list                   |
| GET /rules/{id}               | rules get --id {rulesId}     |
| POST /rules                   | rules create                 |
| DELETE /rules/{id}            | rules delete --id {rulesId}  |
| GET /verion                   | version                      |
| PUT /admin/log?level={level}  | admin log --level {level}    |
| GET /admin/log                | admin log                    |
| GET /admin/metrics            | admin metrics                |
| GET /admin/metrics?reset=true | admin emtrics --reset        |
| DELETE /admin/metrics         | admin metrics --delete       |

-   [Perseo FE API - GitHub](https://github.com/telefonicaid/perseo-fe/blob/master/documentation/api.md)

## Perseo CORE

| PESEO CORE API           | NGSI Go commands |
| ------------------------ | ---------------- |
| GET /perseo-core/version | version          |

## Keyrock API 

| Kerrock API                                                                                                                        | NGSI Go commands                                                                                                                 |
| ---------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| GET /version                                                                                                                       | version                                                                                                                          |
| GET /v1/auth/tokens                                                                                                                | token                                                                                                                            |
| POST /v1/auth/tokens                                                                                                               | token                                                                                                                            |
| DELETE /v1/auth/tokens                                                                                                             | (not yet implemented                                                                                                             |
| GET /v1/applications                                                                                                               | application --aid {application_id} list                                                                                          |
| POST /v1/applications                                                                                                              | application create                                                                                                               |
| GET /v1/applications/{application_id}                                                                                              | application --aid {application_id} get                                                                                           |
| DELETE /v1/applications/{application_id}                                                                                           | application --aid {application_id} delete                                                                                        |
| PATCH /v1/applications/{application_id}                                                                                            | application --aid {application_id} update                                                                                        |
| GET /v1/users                                                                                                                      | users list                                                                                                                       |
| POST /v1/users                                                                                                                     | users create                                                                                                                     |
| GET /v1/users/{user_id}                                                                                                            | users --uid {user_id} get                                                                                                        |
| DELETE /v1/users/{user_id}                                                                                                         | users --uid {user_id} delete                                                                                                     |
| PATCH /v1/users/{user_id}                                                                                                          | users --uid {user_id} update                                                                                                     |
| GET /v1/organizations                                                                                                              | organizations --oid {organization_id} list                                                                                       |
| POST /v1/organizations                                                                                                             | organizations --oid {organization_id} create                                                                                     |
| GET /v1/organizations/{organization_id}                                                                                            | organizations --oid {organization_id} get                                                                                        |
| DELETE /v1/organizations/{organization_id}                                                                                         | organizations --oid {organization_id} delete                                                                                     |
| PATCH /v1/organizations/{organization_id}                                                                                          | organizations --oid {organization_id} update                                                                                     |
| GET /v1/applications/{application_id}/roles                                                                                        | applications --aid {application_id} role --rid {role_id} list                                                                    |
| POST /v1/applications/{application_id}/roles                                                                                       | applications --aid {application_id} role --rid {role_id} create                                                                  |
| GET /v1/applications/{application_id}/roles/{role_id}                                                                              | applications --aid {application_id} role --rid {role_id} get                                                                     |
| DELETE /v1/applications/{application_id}/roles/{role_id}                                                                           | applications --aid {application_id} role --rid {role_id} delete                                                                  |
| PATCH /v1/applications/{application_id}/roles/{role_id}                                                                            | applications --aid {application_id} role --rid {role_id} update                                                                  |
| GET /v1/applications/{application_id}/permissions                                                                                  | applications --aid {application_id} permissions list                                                                             |
| POST /v1/applications/{application_id}/permissions                                                                                 | applications --aid {application_id} permissions create                                                                           |
| GET /v1/applications/{application_id}/permissions/{permission_id}                                                                  | applications --aid {application_id} permissions --pid {permission_id} get                                                        |
| DELETE /v1/applications/{application_id}/permissions/{permission_id}                                                               | applications --aid {application_id} permissions --pid {permission_id} delete                                                     |
| PATCH /v1/applications/{application_id}/permissions/{permission_id}                                                                | applications --aid {application_id} permissions --pid {permission_id} update                                                     |
| GET /v1/applications/{application_id}/pep_proxies                                                                                  | applications --aid {application_id} pep list                                                                                     |
| POST /v1/applications/{application_id}/pep_proxies                                                                                 | applications --aid {application_id} pep create                                                                                   |
| DELETE /v1/applications/{application_id}/pep_proxies                                                                               | applications --aid {application_id} pep delete                                                                                   |
| PATCH /v1/applications/{application_id}/pep_proxies                                                                                | applications --aid {application_id} pep reset                                                                                    |
| GET /v1/applications/{application_id}/iot_agents                                                                                   | applications --aid {application_id} iota list                                                                                    |
| POST /v1/applications/{application_id}/iot_agents                                                                                  | applications --aid {application_id} iota create                                                                                  |
| GET /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                   | applications --aid {application_id} iota -iid {iot_agent_id} get                                                                 |
| DELETE /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                | applications --aid {application_id} iota -iid {iot_agent_id} delete                                                              |
| PATCH /v1/applications/{application_id}/permissions/{iot_agent_id}                                                                 | applications --aid {application_id} iota -iid {iot_agent_id} reset                                                               |
| GET /v1/applications/{application_id}/roles/{role_id}/permissions                                                                  | applications --aid {application_id} role --rid {role_id}s permissions --pid {permission_id}                                      |
| POST /v1/applications/{application_id}/roles/{role_id}/permissions/{permission_id}                                                 | applications --aid {application_id} role --rid {role_id}s assign                                                                 |
| DELETE /v1/applications/{application_id}/roles/{role_id}/permissions/{permission_id}                                               | applications --aid {application_id} role --rid {role_id}s unassign                                                               |
| GET /v1/applications/{application_id}/users                                                                                        | applications --aid {application_id} users --uid {user_id} list                                                                   |
| GET /v1/applications/{application_id}/users/{user_id}/roles                                                                        | applications --aid {application_id} users --uid {user_id} get                                                                    |
| PUT /v1/applications/{application_id}/users/{user_id}/roles/{role_id}                                                              | applications --aid {application_id} users --uid {user_id} assign --rid {role_id}                                                 |
| DELETE /v1/applications/{application_id}/users/{user_id}/roles/{role_id}                                                           | applications --aid {application_id} users --uid {user_id} unassign --rid {role_id}                                               |
| GET /v1/applications/{application_id}/organizations                                                                                | applications --aid {application_id} organizations --oid {organization_id} list                                                   |
| GET /v1/applications/{application_id}/organizations/{organization_id}/roles                                                        | applications --aid {application_id} organizations --oid {organization_id} get                                                    |  
| PUT /v1/applications/{application_id}/organizations/{organization_id}/roles/{role_id}/organization_roles/{organization_role_id}    | applications --aid {application_id} organizations --oid {organization_id} assign --rid {role_id} --orid {organization_role_id}   |
| DELETE /v1/applications/{application_id}/organizations/{organization_id}/roles/{role_id}/organization_roles/{organization_role_id} | applications --aid {application_id} organizations --oid {organization_id} unassign --rid {role_id} --orid {organization_role_id} |
| GET /v1/organizations/{organization_id}/users                                                                                      | organizations --oid {organization_id} users --uid {user_id} list                                                                 |
| GET /v1/organizations/{organization_id}/users/{user_id}/organization_roles                                                         | organizations --oid {organization_id} users --uid {user_id} get                                                                  |
| PUT /v1/organizations/{organization_id}/users/{user_id}/organization_roles/{organization_role_id                                   | organizations --oid {organization_id} users --uid {user_id} create --orid {organization_role_id}                                 |
| DELETE /v1/organizations/{organization_id}/users/{user_id}/organization_roles/{organization_role_id}                               | organizations --oid {organization_id} users --uid {user_id} delete --orid {organization_role_id}                                 |
| GET /v1/applications/{application_id}/trusted_applications                                                                         | applications --aid {application_id} trusted list                                                                                 |
| PUT /v1/applications/{application_id}/trusted_applications/{trustedApplicationId}                                                  | applications --aid {application_id} trusted add --tid {trustedApplicationId}                                                     |
| DELETE /v1/applications/{application_id}/trusted_applications/{trustedApplicationId}                                               | applications --aid {application_id} trusted delete --tid {trustedApplicationId}                                                  |
| GET /v1/service_providers/configs                                                                                                  | providers                                                                                                                        |

-   [Keyrock API - GitHub](https://github.com/ging/fiware-idm/blob/master/apiary.apib)

## WireCloud

| WireCloud API                                                                    | NGSI Go commands                                                       |
| -------------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| GET /api/features                                                                | version                                                                |
| GET /api/preferences/platform                                                    | preferences get                                                        |
| POST /api/preferences/platform                                                   | (not yet implemented)                                                  |
| GET /api/workspaces                                                              | workspaces list                                                        |
| POST /api/workspaces                                                             | (not yet implemented)                                                  |
| GET /api/workspace/{workspace_id}                                                | workspaces get --wid {workspace_id}                                    |
| POST /api/workspace/{workspace_id}                                               | (not yet implemented)                                                  |
| DELETE /api/workspace/{workspace_id}                                             | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/preferences                                   | (not yet implemented)                                                  |
| PATCH /api/workspace/{workspace_id}/wiring                                       | (not yet implemented)                                                  |
| PUT /api/workspace/{workspace_id}/wiring                                         | (not yet implemented)                                                  |
| POST /api/worksace/{workspace_id}/tabs                                           | (not yet implemented)                                                  |
| GET /api/workspace/{workspace_id}/tab/{tab_id}                                   | tabs get --wid {workspace_id} --tid {tab_id}                           |
| UPDATE /api/workspace/{workspace_id}/tab/{tab_id}                                | (not yet implemented)                                                  |
| DELETE /api/workspace/{workspace_id}/tab/{tab_id}                                | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/preferences                      | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidgets                         | (not yet implemented)                                                  |
| DELETE /api/workspace/{workspace_id}/tab/{tab_id}/iwidgets                       | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}             | (not yet implemented)                                                  |
| GET /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/preferences  | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/preferences | (not yet implemented)                                                  |
| GET /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/properties   | (not yet implemented)                                                  |
| POST /api/workspace/{workspace_id}/tab/{tab_id}/iwidget/{iwidget_id}/properties  | (not yet implemented)                                                  |
| GET /api/workspace/{workspace_id}/operators/{operator_id}/variables              | (not yet implemented)                                                  |
| GET /api/resources                                                               | macs list                                                              |
| POST /api/resources                                                              | macs install {mashable_application_component_file}                     |
| GET /api/resource/{vendor}/{name}/{version}                                      | macs download --vendor {vendor} --name {name} --version {version}      |
| DELETE /api/resource/{vendor}/{name}/{version}                                   | macs uninstall --vendor {vendor} --name {name} --version {version}     |
| DELETE /api/resource/{vendor}/{name}                                             | macs uninstall --vendor {vendor} --name {name}                         |

-   [WireCloud API - GitHub](https://github.com/Wirecloud/wirecloud/blob/develop/docs/restapi/applicationmashup.apib)
