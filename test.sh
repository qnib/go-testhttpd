#!/bin/bash
set -xe

govendor fetch +missing
echo "> govendor remove +unused"
govendor remove +unused
echo "> govendor update +local"
govendor update +local
echo "> govendor sync +external"
govendor sync +external
if [ ! -d resources/coverity ];then
    mkdir -p resources/coverity
fi
go test -cover -coverprofile=main.cover >>/dev/null
COVER_FILES="main.cover"
for x in $(find . -maxdepth 1 -type d |egrep -v "(\.$|\.git|vendor|bin|resources|deploy)");do
    go test -cover -coverprofile=resources/coverity/${x}.cover ${x} >>/dev/null
    COVER_FILES="${COVER_FILES} resources/coverity/${x}.cover"
done
coveraggregator -o coverage-all.out ${COVER_FILES} >>/dev/null
#go tool cover -func=coverage-all.outcover |tee ./resources/coverity/coverage-all.out
#go tool cover -html=coverage-all.out -o resources/coverity/all.html
