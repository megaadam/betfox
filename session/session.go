package session

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	//httptransport "github.com/go-openapi/runtime/client"

	"github.com/Nyarum/betting"
	"github.com/go-openapi/runtime"
	"github.com/megaadam/betfox/betfair/models"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const ssoURL = "https://identitysso-cert.betfair.se:443/api/certlogin"
const streamURL = "stream-api.betfair.com:443"
const streamIntegrationURL = "stream-api-integration.betfair.com:443"

type creds struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	AppKey string `json:"appKey"`
}

// '{\\\"authentication\\\":{\\\"appKey\\\":\\\"jvisaiEIJz06FUlF\\\",\\\
//"session\\\":\\\"HBD7dsV9sSOj3cS/T1r9JFuF0Q7pceltvZwQGJBabVs=\\\",\\\
// "op\\\":\\\"AuthenticationMessage\\\"},\\\""

type streamAuth struct {
	AppKey  string `json:"appKey"`
	Session string `json:"session"`
	Op      string `json:"op"`
}

// NyarumClient -- Wrapper type to expand type betting.Betfair
type NyarumClient betting.Betfair

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

// Details --
func (cli *NyarumClient) Details() (betting.AccountDetails, error) {

	details, err := cli.GetAccountDetails()
	perr(err)

	fmt.Println(details)
	return details, err
}

// Funds --
func (cli *NyarumClient) Funds() (betting.AccountFunds, error) {
	filter := betting.Filter{Wallet: betting.W_UK}
	funds, err := cli.GetAccountFunds(filter)
	perr(err)

	fmt.Println(funds)
	return funds, err
}

// Markets --
func (cli *NyarumClient) Markets() ([]betting.MarketCatalogue, error) {
	mf := betting.MarketFilter{InPlayOnly: newTrue(),
		EventTypeIDs:    []string{"1"},
		MarketTypeCodes: []string{"MATCH_ODDS"},
	}

	f2 := betting.Filter{MarketFilter: &mf,
		MaxResults:   20,
		FromCurrency: "SEK",
	}
	mt, err := cli.ListMarketTypes(f2)

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
	f2.Sort = "MAXIMUM_TRADED"
	markets, err := cli.ListMarketCatalogue(f2)

	marketIDs := cli.getMarketIDs(markets)
	books, err := cli.MarketBooks(marketIDs)
	for _, book := range books {
		p := message.NewPrinter(language.German)
		p.Printf("%d\n", (int)(book.TotalMatched))
	}

	for _, market := range markets {
		mb, err := cli.MarketBook(market.MarketID)
		perr(err)
		p := message.NewPrinter(language.German)

		delay := mb.BetDelay
		p.Printf("%d \t%d\t", (int)(mb.TotalMatched), delay)

		//fmt.Printf("%d \t", (int)(market.TotalMatched))
		fmt.Println(market.Event.Name, "\t", market.Competition.Name)

	}

	return markets, err
} // Markets()

// Stream --
func Stream(apiKey, sessionKey, market, eventID string) {
	log.SetOutput(os.Stdout)
	config := loadConfig()

	cert, err := tls.LoadX509KeyPair(config.CertPem, config.CertKey)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	tlscfg := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", streamIntegrationURL, &tlscfg)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	{
		reply := make([]byte, 256)
		n, err := conn.Read(reply)
		log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
		fmt.Printf("client: read %q (%d bytes)\n", string(reply[:n]), n)
		rep := string(reply[:n])
		fmt.Println(rep)
		_ = err
		_ = rep
	}
	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		fmt.Println(v.Subject)
	}
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	rm := new(models.AllRequestTypesExample)
	am := models.AuthenticationMessage{}
	am.AppKey = apiKey
	am.Session = sessionKey
	marsh, err := am.MarshalJSON()
	mx := string(marsh) + "\r\n"

	_ = mx
	rm.Authentication = &am
	rm.OpTypes = "authentication"

	msg, err := json.Marshal(rm)
	msgstr := string(msg)

	m2 := new(streamAuth)
	m2.AppKey = apiKey
	m2.Session = sessionKey
	m2.Op = "authentication"

	msg2, err := json.Marshal(m2)
	msg2str := string(msg2) + "\r\n"
	_ = msg2str
	n, err := io.WriteString(conn, msg2str)
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", msgstr, n)

	reply := make([]byte, 1256)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	fmt.Printf("client: read %q (%d bytes)\n", string(reply[:n]), n)
	rep := string(reply[:n])
	_ = rep

	example := fmt.Sprintf("{\"op\":\"marketSubscription\",\"id\":2,\"marketFilter\":{\"marketIds\":[\"%s\"],\"bspMarket\":true,\"bettingTypes\":"+
		"[\"ODDS\"],\"eventTypeIds\":[\"1\"],\"eventIds\":[\"%s\"],\"turnInPlayEnabled\":true,\"marketTypes\":[\"MATCH_ODDS\"]}"+
		",\"marketDataFilter\":{}}", market, eventID)

	sockWrite(conn, example)

	for i := 1; i < 20; i++ {
		sockRead(conn)
		time.Sleep(2 * time.Second)
	}
	log.Print("client: exiting")
}

func sockWrite(conn *tls.Conn, m string) {
	n, err := io.WriteString(conn, m+"\r\n")
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", m, n)

	reply := make([]byte, 1256)
	n, err = conn.Read(reply)
	rep := string(reply[:n])
	log.Printf("client: read %q (%d bytes)", rep, n)
	fmt.Printf("client: read %q (%d bytes)\n", rep, n)
}

func sockRead(conn *tls.Conn) {
	reply := make([]byte, 1256)
	n, err := conn.Read(reply)
	_ = err
	rep := string(reply[:n])
	log.Printf("client: read %q (%d bytes)", rep, n)
	fmt.Printf("client: read %q (%d bytes)\n", rep, n)
}

// MarketBook --
func (cli *NyarumClient) MarketBook(marketID string) (betting.MarketBook, error) {
	filter := betting.Filter{MarketIDs: []string{marketID}}
	mb, err := cli.ListMarketBook(filter)
	return mb[0], err
}

// MarketBooks --
func (cli *NyarumClient) MarketBooks(marketIDs []string) ([]betting.MarketBook, error) {
	filter := betting.Filter{MarketIDs: marketIDs,
		Sort:         "MAXIMUM_TRADED",
		MaxResults:   20,
		FromCurrency: "SEK"}
	mb, err := cli.ListMarketBook(filter)
	return mb, err
}

// Login --
func Login() *betting.Betfair {
	// creds
	file, _ := ioutil.ReadFile("/home/a/.betfair/creds.json")
	creds := creds{}
	_ = json.Unmarshal([]byte(file), &creds)

	nyarumClient := GetNyarumClient(creds.AppKey)
	return nyarumClient
}

func (cli *NyarumClient) getMarketIDs(markets []betting.MarketCatalogue) []string {
	var IDs []string

	for _, market := range markets {
		IDs = append(IDs, market.MarketID)
	}
	return IDs
}

func testSwag() {
	flt := models.MarketFilter{}
	_ = flt
	return
}

func perr(err error) {
	if err != nil {
		fmt.Println(err)
	}
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
