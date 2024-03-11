package mongoCloud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Mongo Cloud URL
const HostURL string = "https://services.cloud.mongodb.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken   string `json:"access_token"`
	UserId string `json:"user_id"`
	DeviceId    string `json:"device_id"`
}

// NewClient -
func NewClient(username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Mongo Cloud URL
		HostURL: HostURL,
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}
	

	ar, err := c.GetUserTokenSignIn()
	if err != nil {
		return nil, err
	}

	c.Token = fmt.Sprintf("Bearer %s", ar.AccessToken)

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}