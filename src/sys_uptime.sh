#!/bin/sh
set -e

. ./request.sh

URL="https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc/sys?auth=${OPENWRT_TOKEN}"

request "{ \"method\": \"uptime\" }"
