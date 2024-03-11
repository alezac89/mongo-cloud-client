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

	out := make([]FunctionDTO, 0, len(functions))
    for _, f := range functions {
        out = append(out, FunctionDTO{ID: f.ID, Name: f.Name, Source: f.Source, Private: f.Private, LastModified: f.LastModified, ReadOnly: f.ReadOnly})
    }

	return functions, nil
}

// CreateFunction - Create new function
func (c *Client) CreateFunction(groupId string, appId string, f FunctionDTO) (*FunctionDTO, error) {


	incomingfunction := Function{
		ID: f.ID,
		Name: f.Name,
		Source: f.Source,
		Private: f.Private,
		LastModified: f.LastModified,
		ReadOnly: f.ReadOnly,
	}

	rb, err := json.Marshal(incomingfunction)
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

	functionDTO := FunctionDTO{ID: function.ID, 
		Name: function.Name, 
		Source: function.Source, 
		Private: function.Private, LastModified: function.LastModified, ReadOnly: function.ReadOnly}

	return &functionDTO, nil
}

// UpdateFunction - Updates a function
func (c *Client) UpdateFunction(groupId string, appId string, functionId string, functionToUpdate FunctionDTO) (*FunctionDTO, error) {

	incomingfunction := Function{
		ID: functionToUpdate.ID,
		Name: functionToUpdate.Name,
		Source: functionToUpdate.Source,
		Private: functionToUpdate.Private,
		LastModified: functionToUpdate.LastModified,
		ReadOnly: functionToUpdate.ReadOnly,
	}

	rb, err := json.Marshal(incomingfunction)
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

	functionDTO := FunctionDTO{ID: function.ID, 
		Name: function.Name, 
		Source: function.Source, 
		Private: function.Private, LastModified: function.LastModified, ReadOnly: function.ReadOnly}

	return &functionDTO, nil
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
