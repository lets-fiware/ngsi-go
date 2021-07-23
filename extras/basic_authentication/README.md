# Access a broker with Basic authentication

This documentation describes how to access an Orion Context Broker that the endpoints of NGSI API
are protected by Basic Authentication.

## Prepare

Clone the NGSI Go repository and move to `ngsi-go/extras/basic_authentication` directory.

```
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go/extras/basic_authentication
```

## Add user

Add username and password to the htpasswd file.

```
NAME=fiware PASS='1234'; echo $NAME:`openssl passwd -6 $PASS` >> htpasswd
```

## Start up

Start up an Orion context broker.

```
docker-compose up -d
```

You can see three containers.

```
            Name                          Command               State                Ports
--------------------------------------------------------------------------------------------------------
basic_authentication_mongo_1   docker-entrypoint.sh --noj ...   Up      27017/tcp
basic_authentication_nginx_1   /docker-entrypoint.sh ngin ...   Up      0.0.0.0:80->80/tcp,:::80->80/tcp
basic_authentication_orion_1   /usr/bin/contextBroker -fg ...   Up      1026/tcp
```

### Sanity check

Check the service is ready by executing the following command.

```
curl localhost:1026/version
```

```
<html>
<head><title>401 Authorization Required</title></head>
<body>
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx/1.19.10</center>
</body>
</html>
```


## Add the broker

Add the broker to NGSI Go configuration.

```
ngsi broker add \
  --host orion-with-basic-auth \
  --ngsiType v2 \
  --brokerHost http://localhost:1026/ \
  --idmType basic \
  --username fiware \
  --password 1234
```

## Access the broker

The following command allows you to access the broker with Basic authentication.

```
ngsi version --host orion-with-basic-auth
```
