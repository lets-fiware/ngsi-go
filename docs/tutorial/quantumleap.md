# NGSI Go tutorial for QuantumLeap

## Get settings of alias

```console
ngsi server get --host quantumleap
```

```json
serverType quantumleap
serverHost http://localhost:8668
FIWARE-Serivce openiot
FIWARE-SerivcePath /
```

## Print version

```console
ngsi version --host quantumleap
```

```json
{
  "version": "0.7.6"
}
```

## Print health status

```console
ngsi health --host quantumleap
```

```json
{
  "status": "pass"
}
```

## List of all the entityId

### Query the entities

```console
ngsi hget entities
```

```json
[
  {
    "id": "Event001",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  },
  {
    "id": "Event002",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  },
  {
    "id": "device001",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Thing"
  },
  {
    "id": "device002",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Thing"
  }
]
```

### Query the entities with an entity type

```console
ngsi hget entities --type Event
```

```json
[
  {
    "id": "Event001",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  },
  {
    "id": "Event002",
    "index": [
      "2016-09-13T01:39:00.000+00:00"
    ],
    "type": "Event"
  }
]
```

## History of an attribute  - hget attr

### History of an attribute of a given entity instance

```console
ngsi hget attr --id device001 --attrName A1 --lastN 3
```

```json
{
  "attrName": "A1",
  "entityId": "device001",
  "index": [
    "2016-09-13T01:37:00.000+00:00",
    "2016-09-13T01:38:00.000+00:00",
    "2016-09-13T01:39:00.000+00:00"
  ],
  "values": [
    98.0,
    99.0,
    100.0
  ]
}
```

### History values of an attribute of a given entity instance

```console
ngsi hget attr --id device001 --attrName A1 --lastN 3 --value
```

```json
{
  "index": [
    "2016-09-13T01:37:00.000+00:00",
    "2016-09-13T01:38:00.000+00:00",
    "2016-09-13T01:39:00.000+00:00"
  ],
  "values": [
    98.0,
    99.0,
    100.0
  ]
}
```

### History of an attribute of N entities of the same type

```console
ngsi hget attr --sameType --type Thing --attrName A2 --hLimit 4
```

```json
{
  "attrName": "A2",
  "entities": [
    {
      "entityId": "device001",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ],
      "values": [
        2.0,
        3.0
      ]
    },
    {
      "entityId": "device002",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ],
      "values": [
        4.0,
        5.0
      ]
    }
  ],
  "entityType": "Thing"
}
```

### History values of an attribute of N entities of the same type

```console
ngsi hget attr --sameType --type Thing --attrName A2 --hLimit 4 --value
```

```json
{
  "values": [
    {
      "entityId": "device001",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ],
      "values": [
        2.0,
        3.0
      ]
    },
    {
      "entityId": "device002",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ],
      "values": [
        4.0,
        5.0
      ]
    }
  ]
}
```

### History of an attribute of N entities of N types

```console
ngsi hget attr --nTypes --attrName A2 --hLimit 4
```

```json
{
  "attrName": "A2",
  "types": [
    {
      "entities": [
        {
          "entityId": "Event001",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            6.0,
            7.0
          ]
        },
        {
          "entityId": "Event002",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            8.0,
            9.0
          ]
        }
      ],
      "entityType": "Event"
    },
    {
      "entities": [
        {
          "entityId": "device001",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            2.0,
            3.0
          ]
        },
        {
          "entityId": "device002",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            4.0,
            5.0
          ]
        }
      ],
      "entityType": "Thing"
    }
  ]
}
```

### History values of of an attribute of N entities of N types

```console
ngsi hget attr --nTypes --attrName A2 --hLimit 4 --value
```

```json
{
  "values": [
    {
      "entities": [
        {
          "entityId": "Event001",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            6.0,
            7.0
          ]
        },
        {
          "entityId": "Event002",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            8.0,
            9.0
          ]
        }
      ],
      "entityType": "Event"
    },
    {
      "entities": [
        {
          "entityId": "device001",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            2.0,
            3.0
          ]
        },
        {
          "entityId": "device002",
          "index": [
            "2016-09-13T00:00:00.000+00:00",
            "2016-09-13T00:01:00.000+00:00"
          ],
          "values": [
            4.0,
            5.0
          ]
        }
      ],
      "entityType": "Thing"
    }
  ]
}
```

## History of N attributes - hget attrs

### History of N attributes of a given entity instance.

```console
ngsi hget attrs --id device001 --attrs A1,A2 --hLimit 3
```

```json
{
  "attributes": [
    {
      "attrName": "A1",
      "values": [
        1.0,
        2.0,
        3.0
      ]
    },
    {
      "attrName": "A2",
      "values": [
        2.0,
        3.0,
        4.0
      ]
    }
  ],
  "entityId": "device001",
  "index": [
    "2016-09-13T00:00:00.000+00:00",
    "2016-09-13T00:01:00.000+00:00",
    "2016-09-13T00:02:00.000+00:00"
  ]
}
```

### History values of N attributes of a given entity instance.

```console
ngsi hget attrs --id device001 --attrs A1,A2 --hLimit 3 --value
```

```json
{
  "attributes": [
    {
      "attrName": "A1",
      "values": [
        1.0,
        2.0,
        3.0
      ]
    },
    {
      "attrName": "A2",
      "values": [
        2.0,
        3.0,
        4.0
      ]
    }
  ],
  "index": [
    "2016-09-13T00:00:00.000+00:00",
    "2016-09-13T00:01:00.000+00:00",
    "2016-09-13T00:02:00.000+00:00"
  ]
}
```

### History of N attributes of N entities of the same type

```console
ngsi hget attrs --sameType --type Thing --attrs A1,A2 --hLimit 4
```

```json
{
  "entities": [
    {
      "attributes": [
        {
          "attrName": "A1",
          "values": [
            1.0,
            2.0
          ]
        },
        {
          "attrName": "A2",
          "values": [
            2.0,
            3.0
          ]
        }
      ],
      "entityId": "device001",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ]
    },
    {
      "attributes": [
        {
          "attrName": "A1",
          "values": [
            null,
            null
          ]
        },
        {
          "attrName": "A2",
          "values": [
            4.0,
            5.0
          ]
        }
      ],
      "entityId": "device002",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ]
    }
  ],
  "entityType": "Thing"
}
```

### History values of N attributes of N entities of the same type

```console
ngsi hget attrs --sameType --type Thing --attrs A1,A2 --hLimit 4 --value
```

```json
{
  "values": [
    {
      "attributes": [
        {
          "attrName": "A1",
          "values": [
            1.0,
            2.0
          ]
        },
        {
          "attrName": "A2",
          "values": [
            2.0,
            3.0
          ]
        }
      ],
      "entityId": "device001",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ]
    },
    {
      "attributes": [
        {
          "attrName": "A1",
          "values": [
            null,
            null
          ]
        },
        {
          "attrName": "A2",
          "values": [
            4.0,
            5.0
          ]
        }
      ],
      "entityId": "device002",
      "index": [
        "2016-09-13T00:00:00.000+00:00",
        "2016-09-13T00:01:00.000+00:00"
      ]
    }
  ]
}
```

### History of N attributes of N entities of N types

```console
ngsi hget attrs --nTypes --hLimit 2
```

```json
{
  "attrs": [
    {
      "attrName": "A1",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                5.0
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                1.0
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    },
    {
      "attrName": "A2",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                6.0
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                8.0
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                2.0
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                4.0
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    },
    {
      "attrName": "A3",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                7.0
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                3.0
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    }
  ]
}
```

### History values of N attributes of N entities of N types

```console
ngsi hget attrs --nTypes --hLimit 2 --value
```

```json
{
  "values": [
    {
      "attrName": "A1",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                5.0
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                1.0
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    },
    {
      "attrName": "A2",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                6.0
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                8.0
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                2.0
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                4.0
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    },
    {
      "attrName": "A3",
      "types": [
        {
          "entities": [
            {
              "entityId": "Event001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            },
            {
              "entityId": "Event002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                7.0
              ]
            }
          ],
          "entityType": "Event"
        },
        {
          "entities": [
            {
              "entityId": "device001",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                null
              ]
            },
            {
              "entityId": "device002",
              "index": [
                "2016-09-13T00:00:00.000+00:00"
              ],
              "values": [
                3.0
              ]
            }
          ],
          "entityType": "Thing"
        }
      ]
    }
  ]
}
```

## Delete historical data

### Delete historical data of a certain entity

```console
ngsi hdelete entity --id device003 --run
```

### Delete historical data of all entities of a certain type

```console
ngsi hdelete entities --type Event --run
```
