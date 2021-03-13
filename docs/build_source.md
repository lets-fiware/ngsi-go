# Building from sources

The building process of binaries are run in a Docker container.
So that you do not need to prepare a Golang building environment.

```console
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go
make release
```

The binaries will be put in `build/` directory.

## Linux AMD64

```console
make linux_amd64
```

## Linux ARM64

```console
make linux_arm64
```

## Linux ARM

```console
make linux_arm
```

## Darwin AMD64

```console
make darwin_amd64
```

## Darwin ARM64

```console
make darwin_arm64
```
