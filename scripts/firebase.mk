##
##                Firebase Emulator
##
##  This Makefile module contains implementation of:
##    - shortcut to run firestore emulator
##    - start/stop operations on Firestore emulator in detached mode
##
##  It should be used in serive Makefile to run emulator pre-test and stop it post-test.
##
##  \e[1mFirebase Emulator variables\e[0m
##   \e[34mEMULATOR_PORT\e[0m
##       Here can be changed the host-port of emulator
EMULATOR_PORT=42042
##
##  \e[1mFirebase Emulator targets\e[0m
##   \e[34mfirestore_emulator\e[0m
##       Starts the firestore emulator in normal mode. Useful when you are running tests from GoLand.
##       Also if you can open this one in the sepparate terminal and then 'go test ./...' command works just fine.
firestore_emulator:
	gcloud beta emulators firestore start --host-port=localhost:${EMULATOR_PORT}

##   \e[34mfirestore_emulator/start\e[0m
##       Starts the firestore emulator in detached mode. This one is used by `make test` to run test without
##       Requiring to start emulator manually. It is starting it on background. which makes this a good
##       candidate for CI/CD pipeline
firestore_emulator/start:
	gcloud beta emulators firestore start --host-port=localhost:${EMULATOR_PORT} >.firestore.logs 2>&1 & \

##   \e[34mfirestore_emulator/stop\e[0m
##       Stops the running firestore emulator in detached mode. See '\e[34mfirestore_emulator/start\e[0m'
firestore_emulator/stop:
	set -x; \
	PIDS=`ps -aux | awk "/--port=$(EMULATOR_PORT)/ && !/awk/" | awk -F' ' '{print $$2}' | tr '\n' ' '`; \
	if [[ "$${PIDS}" != "" ]]; then \
	     echo "INFO: Stopping Firestore Emulator with PIDS: $${PIDS}"; \
	     kill $${PIDS}; \
	fi
