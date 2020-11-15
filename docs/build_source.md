# Building from sources

The building process of binaries are run in a Docker container.
So that you don't need to prepare a Golang building environment.

```bash
git clone https://github.com/lets-fiware/ngsi-go.git
cd ngsi-go
make release
```

The binaries will be put in `build/` directory.

## Linux AMD64

```bash
make linux_amd64
```

## Linux ARM64

```bash
make linux_arm64
```

## Linux ARM

```bash
make linux_arm
```

## Darwin

```bash
make darwin_amd64
```
