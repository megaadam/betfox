package main

import (
	"os"

	"fmt"
	// "github.com/megaadam/betfox/lab"
	"github.com/megaadam/betfox/session"
)

func main() {
	cli := (*session.NyarumClient)(session.Login())
	_, err := cli.Details()
	if err != nil {
		os.Exit(3)
	}

	_, err = cli.Funds()

	markets, err := cli.Markets()
	marketID := markets[0].MarketID
	fmt.Println(marketID)
	eventID := markets[0].Event.ID
	session.Stream(cli.Client.ApiKey, cli.Client.SessionKey, marketID, eventID)
	fmt.Println(markets[0].Event.Name, "\t", markets[0].Competition.Name)
}
