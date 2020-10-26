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
)

// ChecklistItem - Trello Checklist Item (member of Checklist)
type ChecklistItem struct {
	client   *Client
	listID   string // back pointer to the parent ID
	State    string `json:"state"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	NameData struct {
		Emoji struct{} `json:"emoji"`
	} `json:"nameData"`
	Pos int `json:"pos"`
}

// Delete - Delete a ChecklistItem from Checklist
// - https://developer.atlassian.com/cloud/trello/rest/api-group-checklists/#api-checklists-id-checkitems-idcheckitem-delete
func (i *ChecklistItem) Delete() error {
	_, err := i.client.Delete("/checklists/" + i.listID + "/checkItems/" + i.ID)
	return err
}

func parseChecklistItem(body []byte, checklistItem *ChecklistItem, client *Client, listID string) (err error) {
	err = json.Unmarshal(body, &checklistItem)
	if err == nil {
		checklistItem.client = client
		checklistItem.listID = listID
	}
	return
}

// func parseListChecklistItems(body []byte, client *Client) (checklistItems []Checklist, err error) {
// 	err = json.Unmarshal(body, &checklistItems)
// 	for i := range checklistItems {
// 		list := checklistItems[i]
// 		list.client = client
// 		for j := range list.CheckItems {
// 			item := list.CheckItems[j]
// 			item.client = client
// 			item.listID = list.ID
// 		}
// 	}
// 	return
// }
