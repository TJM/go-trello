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

// Organization - Trello Organization Type
type Organization struct {
	client      *Client
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Desc        string   `json:"desc"`
	DescData    string   `json:"descData"`
	URL         string   `json:"url"`
	Website     string   `json:"website"`
	LogoHash    string   `json:"logoHash"`
	Products    []string `json:"products"`
	PowerUps    []string `json:"powerUps"`
}

// Organization - Get Organization by orgId (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-organizations/#api-organizations-id-get
func (c *Client) Organization(orgID string) (organization *Organization, err error) {
	body, err := c.Get("/organization/" + orgID)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &organization)
	organization.client = c
	return
}

// Members - Get the Members of an Organization
// - https://developer.atlassian.com/cloud/trello/rest/api-group-organizations/#api-organizations-id-members-get
func (o *Organization) Members() (members []Member, err error) {
	body, err := o.client.Get("/organization/" + o.ID + "/members")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &members)
	for i := range members {
		members[i].client = o.client
	}
	return
}

// Boards - Get Boards in an Organization
// - https://developer.atlassian.com/cloud/trello/rest/api-group-organizations/#api-organizations-id-boards-get
func (o *Organization) Boards() (boards []Board, err error) {
	body, err := o.client.Get("/organizations/" + o.ID + "/boards")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &boards)
	for i := range boards {
		boards[i].client = o.client
	}
	return
}

// AddBoard  - Create a new Board in the organization
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-post
func (o *Organization) AddBoard(name string) (*Board, error) {

	payload := url.Values{}
	payload.Set("name", name)
	payload.Set("idOrganization", o.ID)

	body, err := o.client.Post("/boards", payload)
	if err != nil {
		return nil, err
	}
	var board Board
	if err = json.Unmarshal(body, &board); err != nil {
		return nil, err
	}

	board.client = o.client
	return &board, nil
}
