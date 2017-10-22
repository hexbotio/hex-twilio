package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
)

type HexTwilio struct {
}

func (g *HexTwilio) Perform(args hexplugin.Arguments) (resp hexplugin.Response) {

	// initialize return values
	var output = ""
	var success = true

	client := &http.Client{}
	accountSid := args.Config["account_sid"]
	authToken := args.Config["auth_token"]
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	values := url.Values{}
	values.Set("To", args.Config["send_to"])
	values.Set("From", args.Config["send_from"])
	values.Set("Body", args.Command)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		output = "ERROR - Twilio New Request " + err.Error()
		success = false
	}
	_, err = client.Do(req)
	if err != nil {
		output = "ERROR - Twilio API Error " + err.Error()
		success = false
	}

	if success {
		output = "Message Sent"
	}
	resp = hexplugin.Response{
		Output:  output,
		Success: success,
	}
	return resp
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"action": &hexplugin.HexPlugin{Impl: &HexTwilio{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hexplugin.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}
