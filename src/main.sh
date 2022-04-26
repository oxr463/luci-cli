#!/bin/sh
set -e

LUCI_RPC_URL="https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc"

_request() {
  request_url="${LUCI_RPC_URL}/${1}"
  request_data="${2}"

  curl "${request_url}" --data "${request_data}"
}

luci_rpc_auth() {
  openwrt_username="${1}"
  openwrt_password="${2}"

  _request "auth" "{ \"id\": 1, \"method\": \"login\", \"params\": [ \"${openwrt_username}\", \"${openwrt_password}\" ] }"
}

luci_rpc_sys_net_conntrack() {
  _request "sys?auth=${OPENWRT_TOKEN}" "{ \"method\": \"net.conntrack\" }"
}

luci_rpc_sys_net_devices() {
  _request "sys?auth=${OPENWRT_TOKEN}" "{ \"method\": \"net.devices\" }"
}

luci_rpc_sys_uptime() {
  _request "sys?auth=${OPENWRT_TOKEN}" "{ \"method\": \"uptime\" }"
}

luci_rpc_uci_get_all_network() {
  _request "sys?auth=${OPENWRT_TOKEN}" "{ \"method\": \"get_all\", \"params\": [ \"network\" ] }"
}

