#!/bin/sh

ls -al ./
go tool nm ./hashapi | grep -i crypto
./hashapi

