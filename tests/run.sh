#!/bin/bash
IFS=$'\n'

scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
conf="${scriptdir}/conf.txt"

for el in $(cat "${conf}" | grep -Pv '^#'); do
    in=$(echo "${el}" | grep -Po ".*(?=;)")
    exp=$(echo "${el}" | grep -Po "[^;]+$")

    eval go run src/*.go "${in}"
    eval go run src/*.go "${in}" >/dev/stdout 2>&1 | grep "${exp}" &>/dev/null
    if [[ $? != 0 ]]; then
        echo -e "\\033[0;91mfail \\033[0m${in}"
        echo "--------------------------------"
    fi
done
