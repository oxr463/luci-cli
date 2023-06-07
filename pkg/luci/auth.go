package luci

func GetToken() string {
	username := GetUsername()
	password := GetPassword()

        _ = username
        _ = password

	url := "https://${OPENWRT_HOSTNAME}/cgi-bin/luci/rpc/auth"

        _ = url

	data := RequestData{
          Id: 1,
          Method: "login",
          Params: []string{"fixme"},
        }

        _ = data

	// "{ \"id\": 1, \"method\": \"login\", \"params\": [ \"${OPENWRT_USERNAME}\", \"${OPENWRT_PASSWORD}\" ] }"
        token := "fixme"

	// TODO: parse JSON result for token
	return token
}
