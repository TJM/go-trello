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
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestChecklist(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Checklist tests", func() {
		var board *Board
		var checklist *Checklist
		var checklistItem *ChecklistItem
		var testBoardName string

		// Prerequisites
		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Checklist-%v", time.Now().Unix())
			board, err = client.CreateBoard(testBoardName)
			if err != nil {
				log.Fatal("ERROR Creating Board: " + err.Error())
			}
			lists, err := board.Lists()
			if err != nil {
				log.Fatal("ERROR Retrieving board lists: " + err.Error())
			}
			if len(lists) < 1 {
				log.Fatal("ERROR: There should be at least one list on the test board. (by default)")
			}
			list := &lists[0]
			card, err := list.AddCard(Card{
				Name: "Testing 123",
				Desc: "Does this thing work?",
			})
			if err != nil || card == nil {
				log.Fatal("ERROR: Creating Card")
			}
			checklist, err = card.AddChecklist("TrelloTesting")
			if err != nil || checklist == nil {
				log.Fatal("ERROR: Creating Card")
			}
		})

		g.It("should error if checklist item name is empty", func() {
			checklistItem, err = checklist.AddItem("", "bottom", false)
			Expect(err).NotTo(BeNil())
			Expect(checklistItem).To(BeNil())
		})

		g.It("should error if checklist item name is too long", func() {
			tooLong := strings.Repeat("abcdefghijklmnop", 1024) + "z"
			Expect(len(tooLong)).To(BeNumerically(">", 16384))
			checklistItem, err = checklist.AddItem(tooLong, "bottom", false)
			Expect(err).NotTo(BeNil())
			Expect(checklistItem).To(BeNil())
		})

		g.It("should error if checklist item pos is a word other than top/bottom", func() {
			checklistItem, err = checklist.AddItem("Middle Test", "middle", false)
			Expect(err).NotTo(BeNil())
			Expect(checklistItem).To(BeNil())
		})

		g.It("should error if checklist item pos is a negative number string", func() {
			checklistItem, err = checklist.AddItem("Negative Test", "-1", false)
			Expect(err).NotTo(BeNil())
			Expect(checklistItem).To(BeNil())
		})

		g.It("should error if checklist item pos is a non-integer string", func() {
			checklistItem, err = checklist.AddItem("Negative Test", "1.69", false)
			Expect(err).NotTo(BeNil())
			Expect(checklistItem).To(BeNil())
		})

		g.It("should add an item to a checklist", func() {
			checklistItem, err = checklist.AddItem("Test Item", "bottom", false)
			Expect(err).To(BeNil())
			Expect(checklistItem.Name).To(Equal("Test Item"))
		})

		// Add this board test here since it gets Checklists
		g.It("should get the checklists in a board", func() {
			_, err := board.Checklists()
			Expect(err).To(BeNil())
		})

		// Destructive Actions second to last
		g.It("should delete a checklist item", func() {
			err = checklistItem.Delete()
			Expect(err).To(BeNil())
		})

		g.It("should delete a checklist", func() {
			err = checklist.Delete()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			if err != nil {
				log.Fatal("ERROR Deleting Board: " + err.Error())
			}
		})

	})

}
