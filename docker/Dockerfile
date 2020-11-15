FROM ubuntu:20.04

RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        tar \
        git \
        curl \
        ca-certificates \
        make

RUN \
    mkdir -p /root/go/src/github.com/lets-fiware && \
    curl -L "https://golang.org/dl/go1.15.linux-amd64.tar.gz" -O && \
    tar xvf /go1.15.linux-amd64.tar.gz -C /usr/local

COPY build.sh /root/go/src/github.com/lets-fiware

ENV PATH $PATH:/usr/local/go/bin:/root/go/bin

WORKDIR /root/go/src/github.com/lets-fiware

CMD ["./build.sh"]
