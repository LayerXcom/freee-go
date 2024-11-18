package freee

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"time"

	"golang.org/x/oauth2"
)

const (
	HeaderXAPIVersion     = "X-Api-Version"
	HeaderXFreeeRequestID = "X-Freee-Request-ID"

	APIEndpoint         = "https://api.freee.co.jp"
	APIPath1            = "/api/1"
	XAPIVersion20200615 = "2020-06-15"
)

// Config is a setting for freee APIs.
type Config struct {
	APIEndpoint string
	Log         Logger
	Oauth2      *oauth2.Config
	EarlyExpiry *time.Duration
}

func NewConfig(clientID, clientSecret, redirectURL string) *Config {
	return &Config{
		APIEndpoint: APIEndpoint,
		Oauth2: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:   Oauth2AuthURL,
				TokenURL:  Oauth2TokenURL,
				AuthStyle: oauth2.AuthStyleInParams,
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
	var (
		contentType string
		body        io.Reader
	)
	if method != http.MethodDelete {
		contentType = "application/json"
		jsonParams, err := json.Marshal(postBody)
		if err != nil {
			return oauth2Token, err
		}
		body = bytes.NewBuffer(jsonParams)
	}

	req, err := c.newRequest(ctx, apiPath, method, contentType, queryParams, body)
	if err != nil {
		return oauth2Token, err
	}
	return c.do(ctx, oauth2Token, req, res)
}

func (c *Client) postFiles(ctx context.Context,
	apiPath string, method string,
	oauth2Token *oauth2.Token,
	queryParams url.Values, postBody map[string]string,
	fileName string, file []byte,
	res interface{},
) (*oauth2.Token, error) {
	var (
		contentType string
		body        = &bytes.Buffer{}
	)
	mw := multipart.NewWriter(body)
	fw, err := mw.CreateFormFile("receipt", fileName)
	if err != nil {
		return oauth2Token, err
	}
	if _, err := io.Copy(fw, bytes.NewReader(file)); err != nil {
		return oauth2Token, err
	}
	for k, v := range postBody {
		if err := mw.WriteField(k, v); err != nil {
			return oauth2Token, err
		}
	}
	contentType = mw.FormDataContentType()
	if err := mw.Close(); err != nil {
		return oauth2Token, err
	}

	req, err := c.newRequest(ctx, apiPath, method, contentType, queryParams, body)
	if err != nil {
		return oauth2Token, err
	}
	return c.do(ctx, oauth2Token, req, res)
}

func (c *Client) newRequest(
	ctx context.Context,
	apiPath string, method string,
	contentType string,
	queryParams url.Values,
	body io.Reader,
) (*http.Request, error) {
	// construct url
	u, err := url.Parse(c.config.APIEndpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, APIPath1, apiPath)
	u.RawQuery = queryParams.Encode()
	// request with context
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	// set http headers
	req.Header.Set(HeaderXAPIVersion, XAPIVersion20200615)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return req, nil
}

func (c *Client) do(
	ctx context.Context,
	oauth2Token *oauth2.Token,
	req *http.Request,
	res interface{},
) (*oauth2.Token, error) {
	tokenSource := c.config.Oauth2.TokenSource(ctx, oauth2Token)
	if c.config.EarlyExpiry != nil {
		tokenSource = oauth2.ReuseTokenSourceWithExpiry(oauth2Token, tokenSource, *c.config.EarlyExpiry)
	}
	httpClient := oauth2.NewClient(ctx, tokenSource)
	response, err := httpClient.Do(req)
	if err != nil {
		e := &oauth2.RetrieveError{}
		if errors.As(err, &e) {
			resp := &Error{
				RawError:                err.Error(),
				IsAuthorizationRequired: true,
			}
			if e.Response != nil {
				resp.StatusCode = e.Response.StatusCode
			}
			return oauth2Token, resp
		}
		errURL := &url.Error{}
		if errors.As(err, &errURL) {
			err = errURL.Unwrap()
			if v, ok := err.(*Error); ok {
				err = v
			}
		}
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
