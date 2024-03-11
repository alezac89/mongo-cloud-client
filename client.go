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
	GroupId    string
	AppId      string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	ApiKey   string `json:"apiKey"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	UserId      string `json:"user_id"`
	DeviceId    string `json:"device_id"`
}

// NewClient -
func NewClient(username, apiKey, groupId, appId *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Mongo Cloud URL
		HostURL: HostURL,
	}

	// If username or password not provided, return empty client
	if username == nil || apiKey == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		ApiKey:   *apiKey,
	}

	c.GroupId = *groupId
	c.AppId = *appId

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
