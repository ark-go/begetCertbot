package internal

import (
	"fmt"
	"net/url"
)

func makeUrl(commandUrl string, param string) (string, error) {
	baseUrl, err := url.Parse(commandUrl)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return "", err
	}
	// Prepare Query Parameters
	params := url.Values{}
	params.Add("login", Pass.login)
	params.Add("passwd", Pass.password)
	params.Add("input_format", "json")
	params.Add("output_format", "json")
	return baseUrl.String(), nil
}
