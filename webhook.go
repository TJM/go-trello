package trello

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Webhook - Trello Webhook Type
type Webhook struct {
	client      *Client
	ID          string `json:"id"`
	Description string `json:"description"`
	IDModel     string `json:"idModel"`
	CallbackURL string `json:"callbackURL"`
	Active      bool   `json:"active"`
}

// Webhooks - Get Webhooks for a token (string)
// https://developer.atlassian.com/cloud/trello/rest/api-group-tokens/#api-tokens-token-webhooks-get
func (c *Client) Webhooks(token string) (webhooks []Webhook, err error) {

	body, err := c.Get(webhookURL(token))
	if err != nil {
		return []Webhook{}, err
	}
	err = json.Unmarshal(body, &webhooks)
	return
}

// CreateWebhook - Create a Webhook
// - https://developer.atlassian.com/cloud/trello/rest/api-group-webhooks/#api-webhooks-post
func (c *Client) CreateWebhook(hook Webhook) (webhook *Webhook, err error) {
	webhook = &Webhook{}
	payload := url.Values{}
	payload.Set("description", hook.Description)
	payload.Set("callbackURL", hook.CallbackURL)
	payload.Set("idModel", hook.IDModel)
	body, err := c.Post("/webhooks/", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &webhook)
	webhook.client = c
	return
}

// Webhook - Get Webhook by id (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-webhooks/#api-webhooks-id-get
func (c *Client) Webhook(webhookID string) (webhook *Webhook, err error) {
	webhook = &Webhook{}
	url := fmt.Sprintf("/webhooks/%s/", webhookID)
	body, err := c.Get(url)
	if err == nil {
		err = parseWebhook(body, webhook, c)
	}
	return
}

// SetActive - Set Active for WebHook
// true - active, false - inactive
func (w *Webhook) SetActive(active bool) (err error) {
	payload := url.Values{}
	payload.Set("active", strconv.FormatBool(active))
	return w.update(payload)
}

// SetCallbackURL - Set SetCallbackURL for WebHook
func (w *Webhook) SetCallbackURL(callbackURL string) (err error) {
	payload := url.Values{}
	payload.Set("callbackURL", callbackURL)
	return w.update(payload)
}

// SetDescription - Set Description for WebHook
func (w *Webhook) SetDescription(description string) (err error) {
	payload := url.Values{}
	payload.Set("description", description)
	return w.update(payload)
}

// SetIDModel - Set IDModel for WebHook
func (w *Webhook) SetIDModel(model string) (err error) {
	payload := url.Values{}
	payload.Set("idModel", model)
	return w.update(payload)
}

// update - Update a Webhook (with payload)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-webhooks/#api-webhooks-id-put
func (w *Webhook) update(payload url.Values) (err error) {
	body, err := w.client.Put("/webhooks/"+w.ID, payload)
	if err == nil {
		err = parseWebhook(body, w, w.client)
	}
	return
}

// Delete - Delete a Webhook (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-webhooks/#api-webhooks-id-delete
func (w *Webhook) Delete() (err error) {
	return w.client.DeleteWebhook(w.ID)
}

// DeleteWebhook - Delete a Webhook by id (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-webhooks/#api-webhooks-id-delete
func (c *Client) DeleteWebhook(webhookID string) (err error) {

	url := fmt.Sprintf("/webhooks/%s/", webhookID)
	_, err = c.Delete(url)
	return
}

func webhookURL(token string) (url string) {

	return fmt.Sprintf("/tokens/%s/webhooks/", token)
}

func parseWebhook(body []byte, webhook *Webhook, client *Client) (err error) {
	err = json.Unmarshal(body, &webhook)
	if err == nil {
		webhook.client = client
	}
	return
}

// func parseListWebhooks(body []byte, client *Client) (webhooks []Webhook, err error) {
// 	err = json.Unmarshal(body, &webhooks)
// 	for i := range webhooks {
// 		webhooks[i].client = client
// 	}
// 	return
// }
