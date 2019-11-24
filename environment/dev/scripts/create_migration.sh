#!/usr/bin/env bash
name=$1
path_to_migrate=$(which migrate)
 if [[ -x "$path_to_migrate" ]] ; then
    if [[ "$name" == '' ]] ; then
        echo "please specify migration name"
        exit 1
    fi
    echo "creating. . ." &&
    migrate create -ext sql -dir migrations -seq "$name" &&
    echo "created"
 else
    echo "no migrate installed,please, install migrate"
 fi
