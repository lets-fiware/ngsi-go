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

FROM golang:1.15.6-alpine3.12 AS builder

WORKDIR /opt/ngsi-test/

COPY *.go /opt/ngsi-test/

RUN CGO_ENABLED=0 go build -ldflags "-w -s"

FROM ubuntu:20.04 AS runner

COPY --from=builder /opt/ngsi-test/ngsi-test /usr/local/bin/
COPY ngsi-test-config.json /root/.ngsi-test-config.json
COPY ngsi-go-config.json /root/.config/fiware/
COPY 9999_HALT /root/

WORKDIR /root/e2e

ENTRYPOINT ["/usr/local/bin/ngsi-test"]
CMD ["-verbose", "/root/9999_HALT"]