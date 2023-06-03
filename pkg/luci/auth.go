package luci

import (
	"luci/environment"
	"luci/request"
)

func GetToken() {
	username := GetUsername
	password := GetPassword
	
  url := "https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc/auth"
  data := RequestData
  "{ \"id\": 1, \"method\": \"login\", \"params\": [ \"${OPENWRT_USERNAME}\", \"${OPENWRT_PASSWORD}\" ] }"

  // TODO: parse JSON result for token
  return token
}
