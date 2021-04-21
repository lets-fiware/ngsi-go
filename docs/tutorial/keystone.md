# NGSI Go tutorial for IDM Keystone

## Introduction

NGSI Go supports Keystone (part of the OpenStack project) as IDM. More in detail, the security framework in this case is:

* [Keystone](https://docs.openstack.org/keystone/latest) (as IDM)
* [Keypass](https://github.com/telefonicaid/fiware-keypass) (as Access Control)
* [Steelskin](https://github.com/telefonicaid/fiware-pep-steelskin) (as PEP)

Note this is the stack used by [ThinkingCities platform](https://thinking-cities.readthedocs.io/en/master/).

This tutorial describes how to use NGSI Go in this scenario.

## Preconditions

* Orion PEP is running at endpoint `http://orion-pep:1026`
* Keystone IDM is running at endpoint `http://idm:5001`
* A service named `smartgondor` is created in Keystone IDM
* A subservice named `/irrigation` is created within service `smartgondor`
* A subservice named `/watermeter` is created within service `smartgondor`
* The user `admin_smartgondor` with password `admin1234` has permissions on `/irrigation` and `/watermeter`subservices

If you are unfamiliar with the service and subservice concepts [this reference](https://thinking-cities.readthedocs.io/en/master/multitenancy/index.html)
can be useful.

## Using subservice in the IDM configuration

Create the broker using the following command:

```console
ngsi broker add \
  --host mybroker \
  --ngsiType v2 \
  --brokerHost http://orion-pep:1026 \
  --idmType ThinkingCities \
  --idmHost http://idm:5001/v3/auth/tokens \
  --username admin_smartgondor \
  --password admin1234 \
  --service smartgondor \
  --path /irrigation
```

You can now use the `mybroker` broker to do any NGSIv2 operation. For instance, to create an entity in the `/irrigation` subservice:

```console
ngsi create --host mybroker entity --data '{"id":"E", "type": "T", "A": {"value": 1, "type": "Number"}}'
```

The key point is that NGSI Go will deal transparently with all security aspects (i.e. get a token from IDM, renew token when it expires, etc.)
for the user.

More detail on NGSIv2 operations [in this side tutorial](ngsi-v2-crud.md).

## Not using subservice in the IDM configuration

As alternative, you can omit `--path` parameter in the `ngsi broker add` command. This way:

```console
ngsi broker add \
  --host mybroker \
  --ngsiType v2 \
  --brokerHost http://orion-pep:1026 \
  --idmType ThinkingCities \
  --idmHost http://idm:5001/v3/auth/tokens \
  --username admin_smartgondor \
  --password admin1234 \
  --service smartgondor \
```

This allow you to use the same broker specification (`mybroker`) for several subservices, eg:

```console
ngsi create --host mybroker --path /irrigation entity --data '{"id":"E", "type": "T", "A": {"value": 1, "type": "Number"}}'


ngsi create --host mybroker --path /watermeter entity --data '{"id":"E", "type": "T", "A": {"value": 1, "type": "Number"}}'
```

If you don't specify `--path` then `/` is used as default.
