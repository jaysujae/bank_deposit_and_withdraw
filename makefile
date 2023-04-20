.PHONY: all run-tigerbeetle-format run-tigerbeetle-start run-temporalite-start run-tctl-register stop run-create-accounts

all: run-tigerbeetle-start run-temporalite-start run-tctl-register

run-tigerbeetle-format:
	tigerbeetle format --cluster=0 --replica=0 --replica-count=1 0_0.tigerbeetle

run-tigerbeetle-start:
	tigerbeetle start --addresses=3000 0_0.tigerbeetle

run-temporalite-start:
	temporalite start --ephemeral

run-tctl-register:
	tctl --ns default namespace register

run-create-accounts:
	node createAccounts.js
