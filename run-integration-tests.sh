#!/bin/sh
alias dc="docker-compose -f ./docker-compose-test.yml"
sn=test
dc up -d
dc logs -f $sn
res=$(docker inspect --format "{{.State.ExitCode}}" "$(dc ps -q $sn)")
exit $res