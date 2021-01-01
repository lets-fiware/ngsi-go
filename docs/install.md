# Install

## Install NGSI Go binary

The NGSI Go binary is installed in `/usr/local/bin`.

### Installation on UNIX

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.5.0/ngsi-v0.5.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.5.0-linux-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.5.0-linux-arm.tar.gz` and `ngsi-v0.5.0-linux-arm64.tar.gz` binaries are also available.

### Installation on Mac

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.5.0/ngsi-v0.5.0-darwin-amd64.tar.gz
sudo tar zxvf ngsi-v0.5.0-darwin-amd64.tar.gz -C /usr/local/bin
```

## Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```console
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```
