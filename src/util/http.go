package util

import (
	"io/ioutil"
	"net/http"
)

func DoGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ""
	}
	return string(body)
}
