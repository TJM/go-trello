/*
Copyright 2014 go-trello authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trello

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// List - Trello List Type
type List struct {
	client  *Client
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Closed  bool    `json:"closed"`
	IDBoard string  `json:"idBoard"`
	Pos     float32 `json:"pos"`
	cards   []Card
}

// List - Get List by listID (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-get
func (c *Client) List(listID string) (list *List, err error) {
	list = &List{}
	body, err := c.Get("/lists/" + listID)
	if err == nil {
		err = parseList(body, list, c)
	}
	return
}

// Cards - Get Cards in a List
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-cards-get
func (l *List) Cards() (cards []Card, err error) {
	body, err := l.client.Get("/lists/" + l.ID + "/cards")
	if err == nil {
		cards, err = parseListCards(body, l.client)
	}
	return
}

// Actions - Get Actions for a List
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-actions-get
func (l *List) Actions() (actions []Action, err error) {
	body, err := l.client.Get("/lists/" + l.ID + "/actions")
	if err == nil {
		actions, err = parseListActions(body, l.client)
	}
	return
}

// AddCard creates with the attributes of the supplied Card struct
// https://developers.trello.com/advanced-reference/card#post-1-cards
func (l *List) AddCard(opts Card) (card *Card, err error) {
	card = &Card{}
	opts.IDList = l.ID

	payload := url.Values{}
	payload.Set("name", opts.Name)
	payload.Set("desc", opts.Desc)
	if opts.Pos == 0.0 {
		payload.Set("pos", "bottom")
	} else {
		payload.Set("pos", strconv.FormatFloat(opts.Pos, 'g', -1, 64))
	}
	payload.Set("due", opts.Due)
	payload.Set("idList", opts.IDList)
	payload.Set("idMembers", strings.Join(opts.IDMembers, ","))

	body, err := l.client.Post("/cards", payload)
	if err == nil {
		err = parseCard(body, card, l.client)
	}
	return
}

// Archive - Archive List
//If mode is true, list is archived, otherwise it's unarchived (returns to the board)
func (l *List) Archive(mode bool) (err error) {
	payload := url.Values{}
	payload.Set("value", strconv.FormatBool(mode))

	body, err := l.client.Put("/lists/"+l.ID+"/closed", payload)
	if err == nil {
		err = parseList(body, l, l.client)
	}
	return
}

// Move - Move a List (Update a List)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-put
//pos can be "bottom", "top" or a positive number
func (l *List) Move(pos string) (err error) {
	payload := url.Values{}
	payload.Set("value", pos)

	body, err := l.client.Put("/lists/"+l.ID+"/pos", payload)
	if err == nil {
		err = parseList(body, l, l.client)
	}
	return
}

func parseList(body []byte, list *List, client *Client) (err error) {
	err = json.Unmarshal(body, &list)
	if err == nil {
		list.client = client
	}
	return
}

func parseListLists(body []byte, client *Client) (lists []List, err error) {
	err = json.Unmarshal(body, &lists)
	for i := range lists {
		lists[i].client = client
	}
	return
}
