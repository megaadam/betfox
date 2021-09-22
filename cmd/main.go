package main

import (
	"os"

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

	_, err = cli.Markets()

}
