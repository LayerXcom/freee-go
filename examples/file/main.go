package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/LayerXcom/freee-go"
	"golang.org/x/oauth2"
)

func main() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")
	conf := freee.NewConfig(clientID, clientSecret, redirectURL)
	conf.Log = log.New(os.Stdout, "", log.LstdFlags)
	client := freee.NewClient(conf)

	ctx := context.Background()
	token := &oauth2.Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}

	companyID, err := strconv.Atoi(os.Getenv("COMPANY_ID"))
	if err != nil {
		log.Fatal(err)
	}
	file := os.Getenv("FILE")
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	params := freee.CreateReceiptParams{
		CompanyID: int32(companyID),
		Receipt:   f,
	}
	resp, token, err := client.CreateReceipt(ctx, token, params, filepath.Base(file))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp)
	fmt.Printf("%#v\n", token)
}
