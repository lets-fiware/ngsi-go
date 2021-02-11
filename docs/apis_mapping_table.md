# FIWARE Open APIs mapping table

These tables shows the mapping from FIWARE Open APIs to NGSI Go commands.

## STH-Comet API 

| STH-Comet API                                                                                             | NGSI Go commands                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| GET /version                                                                                               | version                                                                                                         |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&hLimit={n}&hOffset={n}                  | hget attr --hLimit {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?entityType={entityType}&lastN={n}                         | hget attr --lastN {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                 |
| GET /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&aggrMethod={method}&aggrPeriod={period} | hget attr --arrgMethod {method} --aggrPeriod {period} --type {entityType} --id {enttiyId} --attrName {attrName} |
| DELETE /STH/v1/contextEntities                                                                             | hdelete                                                                                                         |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}                                             | hdelete --type {entityType} --id {enttiyId}                                                                     |
| DELETE /STH/v1/contextEntities/type/{entityType}/id/{entityId}/attributes/{attrName}                       | hdelete --type {entityType} --id {enttiyId} --attrName {attrName}                                               |

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
| GET /v2/entities/{entityId}/attrs/{attrName}       | hget attr --id {entityId} --attrName {attrName}                        |
| GET /v2/entities/{entityId}/attrs/{attrName}/value | hget attr --id {entityId} --attrName {attrName} --value                |
| GET /v2/entities/{entityId}                        | hget attrs --id {entityId}                                             |
| GET /v2/entities/{entityId}/value                  | hget attrs --id {entityId} --value                                     |
| GET /v2/types/{entityType}/attrs/{attrName}        | hget attr --sameType --type {entityType} --attrName {attrName}         |
| GET /v2/types/{entityType}/attrs/{attrName}/value  | hget attr --sameType --type {entityType} --attrName {attrName} --value |
| GET /v2/types/{entityType}                         | hget attrs --sameType --type {entityType}                              |
| GET /v2/types/{entityType}/value                   | hget attrs --sameType --type {entityType} --value                      |
| GET /v2/attrs/{attrName}                           | hget attr --nTypes --attrName {attrName}                               |
| GET /v2/attrs/{attrName}/value                     | hget attr --nTypes --attrName {attrName} --value                       |
| GET /v2/attrs                                      | hget attrs --nTypes                                                    |
| GET /v2/attrs/value                                | hget attrs --nTypes --value                                            |
| DELETE /v2/entities/{entityId}                     | hdelete entity --id {entityId}                                         |
| DELETE /v2/types/{entityType}                      | hdelete entities --type {entityType}                                   |

-   [QuantumLeap API - GitHub](https://github.com/smartsdk/ngsi-timeseries-api/blob/master/specification/quantumleap.yml)

## IoT Agent Provision API

| IoT Agent Provision API     | NGSI Go commands                |
| --------------------------- | ------------------------------- |
| GET /services               | services list                   |
| POST /services              | services create                 |
| PUT /serivces               | services update                 |
| DELETE /services            | services delete                 |
| GET /devices                | devices list                    |
| GET /devices/{device_id}    | devices get --id {device_id}    |
| POST /devices/{device_id}   | devices create --id {device_id} |
| PUT /devices/{device_id}    | devices update --id {device_id} |
| DELETE /devices/{device_id} | devices delete --id {device_id} |

-   [IoT Agent Provision API - GitHub](https://github.com/telefonicaid/iotagent-node-lib/blob/master/doc/apiary/iotagent.apib)

## PERSEO FE

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

-   [PERSEO FE API - GitHub](https://github.com/telefonicaid/perseo-fe/blob/master/documentation/api.md)

## PERSEO CORE

| PESEO CORE API           | NGSI Go commands |
| ------------------------ | ---------------- |
| GET /perseo-core/version | version          |
