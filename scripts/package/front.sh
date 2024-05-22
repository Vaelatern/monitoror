#!/usr/bin/env bash
# Do not use this script manually, Use makefile

set -e

#######################################################
# This script is used to package ui/dist in go source #
#######################################################

go get github.com/GeertJohan/go.rice
rm -f service/rice-box.go
rice embed-go -i ./service
