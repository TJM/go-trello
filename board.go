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

// Board Type for Trello Board
type Board struct {
	client   *Client
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	DescData struct {
		Emoji struct{} `json:"emoji"`
	} `json:"descData"`
	Closed         bool   `json:"closed"`
	IDOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers"`
		CardAging             string            `json:"cardAging"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage"`
		BackgroundImageScaled []BoardBackground `json:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate"`
		CanInvite             bool              `json:"canInvite"`
	} `json:"prefs"`
	LabelNames struct {
		Red    string `json:"red"`
		Orange string `json:"orange"`
		Yellow string `json:"yellow"`
		Green  string `json:"green"`
		Blue   string `json:"blue"`
		Purple string `json:"purple"`
	} `json:"labelNames"`
}

// BoardBackground Type
type BoardBackground struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Board - Get board by boardID
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-get
func (c *Client) Board(boardID string) (board *Board, err error) {
	board = &Board{}
	body, err := c.Get("/boards/" + boardID)
	if err == nil {
		err = parseBoard(body, board, c)
	}
	return
}

// CreateBoard - Create Board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-post
func (c *Client) CreateBoard(name string) (board *Board, err error) {
	board = &Board{}
	payload := url.Values{}
	payload.Set("name", name)

	body, err := c.Post("/boards", payload)
	if err == nil {
		err = parseBoard(body, board, c)
	}
	return
}

// Duplicate - Duplicate (Copy) Board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-post
func (b *Board) Duplicate(name string, keepCards bool) (board *Board, err error) {
	board = &Board{}
	keepFromSource := "none"
	if keepCards {
		keepFromSource = "cards"
	}
	payload := url.Values{}
	payload.Set("idBoardSource", b.ID)
	payload.Set("keepFromSource", keepFromSource)
	payload.Set("name", name)

	body, err := b.client.Post("/boards", payload)
	if err == nil {
		err = parseBoard(body, board, b.client)
	}
	return
}

// SetBackground - Sets background on board
// background can be a color or a background id
func (b *Board) SetBackground(background string) (err error) {
	return b.Update("prefs/background", background)
}

// SetDescription - Sets background on board
func (b *Board) SetDescription(description string) (err error) {
	return b.Update("desc", description)
}

// Update - Update a Board (path and value, see API docs for details)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-put
func (b *Board) Update(path, value string) (err error) {
	payload := url.Values{}
	payload.Set("value", value)

	body, err := b.client.Put("/boards/"+b.ID+"/"+path, payload)
	if err == nil {
		err = parseBoard(body, b, b.client)
	}
	return
}

// Delete - Update a Board (path and value, see API docs for details)
//  *WARNING* - No Confirmation Dialog!
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-delete
func (b *Board) Delete() (err error) {
	_, err = b.client.Delete("/boards/" + b.ID)
	return
}

// Lists - Get lists on a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-lists-get
func (b *Board) Lists() (lists []List, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/lists")
	if err == nil {
		lists, err = parseListLists(body, b.client)
	}
	return
}

// Members - Get the members of a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-members-get
func (b *Board) Members() (members []Member, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/members")
	if err == nil {
		members, err = parseListMembers(body, b.client)
	}
	return
}

// Cards - Get cards on a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-cards-get
func (b *Board) Cards() (cards []Card, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/cards")
	if err == nil {
		cards, err = parseListCards(body, b.client)
	}
	return
}

// Card - Get a card on a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-cards-idcard-get
func (b *Board) Card(IDCard string) (card *Card, err error) {
	card = &Card{}
	body, err := b.client.Get("/boards/" + b.ID + "/cards/" + IDCard)
	if err == nil {
		err = parseCard(body, card, b.client)
	}
	return
}

// Checklists - Get checklists on a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-checklists-get
func (b *Board) Checklists() (checklists []Checklist, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/checklists")
	if err == nil {
		checklists, err = parseListChecklists(body, b.client)
	}
	return
}

// MemberCards - Get cards for a member ID (string) on a board?
// - URL Link?
func (b *Board) MemberCards(IDMember string) (cards []Card, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/members/" + IDMember + "/cards")
	if err == nil {
		cards, err = parseListCards(body, b.client)
	}
	return
}

// Actions - Get Actions for a Board
// - URL LINK?
func (b *Board) Actions(arg ...*Argument) (actions []Action, err error) {
	ep := "/boards/" + b.ID + "/actions"
	if query := EncodeArgs(arg); query != "" {
		ep += "?" + query
	}

	body, err := b.client.Get(ep)
	if err == nil {
		actions, err = parseListActions(body, b.client)
	}
	return
}

// AddList - Add a List to a Board
func (b *Board) AddList(opts List) (list *List, err error) {
	list = &List{}
	opts.IDBoard = b.ID

	payload := url.Values{}
	payload.Set("name", opts.Name)
	payload.Set("idBoard", opts.IDBoard)
	payload.Set("pos", strconv.FormatFloat(float64(opts.Pos), 'g', -1, 32))

	body, err := b.client.Post("/lists", payload)
	if err == nil {
		err = parseList(body, list, b.client)
	}
	return
}

// Labels - Get Labels on a Board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-labels-get
func (b *Board) Labels() (labels []Label, err error) {
	body, err := b.client.Get("/boards/" + b.ID + "/labels")
	if err == nil {
		labels, err = parseListLabels(body, b.client)
	}
	return
}

// AddLabel - Create a Label on a board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-labels-post
// NOTE: Color can be an empty string
func (b *Board) AddLabel(name, color string) (label *Label, err error) {
	label = &Label{}
	payload := url.Values{}
	payload.Set("name", name)
	payload.Set("color", color)

	body, err := b.client.Post("/boards/"+b.ID+"/labels", payload)
	if err == nil {
		err = parseLabel(body, label, b.client)
	}
	return
}

func parseBoard(body []byte, board *Board, client *Client) (err error) {
	err = json.Unmarshal(body, &board)
	if err == nil {
		board.client = client
	}
	return
}

func parseListBoards(body []byte, client *Client) (boards []Board, err error) {
	err = json.Unmarshal(body, &boards)
	if err == nil {
		for i := range boards {
			boards[i].client = client
		}
	}
	return
}
