package session

import (
	"encoding/json"
	"io/ioutil"
)

type creds struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

// Fundamental Login
func Login() {
	file, _ := ioutil.ReadFile("/home/a/.betfair/creds.json")
	creds := creds{}
	_ = json.Unmarshal([]byte(file), &creds)

}
