#!/bin/bash
set -ex

if [ "${1:0:1}" = '-' ] || [ "${1:0:1}" = '' ]; then
  set -- ./web-socket-show  "$@"
fi

exec "$@"

