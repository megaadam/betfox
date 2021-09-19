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
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/megaadam/betfox/betfair/client"
	"github.com/megaadam/betfox/betfair/client/operations"
	"github.com/megaadam/betfox/betfair/models"
)

const ssoURL = "https://identitysso-cert.betfair.se:443/api/certlogin"
const streamURL = "stream-api.betfair.com:443"
const streamIntegrationURL = "stream-api-integration.betfair.com:443"

type creds struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	AppKey string `json:"appKey"`
}

func opFunc(opts *runtime.ClientOperation) {
	opts.Schemes = []string{"https"}
	// opts.PathPattern = ""
}

// GetNyarumClient --
func GetNyarumClient(appKey string) *betting.Betfair {
	client := betting.NewBetfair(appKey)
	config := loadConfig()

	err := client.GetSession(config.CertPem, config.CertKey, config.Login, config.Password)
	if err != nil {
		fmt.Println(err)
	}
	return client
}

func newTrue() *bool {
	b := true
	return &b
}

// Login --
func Login() {
	// creds
	file, _ := ioutil.ReadFile("/home/a/.betfair/creds.json")
	creds := creds{}
	_ = json.Unmarshal([]byte(file), &creds)

	nyarumClient := GetNyarumClient(creds.AppKey)

	details, err := nyarumClient.GetAccountDetails()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(details)

	filter := betting.Filter{Wallet: betting.W_UK}
	funds, err := nyarumClient.GetAccountFunds(filter)

	fmt.Println(funds)
	mf := betting.MarketFilter{InPlayOnly: newTrue(),
		EventTypeIDs:    []string{"1"},
		MarketTypeCodes: []string{"MATCH_ODDS"},
	}

	f2 := betting.Filter{MarketFilter: &mf,
		MaxResults:   20,
		FromCurrency: "SEK",
	}
	mt, err := nyarumClient.ListMarketTypes(f2)

	_ = mt
	// events, err := nyarumClient.ListEvents(f2)
	// fmt.Println(events[0])

	mp := []betting.EMarketProjection{"COMPETITION",
		"EVENT",
		"EVENT_TYPE",
		//		"EVENT_TYPE",
		"MARKET_DESCRIPTION",
		"RUNNER_DESCRIPTION",
		//		"RUNNER_METADATA",
	}

	f2.MarketProjection = &mp
	f2.Sort = "LAST_TO_START"
	markets, err := nyarumClient.ListMarketCatalogue(f2)

	for _, market := range markets {
		_ = market
	}

	return

	///////////////////////////////////////
	config := loadConfig()

	cfg := client.DefaultTransportConfig()
	cfg.Host = streamIntegrationURL
	cli := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	tls := new(httptransport.TLSClientOptions)
	tls.Certificate = config.CertPem
	tls.Key = config.CertKey
	tls.InsecureSkipVerify = true

	tlst, err := httptransport.TLSTransport(*tls)
	fmt.Println(tlst)

	prp := operations.NewPostRequestParams()
	rm := new(models.AllRequestTypesExample)
	rm.Authentication = new(models.AuthenticationMessage)
	rm.Authentication.AppKey = creds.AppKey
	rm.Authentication.Session = nyarumClient.SessionKey

	prp.RequestMessage = rm
	eee := cli.Operations.PostRequest(prp, opFunc)

	m := models.AuthenticationMessage{}
	m.AppKey = creds.AppKey
	m.Session = nyarumClient.SessionKey

	//res, err = models.AllRequestTypesExample

	fmt.Println(m, eee)
}

func testSwag() {
	flt := models.MarketFilter{}
	_ = flt
	return
}

func classicLogin() {
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
	clx := http.Client{
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

	res, err := clx.Do(req)

	fmt.Println(res, err)
	// read response body
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", data)

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
