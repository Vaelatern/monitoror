#!/usr/bin/env bash
# Do not use this script manually, Use makefile

set -e

# gotestsum, used by `make test`. Test utilities
echo "Installing gotestsum"
go install gotest.tools/gotestsum@latest

# rice, used by `make build`. Embed UI dist into go binary
echo "Installing rice"
go install github.com/GeertJohan/go.rice/rice@latest

# mockery, used by `make mocks`. Generating mock for backend
echo "Installing mockery"
go install github.com/vektra/mockery/.../

# revproxy, usef by `make proxy`. Start a proxy for test
echo "Installing revproxy"
go install github.com/jsdidierlaurent/revproxy@latest
