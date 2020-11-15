#!/bin/sh
git clone "https://github.com/lets-fiware/ngsi-go.git"
cd ngsi-go
rm -fr build
ln -s /build
rm -fr build/*
make devel-deps
echo "TARGET: $NGSIGO_TARGET"
make $NGSIGO_TARGET
chown -R $USER_ID:$GROUP_ID /build
