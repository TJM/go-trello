package trello

import (
	"encoding/json"
	"net/url"
)

// Label - Label Type
type Label struct {
	client *Client

	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}

// UpdateName - Update Name on a Label (Update a Label)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-put
func (l *Label) UpdateName(name string) (err error) {
	payload := url.Values{}
	payload.Set("value", name)

	body, err := l.client.Put("/labels/"+l.ID+"/name", payload)
	if err != nil {
		return
	}
	return json.Unmarshal(body, l)
}

// UpdateColor - Update Color for Label (Update a Label)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-put
// Color can be null
func (l *Label) UpdateColor(color string) (err error) {
	payload := url.Values{}
	payload.Set("value", color)

	body, err := l.client.Put("/labels/"+l.ID+"/color", payload)
	if err != nil {
		return
	}
	return json.Unmarshal(body, l)
}

// DeleteLabel - Delete a Label
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-delete
func (l *Label) DeleteLabel() error {
	_, err := l.client.Delete("/labels/" + l.ID)
	return err
}
