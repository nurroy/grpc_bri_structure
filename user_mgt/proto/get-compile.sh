#!/bin/bash

echo "Starting compile proto"

SRC=$(realpath $(cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd ))

GITVER=$(git describe --exact-match 2> /dev/null || echo "`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:\"%h\" -1`")

set -ve

./get-googleapi.sh

pushd $SRC &> /dev/null

#export PATH=$PATH:$SRC/node_modules/.bin
export PATH=$PATH:$GOPATH/bin

for VER in v*; do
  pushd $VER &> /dev/null

  for i in $(find ./ -mindepth 1 -maxdepth 1 -type d); do
    SERVICE=$(basename $i)

    pushd $SRC/$VER/$SERVICE &> /dev/null

    protoc \
      -I. \
      -I/usr/local/include \
      -I${GOPATH}/src \
      -I${GOPATH}/src/user_mgt/proto \
      -I${GOPATH}/src/user_mgt/proto/lib \
      --go_out=plugins=grpc:$GOPATH/src/user_mgt/proto/v1/users \
      --swagger_out=logtostderr=true:$GOPATH/src/user_mgt/swagger \
      --grpc-gateway_out=logtostderr=true:$GOPATH/src/user_mgt/proto/v1/users \
      *.proto

    popd &> /dev/null
  done

  popd &> /dev/null
done

# remove docustomer.json and generate merge swagger json file and remove each raw files
rm -rf $GOPATH/src/bitbucket.org/customer/swagger/docs.json
jq -s 'reduce .[] as $item ({}; . * $item)|del(.info)|del(..|.tags?)' \
$GOPATH/src/bitbucket.org/customer/swagger/*swagger.json >> $GOPATH/src/bitbucket.org/customer/swagger/docs.json
rm -rf $GOPATH/src/bitbucket.org/customer/swagger/*swagger.json


# build/install go output
for VER in v*; do
  pushd $VER &> /dev/null

  for i in $(find ./ -mindepth 1 -maxdepth 1 -type d); do
    SERVICE=$(basename $i)

    pushd $SRC/$VER/$SERVICE &> /dev/null
    go install
    popd &> /dev/null
  done

  popd &> /dev/null
done

popd &> /dev/null

echo "Successfully compile proto"