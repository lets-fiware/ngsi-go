# MIT License
#
# Copyright (c) 2020-2021 Kazuhito Suda
#
# This file is part of NGSI Go
#
# https://github.com/lets-fiware/ngsi-go
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

FROM alpine:3.14 AS builder

RUN \
    apk update && \
    apk add --no-cache tar curl && \
    curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.8.4/ngsi-v0.8.4-linux-amd64.tar.gz && \
    tar zxvf ngsi-v0.8.4-linux-amd64.tar.gz -C /usr/local/bin

FROM alpine:3.14 AS runner

COPY --from=builder /usr/local/bin/ngsi /usr/local/bin/ngsi
COPY ./ngsi-go-config.json /opt/ngsi-go/ngsi-go-config.json

WORKDIR /opt/ngsi-go/

RUN chown -R nobody:nobody /opt/ngsi-go

USER nobody 

ENTRYPOINT ["/usr/local/bin/ngsi"]
CMD ["--stderr", "info", "--config", "./ngsi-go-config.json", "--cache", "ngsi-go-token-cache.json", "regproxy", "--host", "remote-orion", "--verbose"]
