package mongoCloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllFunctions - Returns all app's functions
func (c *Client) GetAllFunctions(groupId string, appId string) ([]Function, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/functions", c.HostURL, groupId, appId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	functions := []Function{}
	err = json.Unmarshal(body, &functions)
	if err != nil {
		return nil, err
	}

	return functions, nil
}

// CreateFunction - Create new function
func (c *Client) CreateFunction(groupId string, appId string, incomingFunction Function) (*Function, error) {
	rb, err := json.Marshal(incomingFunction)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/functions", c.HostURL, groupId, appId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	function := Function{}
	err = json.Unmarshal(body, &function)
	if err != nil {
		return nil, err
	}

	return &function, nil
}

// UpdateFunction - Updates a function
func (c *Client) UpdateFunction(groupId string, appId string, functionId string, functionToUpdate Function) (*Function, error) {
	rb, err := json.Marshal(functionToUpdate)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/functions/%s", c.HostURL, groupId, appId, functionId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	function := Function{}
	err = json.Unmarshal(body, &function)
	if err != nil {
		return nil, err
	}

	return &function, nil
}

// DeleteFunction - Deletes a function
func (c *Client) DeleteOrder(groupId string, appId string, functionId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/admin/v3.0/groups/%s/apps/%s/functions/%s", c.HostURL, groupId, appId, functionId), nil)
	if err != nil {
		return err
	}

	_, errReq := c.doRequest(req)
	if errReq != nil {
		return errReq
	}

	return nil
}
