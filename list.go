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
	body, err := c.Get("/lists/" + listID)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &list)
	list.client = c
	return
}

// Cards - Get Cards in a List
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-cards-get
func (l *List) Cards() (cards []Card, err error) {
	if l.cards != nil {
		return l.cards, nil
	}
	body, err := l.client.Get("/lists/" + l.ID + "/cards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cards)
	for i := range cards {
		cards[i].client = l.client
		for j := range cards[i].Labels {
			cards[i].Labels[j].client = l.client
		}
	}
	l.cards = cards
	return
}

// FreshCards - nil cards?
func (l *List) FreshCards() (cards []Card, err error) {
	l.cards = nil
	return l.Cards()
}

// Actions - Get Actions for a List
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-actions-get
func (l *List) Actions() (actions []Action, err error) {
	body, err := l.client.Get("/lists/" + l.ID + "/actions")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &actions)
	for i := range actions {
		actions[i].client = l.client
	}
	return
}

// AddCard creates with the attributes of the supplied Card struct
// https://developers.trello.com/advanced-reference/card#post-1-cards
func (l *List) AddCard(opts Card) (*Card, error) {
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
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = l.client
	return &card, nil
}

// Archive - Archive List
//If mode is true, list is archived, otherwise it's unarchived (returns to the board)
func (l *List) Archive(mode bool) error {
	payload := url.Values{}
	payload.Set("value", strconv.FormatBool(mode))

	_, err := l.client.Put("/lists/"+l.ID+"/closed", payload)
	return err
}

// Move - Move a List (Update a List)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-lists/#api-lists-id-put
//pos can be "bottom", "top" or a positive number
func (l *List) Move(pos string) (*List, error) {
	payload := url.Values{}
	payload.Set("value", pos)

	body, err := l.client.Put("/lists/"+l.ID+"/pos", payload)
	if err != nil {
		return nil, err
	}

	var list List
	if err = json.Unmarshal(body, &list); err != nil {
		return nil, err
	}
	list.client = l.client
	return &list, nil
}
