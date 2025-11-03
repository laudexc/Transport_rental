package main

import (
	"Transport_rental/internal/common"
	"Transport_rental/internal/genid"
)

func main() {
	genid.NewID("user", genid.UuidGen{})   // create user ID
	genid.NewID("car", genid.ShortIDGen{}) // create transport ID

	genid.ShowAllID()    // show all IDs
	common.ShowLogsErr() // show all logs with errors
}
