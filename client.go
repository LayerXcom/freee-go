package freee

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/oauth2"
)

const (
	HeaderXAPIVersion = "X-Api-Version"

	APIEndpoint         = "https://api.freee.co.jp"
	APIPath1            = "/api/1"
	XAPIVersion20200615 = "2020-06-15"

	Oauth2TokenURL = "https://accounts.secure.freee.co.jp/public_api/token"
	Oauth2AuthURL  = "https://accounts.secure.freee.co.jp/public_api/authorize"
)

var (
	oauth2Endopint = oauth2.Endpoint{
		AuthURL:  Oauth2AuthURL,
		TokenURL: Oauth2TokenURL,
	}
)

// Config is a setting for freee APIs.
type Config struct {
	APIEndpoint string
	Oauth2      *oauth2.Config
}

func NewConfig(clientID, clientSecret, redirectURL string) *Config {
	return &Config{
		APIEndpoint: APIEndpoint,
		Oauth2: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  Oauth2AuthURL,
				TokenURL: Oauth2TokenURL,
			},
		},
	}
}

// Client represents an API client for freee.
type Client struct {
	httpClient *http.Client
	config     *Config
}

// NewClient returns a new freee API client.
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) call(ctx context.Context,
	apiPath string, method string,
	oauth2Token *oauth2.Token,
	queryParams url.Values, postBody interface{},
	res interface{},
) (oauth2.TokenSource, error) {
	tokenSource := c.config.Oauth2.TokenSource(ctx, oauth2Token)

	// url
	u, err := url.Parse(c.config.APIEndpoint)
	if err != nil {
		return tokenSource, err
	}
	u.Path = path.Join(u.Path, APIPath1, apiPath)
	u.RawQuery = queryParams.Encode()
	var req *http.Request
	if method == http.MethodDelete {
		// headers
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return tokenSource, err
		}
	} else {
		// post payload
		jsonParams, err := json.Marshal(postBody)
		if err != nil {
			return tokenSource, err
		}
		// headers
		req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(jsonParams))
		if err != nil {
			return tokenSource, err
		}
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set(HeaderXAPIVersion, XAPIVersion20200615)
	req = req.WithContext(ctx)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	response, err := httpClient.Do(req)
	if err != nil {
		return tokenSource, err
	}
	defer response.Body.Close()

	var r io.Reader = response.Body
	// r = io.TeeReader(r, os.Stderr)
	code := response.StatusCode
	if code >= http.StatusBadRequest {
		byt, err := ioutil.ReadAll(r)
		if err != nil {
			return tokenSource, err
		}
		return tokenSource, errors.New(string(byt))
	}
	if res == nil {
		return tokenSource, nil
	}

	return tokenSource, json.NewDecoder(r).Decode(&res)
}
