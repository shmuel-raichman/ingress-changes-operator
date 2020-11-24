#!/bin/bash
# B"H


IMGD=smuel770/golang-operator-sdk:$1

make docker-build IMG=$IMGD
make docker-push IMG=$IMGD
make deploy IMG=$IMGD