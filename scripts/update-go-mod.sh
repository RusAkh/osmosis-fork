#!/bin/bash

# Script for checking `git diff` between two commits and updating osmoutils, osmomath, epochs or ibc-hooks if any were changed between two commits
# Used by Go Mod Auto Version Update workflow
# First argument: sha of a first commit
# Second argument: sha of a second commit

is_updated() {
	if [ "${1}" != "" ]
	then
		return 1
	fi
	return 0
}

commit_before=$1
commit_after=$2

changed_osmoutils=$(git diff --name-only $commit_before $commit_after | grep osmoutils)
changed_osmomath=$(git diff --name-only $commit_before $commit_after | grep osmomath)
changed_ibc_hooks=$(git diff --name-only $commit_before $commit_after | grep x/ibc-hooks)
changed_epochs=$(git diff --name-only $commit_before $commit_after | grep x/epochs)

is_updated $changed_osmoutils
update_osmoutils=$?

is_updated $changed_osmomath
update_osmomath=$?

is_updated $changed_ibc_hooks
update_ibc_hooks=$?

is_updated $changed_epochs
update_epochs=$?

if [ $update_osmoutils -eq 1 ]
then 
	sed -i 's|module github.com/osmosis-labs/osmosis/v15|module github.com/osmosis-labs/osmosis/v15 // changed osmoutils |g' go.mod
fi

if [ $update_osmomath -eq 1 ]
then 
	sed -i "s|module github.com/osmosis-labs/osmosis/v15|module github.com/osmosis-labs/osmosis/v15 // changed osmomath ($commit_before) ($commit_after) |g" go.mod
fi

if [ $update_ibc_hooks -eq 1 ]
then 
	sed -i 's|module github.com/osmosis-labs/osmosis/v15|module github.com/osmosis-labs/osmosis/v15 // changed ibc-hooks |g' go.mod
fi

if [ $update_epochs -eq 1 ]
then 
	sed -i 's|module github.com/osmosis-labs/osmosis/v15|module github.com/osmosis-labs/osmosis/v15 // changed epochs |g' go.mod
fi

go mod tidy