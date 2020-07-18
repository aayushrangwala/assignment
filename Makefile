# If the "make" command is run without any argument, the first default goal target will run.
# A default goal is the first target without a (.) at the begining of its name.
# The default goal can be override by specifying ".DEFAULT_GOAL=<target>".
# The convention is to keep "all: <all targets, space separated>" at the begining of the Makefile

# export <ENV_VAR> command in make exports that env var with its value

# CGO_ENABLED is an env var used at the time of compiling and building the programs.
# It needs to be enabled (1) for cross compiling and disabled (0) for native builds
export CGO_ENABLED=0

# GO111MODULE is the env var used by the mod tool (go.mod file) is useful for enabling the module behaviour
export GO111MODULE=on

# Declaring the binary name
GO_APP_BINARY ?= www

# A Verb with some commands under it is called as target in Makefile.
# Target is used to run as an argument along with "make" command. It basically runs the commands defined under it

# all target runs all the targets specified.
all: lint test coverage run clean

# fmt target to format the go code
fmt:
	go fmt ./...

# vet to run the vet linter on the go code
vet:
	go vet ./...

# lint target is used to run the golangci-lint binary to check for the linting errors
lint:
	golangci-lint run --skip-dirs='(vendor)' -vc ./.golangci.yaml ./...

# yaml-lint will run the linter for all the yaml files in the root directory or sub-directory
yaml-lint:
	yamllint -c .yamllint.conf ./

# test is the target to run the tests for all the directories and sub directories
test:
	go test -v ./... -coverprofile coverage.out

# coverage taret will run a script which will check the test coverage of the project, if it is greater than 85% or not
coverage:
	scripts/gocoverage.sh

# run is the target used to compile and build the program (main.go) by calling the 'build' target and run
run: build
	./$(GO_APP_BINARY)

# build target is used to only to compile and build the program (main.go) with running fmt and vet targets also
build: clean fmt vet
	go build -o $(GO_APP_BINARY)

# clean is the target which will clean the object files in the temporary source directory and the binary
# ,which are created at the time of build
clean:
	go clean
	rm -f $(GO_APP_BINARY)

# .PHONY is a special built in target which is used to specify the target names explicitely
# so that it is not conflicted with the file names and also it improves performance
.PHONY: clean lint test coverage build run
