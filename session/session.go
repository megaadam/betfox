package session

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Nyarum/betting"
)

type creds struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	AppKey string `json:"appKey"`
}

// Login --stuff
func Login() {
	// creds
	file, _ := ioutil.ReadFile("/home/a/.betfair/creds.json")
	creds := creds{}
	_ = json.Unmarshal([]byte(file), &creds)

	// cert
	// the CertPool wants to add a root as a []byte so we read the file ourselves
	caCert, err := ioutil.ReadFile("/path/to/ca.crt")
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)
	// LoadX509KeyPair reads files, so we give it the paths
	clientCert, err := tls.LoadX509KeyPair("/home/a/.betfair/client-2048.crt", "/home/a/.betfair/client-2048.key")
	tlsConfig := tls.Config{
		RootCAs:            pool,
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := http.Client{
		Transport: &transport,
	}

	reqURL, _ := url.Parse("https://identitysso-cert.betfair.se:443/api/certlogin")

	// create request body
	reqBody := ioutil.NopCloser(strings.NewReader("username=" + creds.User + "&password=" + creds.Pass))

	// create a request object
	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
			//			"X-Application": {"jvisaiEIJz06FUlF"},
			"X-Application": {"dummy"},
		},
		Body: reqBody,
	}

	res, err := client.Do(req)

	fmt.Println(res, err)
	// read response body
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", data)

	bet := betting.NewBet(creds.AppKey)
	config := loadConfig()

	err = bet.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		fmt.Println(err)
	}

	details, err := bet.GetAccountDetails()
	fmt.Println(details)

	filter := betting.Filter{}
	filter.Wallet = betting.W_UK
	funds, err := bet.GetAccountFunds(filter)

	fmt.Println(funds)
}

type Test struct {
	ApiKey   string `json:"api_key"`
	Login    string `json:"login"`
	Password string `json:"password"`
	CertPem  string `json:"cert_pem"`
	CertKey  string `json:"cert_key"`
	Debug    bool   `json:"debug"`
}

func loadConfig() (test Test) {
	loadFile, err := ioutil.ReadFile("/home/a/.betfair/nyarum.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(loadFile, &test)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
