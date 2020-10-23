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
)

// Card - Trello Card Type
type Card struct {
	client                *Client
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	Email                 string   `json:"email"`
	IDShort               int      `json:"idShort"`
	IDAttachmentCover     string   `json:"idAttachmentCover"`
	IDCheckLists          []string `json:"idCheckLists"`
	IDBoard               string   `json:"idBoard"`
	IDList                string   `json:"idList"`
	IDMembers             []string `json:"idMembers"`
	IDMembersVoted        []string `json:"idMembersVoted"`
	ManualCoverAttachment bool     `json:"manualCoverAttachment"`
	Closed                bool     `json:"closed"`
	Pos                   float64  `json:"pos"`
	ShortLink             string   `json:"shortLink"`
	DateLastActivity      string   `json:"dateLastActivity"`
	ShortURL              string   `json:"shortUrl"`
	Subscribed            bool     `json:"subscribed"`
	URL                   string   `json:"url"`
	Due                   string   `json:"due"`
	Desc                  string   `json:"desc"`
	DescData              struct {
		Emoji struct{} `json:"emoji"`
	} `json:"descData"`
	CheckItemStates []struct {
		IDCheckItem string `json:"idCheckItem"`
		State       string `json:"state"`
	} `json:"checkItemStates"`
	Badges struct {
		Votes              int    `json:"votes"`
		ViewingMemberVoted bool   `json:"viewingMemberVoted"`
		Subscribed         bool   `json:"subscribed"`
		Fogbugz            string `json:"fogbugz"`
		CheckItems         int    `json:"checkItems"`
		CheckItemsChecked  int    `json:"checkItemsChecked"`
		Comments           int    `json:"comments"`
		Attachments        int    `json:"attachments"`
		Description        bool   `json:"description"`
		Due                string `json:"due"`
	} `json:"badges"`
	Labels []Label `json:"labels"`
}

// Card - Retrieve card by card ID
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-get
func (c *Client) Card(CardID string) (card *Card, err error) {
	body, err := c.Get("/card/" + CardID)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &card)
	card.client = c
	return
}

// Checklists - Get Checklists on a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-checklists-get
func (c *Card) Checklists() (checklists []Checklist, err error) {
	body, err := c.client.Get("/card/" + c.ID + "/checklists")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &checklists)
	for i := range checklists {
		list := &checklists[i]
		list.client = c.client
		for i := range list.CheckItems {
			item := &list.CheckItems[i]
			item.client = c.client
			item.listID = list.ID
		}
	}
	return
}

// Members - Get the Members of a card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-members-get
func (c *Card) Members() (members []Member, err error) {
	body, err := c.client.Get("/cards/" + c.ID + "/members")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &members)
	for i := range members {
		members[i].client = c.client
	}
	return
}

// AddMember - Add a member to a card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-idmembers-post
// The AddMember function requires a member (pointer) to add
// It returns the resulting member-list
// https://developers.trello.com/v1.0/reference#cardsididmembers
func (c *Card) AddMember(member *Member) (members []Member, err error) {
	payload := url.Values{}
	payload.Set("value", member.ID)
	body, err := c.client.Post("/cards/"+c.ID+"/idMembers", payload)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &members); err != nil {
		return nil, err
	}

	// To enable our members to execute operations using our client, we need to pass each our client object
	for i := range members {
		members[i].client = c.client
	}
	return members, nil
}

// RemoveMember - Remove a member from a card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-idmembers-idmember-delete
// The RemoveMember function requires a member (pointer) to delete
// It returns the resulting member-list
func (c *Card) RemoveMember(member *Member) (members []Member, err error) {
	body, err := c.client.Delete("/cards/" + c.ID + "/idMembers/" + member.ID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &members)
	if err != nil {
		return nil, err
	}

	// To enable our members to execute operations using our client, we need to pass each our client object
	for i := range members {
		members[i].client = c.client
	}
	return members, nil
}

// Attachments - Get Attachments on a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-attachments-get
func (c *Card) Attachments() (attachments []Attachment, err error) {
	body, err := c.client.Get("/cards/" + c.ID + "/attachments")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &attachments)
	for i := range attachments {
		attachments[i].client = c.client
	}
	return
}

// Attachment will return the specified attachment on the card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-attachments-idattachment-get
// https://developers.trello.com/advanced-reference/card#get-1-cards-card-id-or-shortlink-attachments-idattachment
func (c *Card) Attachment(attachmentID string) (*Attachment, error) {
	body, err := c.client.Get("/cards/" + c.ID + "/attachments/" + attachmentID)
	if err != nil {
		return nil, err
	}

	attachment := &Attachment{}
	err = json.Unmarshal(body, attachment)
	attachment.client = c.client
	return attachment, err
}

// Actions - Get Actions on a card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-actions-get
func (c *Card) Actions() (actions []Action, err error) {
	body, err := c.client.Get("/cards/" + c.ID + "/actions")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &actions)
	for i := range actions {
		actions[i].client = c.client
	}
	return
}

// AddChecklist - Create a Checklist on a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-checklists-post
func (c *Card) AddChecklist(name string) (*Checklist, error) {
	newList := &Checklist{}

	payload := url.Values{}
	payload.Set("name", name)
	body, err := c.client.Post("/cards/"+c.ID+"/checklists", payload)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, newList); err != nil {
		return nil, err
	}
	newList.client = c.client
	// the new list has no items, no need to walk those adding client
	return newList, err
}

// AddComment will add a new comment to the card
// https://developers.trello.com/advanced-reference/card#post-1-cards-card-id-or-shortlink-actions-comments
func (c *Card) AddComment(text string) (*Action, error) {
	newAction := &Action{}

	payload := url.Values{}
	payload.Set("text", text)

	body, err := c.client.Post("/cards/"+c.ID+"/actions/comments", payload)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, newAction); err != nil {
		return nil, err
	}
	newAction.client = c.client
	return newAction, nil
}

// MoveToList - Move a card to a list
func (c *Card) MoveToList(dstList List) (*Card, error) {
	payload := url.Values{}
	payload.Set("value", dstList.ID)

	body, err := c.client.Put("/cards/"+c.ID+"/idList", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = c.client
	return &card, nil
}

// Move - Move Card Position (Update a Card)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-put
//pos can be "bottom", "top" or a positive number
func (c *Card) Move(pos string) (*Card, error) {
	payload := url.Values{}
	payload.Set("value", pos)

	body, err := c.client.Put("/cards/"+c.ID+"/pos", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = c.client
	return &card, nil
}

// Delete - Delete a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-delete
func (c *Card) Delete() error {
	_, err := c.client.Delete("/cards/" + c.ID)
	return err
}

// Archive - Archive (close) a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-put
//If mode is true, card is archived, otherwise it's unarchived (returns to the board)
func (c *Card) Archive(mode bool) error {
	payload := url.Values{}
	payload.Set("value", strconv.FormatBool(mode))

	_, err := c.client.Put("/cards/"+c.ID+"/closed", payload)
	return err
}

// SetName - Set Name on Card (Update a Card)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-put
func (c *Card) SetName(name string) (*Card, error) {
	payload := url.Values{}
	payload.Set("value", name)

	body, err := c.client.Put("/cards/"+c.ID+"/name", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = c.client
	return &card, nil
}

// SetDescription - Set Description on Card (Update a Card)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-put
func (c *Card) SetDescription(desc string) (*Card, error) {
	payload := url.Values{}
	payload.Set("value", desc)

	body, err := c.client.Put("/cards/"+c.ID+"/desc", payload)
	if err != nil {
		return nil, err
	}

	var card Card
	if err = json.Unmarshal(body, &card); err != nil {
		return nil, err
	}
	card.client = c.client
	return &card, nil
}

// AddLabel - Add Label to a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-idlabels-post
// Returns an array of cards labels ids
func (c *Card) AddLabel(id string) ([]string, error) {
	payload := url.Values{}
	payload.Set("value", id)

	body, err := c.client.Post("/cards/"+c.ID+"/idLabels", payload)
	if err != nil {
		return nil, err
	}

	var ids []string
	if err = json.Unmarshal(body, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

// AddNewLabel - Add a Label to a Card
// - https://developer.atlassian.com/cloud/trello/rest/api-group-cards/#api-cards-id-idlabels-post
func (c *Card) AddNewLabel(name, color string) (*Label, error) {
	payload := url.Values{}
	payload.Set("name", name)
	payload.Set("color", color)

	body, err := c.client.Post("/cards/"+c.ID+"/labels", payload)
	if err != nil {
		return nil, err
	}

	var label Label
	if err = json.Unmarshal(body, &label); err != nil {
		return nil, err
	}

	label.client = c.client
	return &label, nil
}
