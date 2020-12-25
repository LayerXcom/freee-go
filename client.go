package freee

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/oauth2"
)

const (
	HeaderXAPIVersion     = "X-Api-Version"
	HeaderXFreeeRequestID = "X-Freee-Request-ID"

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
	Log         Logger
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
) (*oauth2.Token, error) {
	// url
	u, err := url.Parse(c.config.APIEndpoint)
	if err != nil {
		return oauth2Token, err
	}
	u.Path = path.Join(u.Path, APIPath1, apiPath)
	u.RawQuery = queryParams.Encode()
	var req *http.Request
	if method == http.MethodDelete {
		// headers
		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return oauth2Token, err
		}
	} else {
		// post payload
		jsonParams, err := json.Marshal(postBody)
		if err != nil {
			return oauth2Token, err
		}
		// headers
		req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(jsonParams))
		if err != nil {
			return oauth2Token, err
		}
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set(HeaderXAPIVersion, XAPIVersion20200615)
	req = req.WithContext(ctx)
	tokenSource := c.config.Oauth2.TokenSource(ctx, oauth2Token)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	response, err := httpClient.Do(req)
	if err != nil {
		return oauth2Token, err
	}
	defer response.Body.Close()
	c.logf("[freee] %s: %s", HeaderXFreeeRequestID, response.Header.Get(HeaderXFreeeRequestID))
	c.logf("[freee] %s: %v %v%v", response.Status, req.Method, req.URL.Host, req.URL.Path)

	oauth2Token, err = tokenSource.Token()
	if err != nil {
		// error occured, but ignored.
		c.logf("[freee] OAuth2: %v", err)
	}

	var r io.Reader = response.Body
	// r = io.TeeReader(r, os.Stderr)

	// Parse freee API errors
	code := response.StatusCode
	if code >= http.StatusBadRequest {
		byt, err := ioutil.ReadAll(r)
		if err != nil {
			// error occured, but ignored.
			c.logf("[freee] HTTP response body: %v", err)
		}
		res := &Error{
			StatusCode: code,
			RawError:   string(byt),
		}
		// Check if re-authorization is required
		if code == http.StatusUnauthorized {
			var e UnauthorizedError
			if err := json.NewDecoder(bytes.NewReader(byt)).Decode(&e); err != nil {
				c.logf("[freee] HTTP response body: %v", err)
				return oauth2Token, res
			}
			if e.Code == UnauthorizedCodeInvalidAccessToken ||
				e.Code == UnauthorizedCodeExpiredAccessToken {
				res.IsAuthorizationRequired = true
			}
		}
		return oauth2Token, res
	}

	if res == nil {
		return oauth2Token, nil
	}

	return oauth2Token, json.NewDecoder(r).Decode(&res)
}
