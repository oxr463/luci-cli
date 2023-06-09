package luci

import (
	"errors"
	"os"
)

func GetHostname() (string, error) {
	hostname, present := os.LookupEnv("OPENWRT_HOSTNAME")

	if !present {
		return "", errors.New("OPENWRT_HOSTNAME not found")
	} else {
		return hostname, nil
	}
}

func GetUsername() (string, error) {
	username, present := os.LookupEnv("OPENWRT_USERNAME")

	if !present {
		return "", errors.New("OPENWRT_USERNAME not found")
	} else {
		return username, nil
	}
}

func GetPassword() (string, error) {
	password, present := os.LookupEnv("OPENWRT_PASSWORD")

	if !present {
		return "", errors.New("OPENWRT_PASSWORD not found")
	} else {
		return password, nil
	}
}
