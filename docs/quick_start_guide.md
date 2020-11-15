# NGSI Go Quick Start Guide

## Install

### Install NGSI Go binary

Install NGSI Go binary in `/usr/local/bin`.

```
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.1.0/ngsi-v0.1.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.1.0-linux-amd64.tar.gz -C /usr/local/bin
```

### Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

## Run

You can get the version of your context broker instance as shown:

```json
$ ngsi version -h localhost:1026
{
"orion" : {
  "version" : "2.5.0",
  "uptime" : "0 d, 5 h, 7 m, 50 s",
  "git_hash" : "63cc107657ae10aa03f1c83bdea0be869d8e26a1",
  "compile_time" : "Fri Oct 30 09:02:37 UTC 2020",
  "compiled_by" : "root",
  "compiled_in" : "320890801dd4",
  "release_date" : "Fri Oct 30 09:02:37 UTC 2020",
  "doc" : "https://fiware-orion.rtfd.io/en/2.5.0/",
  "libversions": {
     "boost": "1_53",
     "libcurl": "libcurl/7.29.0 NSS/3.44 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.0.2k",
     "rapidjson": "1.1.0",
     "mongodriver": "legacy-1.1.2"
  }
}
}
```

You can register an alias to access the broker.

```
ngsi broker add --host letsfiware --brokerHost http://localhost:1026 --ngsiType v2
```

You can get the version by using the alias `letsfiware`.

```
$ ngsi version -h letsfiware
{
"orion" : {
  "version" : "2.5.0",
  "uptime" : "0 d, 5 h, 7 m, 50 s",
  "git_hash" : "63cc107657ae10aa03f1c83bdea0be869d8e26a1",
  "compile_time" : "Fri Oct 30 09:02:37 UTC 2020",
  "compiled_by" : "root",
  "compiled_in" : "320890801dd4",
  "release_date" : "Fri Oct 30 09:02:37 UTC 2020",
  "doc" : "https://fiware-orion.rtfd.io/en/2.5.0/",
  "libversions": {
     "boost": "1_53",
     "libcurl": "libcurl/7.29.0 NSS/3.44 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.0.2k",
     "rapidjson": "1.1.0",
     "mongodriver": "legacy-1.1.2"
  }
}
}
```

Once you access the broker, you can omit to specify the broker.

```
ngsi version
```

If you want to check the current default settings, you can run the following command.

```
ngsi settings list
```
