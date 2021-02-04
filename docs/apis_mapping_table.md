# FIWARE Open APIs mapping table

These tables shows the mapping from FIWARE Open APIs to NGSI Go commands.

## STH-Comet APIs

| STH-Comet APIs                                                                                             | NGSI Go commands                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| get /version                                                                                               | version                                                                                                         |
| get /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&hLimit={n}&hOffset={n}                  | hget attr --hLimit {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                |
| get /STH/v2/entities/{entityId}/attrs/{attrName}?entityType={entityType}&lastN={n}                         | hget attr --lastN {n} --type {entityType} --id {enttiyId} --attrName {attrName}                                 |
| get /STH/v2/entities/{entityId}/attrs/{attrName}?type={entityType}&aggrMethod={method}&aggrPeriod={period} | hget attr --arrgMethod {method} --aggrPeriod {period} --type {entityType} --id {enttiyId} --attrName {attrName} |
| delete /STH/v1/contextEntities                                                                             | hdelete                                                                                                         |
| delete /STH/v1/contextEntities/type/{entityType}/id/{entityId}                                             | hdelete --type {entityType} --id {enttiyId}                                                                     |
| delete /STH/v1/contextEntities/type/{entityType}/id/{entityId}/attributes/{attrName}                       | hdelete --type {entityType} --id {enttiyId} --attrName {attrName}                                               |

-   [STH-Comet APIs - GitHub](https://github.com/telefonicaid/fiware-sth-comet/blob/master/apiary.apib)

## QuantumLeap APIs

| QuantumLeap APIs                                   | NGSI Go commands                                                       |
| -------------------------------------------------- | ---------------------------------------------------------------------- |
| get /v2/                                           | apis                                                                   |
| get /version                                       | version                                                                |
| post /config                                       | (not yet implemented)                                                  |
| get /health                                        | health                                                                 |
| post /notify                                       | (not yet implemented)                                                  |
| post /subscribe                                    | (not yet implemented)                                                  |
| get /v2/entities                                   | hget entities                                                          |
| get /v2/entities/{entityId}/attrs/{attrName}       | hget attr --id {entityId} --attrName {attrName}                        |
| get /v2/entities/{entityId}/attrs/{attrName}/value | hget attr --id {entityId} --attrName {attrName} --value                |
| get /v2/entities/{entityId}                        | hget attrs --id {entityId}                                             |
| get /v2/entities/{entityId}/value                  | hget attrs --id {entityId} --value                                     |
| get /v2/types/{entityType}/attrs/{attrName}        | hget attr --sameType --type {entityType} --attrName {attrName}         |
| get /v2/types/{entityType}/attrs/{attrName}/value  | hget attr --sameType --type {entityType} --attrName {attrName} --value |
| get /v2/types/{entityType}                         | hget attrs --sameType --type {entityType}                              |
| get /v2/types/{entityType}/value                   | hget attrs --sameType --type {entityType} --value                      |
| get /v2/attrs/{attrName}                           | hget attr --nTypes --attrName {attrName}                               |
| get /v2/attrs/{attrName}/value                     | hget attr --nTypes --attrName {attrName} --value                       |
| get /v2/attrs                                      | hget attrs --nTypes                                                    |
| get /v2/attrs/value                                | hget attrs --nTypes --value                                            |
| delete /v2/entities/{entityId}                     | hdelete entity --id {entityId}                                         |
| delete /v2/types/{entityType}                      | hdelete entities --type {entityType}                                   |

-   [QuantumLeap APIs - GitHub](https://github.com/smartsdk/ngsi-timeseries-api/blob/master/specification/quantumleap.yml)
