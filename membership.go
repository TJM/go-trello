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
)

// Membership tello membership struct
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-memberships-get
type Membership struct {
	client      *Client
	Board       *Board
	ID          string `json:"id" pattern:"^[0-9a-fA-F]{32}$"` // Pattern: ^[0-9a-fA-F]{32}$
	IDMember    string `json:"idMember"`
	MemberType  string `json:"memberType"`
	Unconfirmed bool   `json:"unconfirmed"`
	Deactivated bool   `json:"deactivated"`
}

// Update - Update Membership of Member on a Board
// memberType can be admin, normal or observer (if left blank will default to normal)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-id-memberships-idmembership-put
func (m *Membership) Update(memberType, memberFields string) (err error) {
	if memberType == "" {
		memberType = "normal"
	}
	payload := url.Values{}
	payload.Set("type", memberType)
	if memberFields != "" {
		payload.Set("member_fields", memberFields)
	}
	body, err := m.client.Put("/boards/"+m.Board.ID+"/memberships/"+m.ID, payload)
	if err == nil {
		err = parseMembership(body, m, m.Board)
	}
	return
}

func parseMembership(body []byte, membership *Membership, board *Board) (err error) {
	err = json.Unmarshal(body, &membership)
	if err == nil {
		membership.client = board.client
		membership.Board = board
	}
	return
}

func parseListMemberships(body []byte, board *Board) (memberships []*Membership, err error) {
	err = json.Unmarshal(body, &memberships)
	for i := range memberships {
		memberships[i].client = board.client
		memberships[i].Board = board
	}
	return
}
