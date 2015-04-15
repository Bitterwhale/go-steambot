package main

import (
	"io/ioutil"
	"log"

	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/internal/steamlang"
	"github.com/Philipp15b/go-steam/tradeoffer"
)

func main() {
	myLoginInfo := new(steam.LogOnDetails)
	myLoginInfo.Username = ""
	myLoginInfo.Password = ""
	myLoginInfo.AuthCode = "3XYXX"
	hash, _ := ioutil.ReadFile("sentry")
	myLoginInfo.SentryFileHash = steam.SentryHash(hash)
	client := steam.NewClient()
	client.Connect()

	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			log.Print("Connected to Steam")
			client.Auth.LogOn(myLoginInfo)
		case *steam.MachineAuthUpdateEvent:
			log.Print("machineauth")
			ioutil.WriteFile("sentry", e.Hash, 0666)
		case *steam.LoggedOnEvent:
			log.Print("Logged on")
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
			client.Web.LogOn()
		case *steam.WebLoggedOnEvent:
			//, _ := steamid.NewId("STEAM_0:0:24609886")
			tradingClient := tradeoffer.NewClient("0C4EF9B2D79AAA9FEF5CB3166FFC8D58", client.Web.SessionId, client.Web.SteamLogin, client.Web.SteamLoginSecure)
			inventory, _ := tradingClient.GetOwnInventory(2, 730)

			for _, item := range inventory.Descriptions {
				log.Print(item.MarketName, item.Type)
			}
		case steam.FatalErrorEvent:
			log.Print(e)
		case error:
			log.Print(e)
		}
	}
}
