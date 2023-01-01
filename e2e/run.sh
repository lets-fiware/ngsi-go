#!/bin/sh
# MIT License
#
# Copyright (c) 2020-2023 Kazuhito Suda
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

if [ ! -e "$1" ]; then
  echo "$1 not found."
fi

type docker-compose > /dev/null
if [ $? -eq 0 ]; then
  docker-compose exec ngsi-test /usr/local/bin/ngsi-test -verbose "$1"
else
  if [ -e /usr/local/bin/ngsi-test ]; then
    /usr/local/bin/ngsi-test -verbose "$1"
  fi 
fi
