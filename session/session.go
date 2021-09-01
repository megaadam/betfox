package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/megaadam/betfox/betfair/client"
	"github.com/megaadam/betfox/betfair/client/operations"
	"github.com/megaadam/betfox/betfair/models"
)

type creds struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	AppKey string `json:"appKey"`
}

// Login --stuff
func Login() {
	file, _ := ioutil.ReadFile("/home/a/.betfair/creds.json")
	creds := creds{}
	_ = json.Unmarshal([]byte(file), &creds)

	reqURL, _ := url.Parse("https://identitysso-cert.betfair.se:443/api/certlogin")

	// create request body
	reqBody := ioutil.NopCloser(strings.NewReader("username=" + creds.User + "&password=" + creds.Pass))
	// 	{
	// 		"username":creds.user,
	// 		"pasword":creds.password,
	// 	}
	// `))

	// create a request object
	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/x-www-form-urlencoded"},
			"X-Application": {"jvisaiEIJz06FUlF"},
		},
		Body: reqBody,
	}

	res, err := http.DefaultClient.Do(req)

	fmt.Println(res, err)
	// read response body
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", data)

	cm := models.ConnectionMessage{}
	cm.ConnectionID = ""

	params := operations.PostRequestParams{}

	auth := new(models.AuthenticationMessage)
	auth.AppKey = creds.AppKey
	auth.Session = "SESSIONTOKEN"

	params.RequestMessage = new(models.AllRequestTypesExample)
	params.RequestMessage.Authentication = auth
	params.RequestMessage.Heartbeat = &models.HeartbeatMessage{}

	cli := client.Default.Operations.PostRequest(&params)
	fmt.Println(cli)
}
