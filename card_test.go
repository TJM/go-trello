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
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestCard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Card tests", func() {
		var board *Board
		var card *Card
		var member *Member
		var attachments []Attachment
		var testBoardName string

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Card-%v", time.Now().Unix())
			member, err = client.Member("me")
			Expect(err).To(BeNil())
			board, err = client.CreateBoard(testBoardName)
			Expect(err).To(BeNil())
			lists, err := board.Lists()
			Expect(err).To(BeNil())
			list := &lists[0]
			card, err = list.AddCard(Card{
				Name: "Testing 123",
				Desc: "Does this thing work?",
			})
			Expect(err).To(BeNil())
			Expect(card).NotTo(BeNil())
		})

		g.It("should retrieve a card by ID from client", func() {
			c, err := client.Card(card.ID)
			Expect(err).To(BeNil())
			Expect(c).NotTo(BeNil())
		})

		g.It("should retrieve a card by ID from board", func() {
			c, err := board.Card(card.ID)
			Expect(err).To(BeNil())
			Expect(c).NotTo(BeNil())
		})

		g.It("should add a member to a card", func() {
			_, err = card.AddMember(member)
			Expect(err).To(BeNil())
		})

		g.It("should get the members of a card", func() {
			_, err = card.Members()
			Expect(err).To(BeNil())
			// It might be nice to check if "me" is in the members?
		})

		// Add this board test here cause it gets cards
		g.It("should get the membercards in a board", func() {
			_, err = board.MemberCards(member.ID)
			Expect(err).To(BeNil())
		})

		g.It("should remove a member from a card", func() {
			_, err = card.RemoveMember(member)
			Expect(err).To(BeNil())
			// It might be nice to check if "me" is NOT in the members?
		})

		g.It("should get the attachments on a card", func() {
			attachments, err = card.Attachments()
			Expect(err).To(BeNil())
			Expect(attachments).NotTo(BeNil())
		})

		// We need a way to "add" an attachment before we can retriecve it
		// g.It("should get an attachment by id from a card", func() {
		// 	if len(attachments) == 0 {
		// 		t.Skip("No Attachments")
		// 	}
		// 	_, err = card.Attachment(attachments[0].ID)
		// 	Expect(err).To(BeNil())
		// 	// It might be nice to check attachments?
		// })

		g.It("should get an error for get invalid attachmentid from a card", func() {
			_, err = card.Attachment("invalid")
			Expect(err).NotTo(BeNil())
			// It might be nice to check attachments?
		})

		g.It("should get the actions on a card", func() {
			_, err = card.Actions()
			Expect(err).To(BeNil())
			// It might be nice to check attachments?
		})

		g.It("should add a checklist to a card", func() {
			checklist, err := card.AddChecklist("TrelloChecklistTest")
			Expect(err).To(BeNil())
			// Add Item
			_, err = checklist.AddItem("Test Item", "bottom", false)
			Expect(err).To(BeNil())
		})

		g.It("should get the checklists in a card", func() {
			checklists, err := card.Checklists()
			Expect(err).To(BeNil())
			Expect(len(checklists)).To(BeNumerically(">", 0))
		})

		g.It("should add a comment to a card", func() {
			_, err = card.AddComment("Test Comment")
			Expect(err).To(BeNil())
		})

		g.It("should move a card to another list", func() {
			lists, err := board.Lists()
			Expect(err).To(BeNil())
			dest := lists[1]
			err = card.MoveToList(dest)
			Expect(err).To(BeNil())
			Expect(card.IDList).To(Equal(dest.ID))
		})

		g.It("should move a card position in the list (top)", func() {
			err = card.Move("top")
			Expect(err).To(BeNil())
		})

		g.It("should move a card position in the list (numeric)", func() {
			err = card.Move("2")
			Expect(err).To(BeNil())
		})

		g.It("should setName on a card", func() {
			err = card.SetName("A whole new name")
			Expect(err).To(BeNil())
			Expect(card.Name).To(Equal("A whole new name"))
		})

		g.It("should setDescription on a card", func() {
			err = card.SetDescription("A whole new description")
			Expect(err).To(BeNil())
			Expect(card.Desc).To(Equal("A whole new description"))
		})

		g.It("should add a label to a card", func() {
			labels, err := board.Labels()
			Expect(err).To(BeNil())
			_, err = card.AddLabel(labels[0].ID)
			Expect(err).To(BeNil())
		})

		g.It("should add a NEW label to a card", func() {
			_, err = card.AddNewLabel("NewLabel", "purple")
			Expect(err).To(BeNil())
		})

		// Add this board test here, cause it gets cards
		g.It("should get the cards in a board", func() {
			_, err := board.Cards()
			Expect(err).To(BeNil())
		})

		// Destructive Actions second to last
		g.It("should archive (close) a card", func() {
			err = card.Archive(true)
			Expect(err).To(BeNil())
		})

		g.It("should un-archive (re-open) a card", func() {
			err = card.Archive(false)
			Expect(err).To(BeNil())
		})

		g.It("should delete a card", func() {
			err = card.Delete()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
