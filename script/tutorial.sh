#!/bin/sh
#
# MIT License
#
# Copyright (c) 2020-2022 Kazuhito Suda
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

if [ -e tutorial.sh ]; then
  cd ..
fi

if [ ! -e script/tutorial.sh ]; then
  echo "directory error"
  exit 127
fi

for i in  docker docker-compose uname make
do
  type "$i" > /dev/null
  if [ $? -ne 0 ]; then
      echo "$i not found"
      exit 127
  fi
done

if [ $(uname -m) != "x86_64" ]; then
  echo "not x86_64"
  exit 127
fi

usage="usage: tutorial.sh [start|stop|shell|ps]"

if [ $# -ne 1 ]; then
    echo "Illegal number of parameters"
    echo "$usage"
    exit 127
fi

command="$1"
case "${command}" in
    "help")
        echo "$usage"
        ;;
    "start")
        make build
        cd e2e
        if [ ! "$(docker-compose ps -a | wc -l)" = "2" ]; then
            make down
        fi
        make build
        make rmi
        make up
        ./run.sh cases/0000_prepare/
        ./run.sh cases/4000_ngsi-v2/0000_import-data.test
        ./run.sh cases/6000_time_series/1000_comet/0002_create_historical_data.test
        ./run.sh cases/6000_time_series/2000_quantumleap/0003_create_historical_data.test
        ;;
    "up")
        make build
        cd e2e
        if [ ! "$(docker-compose ps -a | wc -l)" = "2" ]; then
            make down
        fi
        make build
        make rmi
        make up
        ./run.sh cases/0000_prepare/
        ;;
    "ps")
        cd e2e
        make ps
        ;;
    "stop")
        cd e2e
        make down
        ;;
    "shell")
        cd e2e
        make shell
        ;;
    *)
        echo "Command not Found."
        echo "$usage"
        exit 127
        ;;
esac
