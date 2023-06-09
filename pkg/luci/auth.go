package luci

import (
        _ "encoding/json"
	"fmt"
        "os"
)

func GetToken() ([]byte, error) {
	hostname, err := GetHostname()
	username, err := GetUsername()
	password, err := GetPassword()

	if err != nil {
          fmt.Println(err)
        }

	url := fmt.Sprintf("https://%s/cgi-bin/luci/rpc/auth", hostname)

	data := RequestData{
		Id:     1,
		Method: "login",
		Params: []string{username, password},
	}

	token, err := Request(url, data)

	// TODO: parse JSON result for token

        os.Setenv("OPENWRT_TOKEN", "fixme")
	return token, err
}
