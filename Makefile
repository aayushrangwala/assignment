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

# A Verb with some commands under it is called as target in Makefile.
# Target is used to run as an argument along with "make" command. It basically runs the commands defined under it

# test is the target to run the tests for all the directories and sub directories
test:
	go test -v ./...

# run is the target used to compile and build the program (main.go) by calling the 'build' target and run
run: build
	./www

# build target is used to only to compile and build the program (main.gos)
build:
	go build -o www
