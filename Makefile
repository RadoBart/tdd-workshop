##             Kvasar permweb assistantMakefile
##
##  Simple Makefile containing implementation of targets for generating protobuf file and various targets useful during development.
##
##  To generate swagger, angular models and go messages&grpc use this command:
##    $ make gen
##  To run unittests, use:
##    $ make test
SHELL=/bin/bash
MAKEFILE_SOURCES=Makefile scripts/firebase.mk
MAKEFLAGS += -j$(nproc)
SHELL=/bin/bash
##
##  \e[1mTargets\e[0m
##   \e[34mhelp\e[0m
##       Shows this help
help:
	@echo -e "$$(sed -n 's/^##//p' Makefile)"

setup_test: firestore_emulator

##   \e[34mtest_%\e[0m
##       Runs all tests in the specified service
test_%:
	make firestore_emulator/start EMULATOR_PORT=$(EMULATOR_PORT)
	sleep 4
	set -x; \
	export FIRESTORE_EMULATOR_HOST=localhost:$(EMULATOR_PORT); \
	export GCLOUD_PROJECT=$(GCP_PROJECT); \
	pushd $*/;\
	go test ./... -v; \
	RESULT=$$!; \
	make firestore_emulator/stop EMULATOR_PORT=$(EMULATOR_PORT); \
	exit $${RESULT}

include scripts/firebase.mk
