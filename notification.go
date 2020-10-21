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

import "encoding/json"

// Notification - Trello Notification Type
type Notification struct {
	client *Client
	ID     string `json:"id"`
	Unread bool   `json:"unread"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	Data   struct {
		ListBefore struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"listBefore"`
		ListAfter struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"listAfter"`
		Board struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
		} `json:"board"`
		Card struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			ShortLink string `json:"shortLink"`
			IDShort   int    `json:"idShort"`
		} `json:"card"`
		Old struct {
			IDList string `json:"idList"`
		} `json:"old"`
	} `json:"data"`
	IDMemberCreator string `json:"idMemberCreator"`
	MemberCreator   struct {
		ID         string `json:"id"`
		AvatarHash string `json:"avatarHash"`
		FullName   string `json:"fullName"`
		Initials   string `json:"initials"`
		Username   string `json:"username"`
	} `json:"memberCreator"`
}

// Notification - Get Notification by notificationId (string)
// - https://developer.atlassian.com/cloud/trello/rest/api-group-notifications/#api-notifications-id-get
func (c *Client) Notification(notificationID string) (notification *Notification, err error) {
	body, err := c.Get("/notifications/" + notificationID)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &notification)
	notification.client = c
	return
}
