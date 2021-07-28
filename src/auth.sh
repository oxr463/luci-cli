#!/bin/sh
set -e

. ./request.sh

URL="https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc/auth"

request "{ \"id\": 1, \"method\": \"login\", \"params\": [ \"${OPENWRT_USERNAME}\", \"${OPENWRT_PASSWORD}\" ] }"
