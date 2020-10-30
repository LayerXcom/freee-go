package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LayerXcom/freee-go"
	"golang.org/x/oauth2"
)

func main() {
	clientID := "CLIENT_ID"
	clientSecret := "CLIENT_SECRET"
	redirectURL := "REDIRECT_URL"
	conf := freee.NewConfig(clientID, clientSecret, redirectURL)

	token := &oauth2.Token{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
	}
	client := freee.NewClient(conf)
	ctx := context.Background()
	me, token, err := client.GetUsersMe(ctx, token, freee.GetUsersMeOpts{})
	if err != nil {
		if v, ok := err.(*freee.Error); ok {
			fmt.Printf("StatusCode: %v\n", v.StatusCode)
			fmt.Printf("IsAuthorizationRequired: %v\n", v.IsAuthorizationRequired)
		}
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", token)
}
