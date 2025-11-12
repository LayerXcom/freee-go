package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/LayerXcom/freee-go"
	"golang.org/x/oauth2"
)

func main() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
	conf := freee.NewConfig(clientID, clientSecret, redirectURL)
	conf.Log = log.New(os.Stdout, "", log.LstdFlags)
	client := freee.NewClient(
		conf, freee.WithHTTPRoundTripper(
			&http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		),
	)

	ctx := context.Background()
	token := &oauth2.Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}
	me, token, err := client.GetUsersMe(ctx, token, freee.GetUsersMeOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", token)
}
