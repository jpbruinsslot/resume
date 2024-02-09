default: build

# -timeout 	timeout in seconds
#  -v		verbose output
test:
	@ echo "+ $@"
	@ go test -timeout=5s -v

# `CGO_ENABLED=0`
# Because of dynamically linked libraries, this will statically compile the
# app with all libraries built in. You won't be able to cross-compile if CGO
# is enabled. This is because Go binary is looking for libraries on the
# operating system it’s running in. We compiled our app, but it still is
# dynamically linked to the libraries it needs to run
# (i.e., all the C libraries it binds to). When using a minimal docker image
# the operating system doesn't have these libraries.
#
# `-a`
# Force rebuilding of package, all import will be rebuilt with cgo disabled,
# which means all the imports will be rebuilt with cgo disabled.
#
# `-installsuffix cgo`
# A suffix to use in the name of the package installation directory
#
# `-o`
# Output
#
# `./bin/[name-of-app]`
# Placement of the binary
#
# `./cmd/[name-of-app]`
# Location of the cmd package
build:
	@ echo "+ $@"
	@ CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/resume ./cmd/resume

build-linux:
	@ echo "+ $@"
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/resume-linux-amd64 ./cmd/resume

build-darwin:
	@ echo "+ $@"
	@ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/resume-darwin-amd64 ./cmd/resume

build-windows:
	@ echo "+ $@"
	@ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o ./bin/resume-windows-amd64 ./cmd/resume

.PHONY: default test build build-linux build-darwin build-windows
