#!/bin/sh
set -e

. ./request.sh

URL="https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc/uci?auth=${OPENWRT_TOKEN}"

request "{ \"method\": \"get_all\", \"params\": [ \"network\" ] }"
