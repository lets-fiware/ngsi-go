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

RUN \
    apk add --no-cache --update sed && \
    cat /etc/passwd | sed -n -e '/^nobody/p' > /opt/passwd && \
    chmod 444 /opt/passwd

WORKDIR /opt/csource/

COPY ./main.go /opt/csource

RUN CGO_ENABLED=0 go build -ldflags "-w -s"

FROM scratch AS runner

COPY --from=builder /opt/csource/csource /opt/csource/
COPY --from=builder /opt/passwd /etc/passwd

WORKDIR /opt/csource/

USER nobody 

ENTRYPOINT ["/opt/csource/csource"]