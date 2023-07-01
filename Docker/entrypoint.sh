#!/bin/bash
set -ex

if [ "${1:0:1}" = '-' ] || [ "${1:0:1}" = '' ]; then
  set -- go run  ./ws_server.go  "$@"
fi

exec "$@"

