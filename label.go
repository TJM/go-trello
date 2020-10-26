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

// SetName - Set Name on a Label (Update a Label)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-put
func (l *Label) SetName(name string) (err error) {
	return l.Update("name", name)
}

// SetColor - Set Color for Label (Update a Label)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-put
// Color can be null
func (l *Label) SetColor(color string) (err error) {
	return l.Update("color", color)
}

// Update - Update a Label (path and value, see API docs for details)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-put
func (l *Label) Update(path, value string) (err error) {
	payload := url.Values{}
	payload.Set("value", value)

	body, err := l.client.Put("/labels/"+l.ID+"/"+path, payload)
	if err == nil {
		err = parseLabel(body, l, l.client)
	}
	return
}

// Delete - Delete a Label
// - https://developer.atlassian.com/cloud/trello/rest/api-group-labels/#api-labels-id-delete
func (l *Label) Delete() error {
	_, err := l.client.Delete("/labels/" + l.ID)
	return err
}

func parseLabel(body []byte, label *Label, client *Client) (err error) {
	err = json.Unmarshal(body, &label)
	if err == nil {
		label.client = client
	}
	return
}

func parseListLabels(body []byte, client *Client) (labels []Label, err error) {
	err = json.Unmarshal(body, &labels)
	for i := range labels {
		labels[i].client = client
	}
	return
}
