package freee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

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

type ErrorResponse struct {
	StatusCode string  `json:status_code`
	Errors     []Error `json:errors`
}

func (e *ErrorResponse) Error() string {
	ems := make([]string, len(e.Errors))
	for i, e := range e.Errors {
		m := strings.Join(e.Messages, ",")
		em := fmt.Sprintf("type: %s, messsage: %s", e.Type, m)
		ems[i] = em
	}
	return strings.Join(ems, ",")
}

type Error struct {
	Type     string   `json:type`
	Messages []string `json:messages`
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
	// post payload
	jsonParams, err := json.Marshal(postBody)
	if err != nil {
		return tokenSource, err
	}
	// headers
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(jsonParams))
	if err != nil {
		return tokenSource, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HeaderXAPIVersion, XAPIVersion20200615)
	req = req.WithContext(ctx)

	httpClient := oauth2.NewClient(ctx, tokenSource)
	response, err := httpClient.Do(req)
	if err != nil {
		return tokenSource, err
	}
	defer response.Body.Close()

	var r io.Reader = response.Body
	code := response.StatusCode
	if code >= http.StatusBadRequest {
		var e ErrorResponse
		json.NewDecoder(r).Decode(&e)
		return tokenSource, &e
	}
	// r = io.TeeReader(r, os.Stderr)
	if res == nil {
		return tokenSource, nil
	}

	return tokenSource, json.NewDecoder(r).Decode(&res)
}
