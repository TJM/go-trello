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
	"strings"
)

// Member tello member struct
type Member struct {
	client     *Client
	ID         string `json:"id"`
	AvatarHash string `json:"avatarHash"`
	Bio        string `json:"bio"`
	BioData    struct {
		Emoji interface{} `json:"emoji,omitempty"`
	} `json:"bioData"`
	Confirmed                bool     `json:"confirmed"`
	FullName                 string   `json:"fullName"`
	IDPremOrgsAdmin          []string `json:"idPremOrgsAdmin"`
	Initials                 string   `json:"initials"`
	MemberType               string   `json:"memberType"`
	Products                 []int    `json:"products"`
	Status                   string   `json:"status"`
	URL                      string   `json:"url"`
	Username                 string   `json:"username"`
	AvatarSource             string   `json:"avatarSource"`
	Email                    string   `json:"email"`
	GravatarHash             string   `json:"gravatarHash"`
	IDBoards                 []string `json:"idBoards"`
	IDBoardsPinned           []string `json:"idBoardsPinned"`
	IDOrganizations          []string `json:"idOrganizations"`
	LoginTypes               []string `json:"loginTypes"`
	NewEmail                 string   `json:"newEmail"`
	OneTimeMessagesDismissed []string `json:"oneTimeMessagesDismissed"`
	Prefs                    struct {
		SendSummaries                 bool   `json:"sendSummaries"`
		MinutesBetweenSummaries       int    `json:"minutesBetweenSummaries"`
		MinutesBeforeDeadlineToNotify int    `json:"minutesBeforeDeadlineToNotify"`
		ColorBlind                    bool   `json:"colorBlind"`
		Locale                        string `json:"locale"`
	} `json:"prefs"`
	Trophies           []string `json:"trophies"`
	UploadedAvatarHash string   `json:"uploadedAvatarHash"`
	PremiumFeatures    []string `json:"premiumFeatures"`
}

// Member returns a member (NOTE: "me" defaults to yourself)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-members/#api-members-id-get
func (c *Client) Member(nick string) (member *Member, err error) {
	member = &Member{}
	body, err := c.Get("/members/" + nick)
	if err == nil {
		err = parseMember(body, member, c)
	}
	return
}

// Boards returns members boards
// - https://developer.atlassian.com/cloud/trello/rest/api-group-members/#api-members-id-boards-get
func (m *Member) Boards(field ...string) (boards []*Board, err error) {
	fields := ""
	if len(field) == 0 {
		fields = "all"
	} else {
		fields = strings.Join(field, ",")
	}

	body, err := m.client.Get("/members/" + m.ID + "/boards?fields=" + fields)
	if err == nil {
		boards, err = parseListBoards(body, m.client)
	}
	return
}

// AddBoard creates a new Board
// - https://developer.atlassian.com/cloud/trello/rest/api-group-boards/#api-boards-post
func (m *Member) AddBoard(name string) (*Board, error) {
	return m.client.CreateBoard(name)
}

// Notifications - https://developer.atlassian.com/cloud/trello/rest/api-group-members/#api-members-id-notifications-get
func (m *Member) Notifications() (notifications []Notification, err error) {
	body, err := m.client.Get("/members/" + m.ID + "/notifications")
	if err == nil {
		notifications, err = parseListNotifications(body, m.client)
	}
	return
}

// AvatarURL returns avatar URL for member
// TODO: Avatar sizes [170, 30]
func (m *Member) AvatarURL() string {
	return "https://trello-avatars.s3.amazonaws.com/" + m.AvatarHash + "/170.png"
}

func parseMember(body []byte, member *Member, client *Client) (err error) {
	err = json.Unmarshal(body, &member)
	if err == nil {
		member.client = client
	}
	return
}

func parseListMembers(body []byte, client *Client) (members []*Member, err error) {
	err = json.Unmarshal(body, &members)
	for i := range members {
		members[i].client = client
	}
	return
}
