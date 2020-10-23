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

package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/TJM/go-trello"
	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestCard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Card tests", func() {
		var Board *trello.Board
		var Card *trello.Card
		var Member *trello.Member
		var TestBoardName string

		g.Before(func() {
			TestBoardName = fmt.Sprintf("GoTestTrello-Card-%v", time.Now())
			Member, err = Client.Member("me")
			Expect(err).To(BeNil())
			Board, err = Client.CreateBoard(TestBoardName)
			Expect(err).To(BeNil())
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			list := &lists[0]
			Card, err = list.AddCard(trello.Card{
				Name: "Testing 123",
				Desc: "Does this thing work?",
			})
			Expect(err).To(BeNil())
		})

		g.It("should retrieve a card by ID", func() {
			_, err = Client.Card(Card.ID)
			Expect(err).To(BeNil())
		})

		g.It("should get the checklists in a card", func() {
			_, err := Card.Checklists()
			Expect(err).To(BeNil())
		})

		g.It("should add a member to a card", func() {
			_, err = Card.AddMember(Member)
			Expect(err).To(BeNil())
		})

		g.It("should get the members of a card", func() {
			_, err = Card.Members()
			Expect(err).To(BeNil())
			// It might be nice to check if "me" is in the members?
		})

		g.It("should remove a member from a card", func() {
			_, err = Card.RemoveMember(Member)
			Expect(err).To(BeNil())
			// It might be nice to check if "me" is NOT in the members?
		})

		g.It("should get the attachments on a card", func() {
			_, err = Card.Attachments()
			Expect(err).To(BeNil())
			// It might be nice to check attachments?
		})

		g.It("should get the actions on a card", func() {
			_, err = Card.Actions()
			Expect(err).To(BeNil())
			// It might be nice to check attachments?
		})

		g.It("should add a checklist to a card", func() {
			_, err = Card.AddChecklist("TrelloTesting")
			Expect(err).To(BeNil())
		})

		g.It("should add a comment to a card", func() {
			_, err = Card.AddComment("Test Comment")
			Expect(err).To(BeNil())
		})

		g.It("should move a card to another list", func() {
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			dest := lists[1]
			Card, err = Card.MoveToList(dest)
			Expect(err).To(BeNil())
			Expect(Card.IDList).To(Equal(dest.ID))
		})

		g.It("should move a card position in the list (top)", func() {
			Card, err = Card.Move("top")
			Expect(err).To(BeNil())
		})

		g.It("should setName on a card", func() {
			Card, err = Card.SetName("A whole new name")
			Expect(err).To(BeNil())
			Expect(Card.Name).To(Equal("A whole new name"))
		})

		g.It("should setDescription on a card", func() {
			Card, err = Card.SetDescription("A whole new description")
			Expect(err).To(BeNil())
			Expect(Card.Desc).To(Equal("A whole new description"))
		})

		g.It("should add a label to a card", func() {
			labels, err := Board.Labels()
			Expect(err).To(BeNil())
			_, err = Card.AddLabel(labels[0].ID)
			Expect(err).To(BeNil())
		})

		g.It("should add a NEW label to a card", func() {
			_, err = Card.AddNewLabel("NewLabel", "purple")
			Expect(err).To(BeNil())
		})

		// Destructive Actions second to last
		g.It("should archive (close) a card", func() {
			err = Card.Archive(true)
			Expect(err).To(BeNil())
		})

		g.It("should un-archive (re-open) a card", func() {
			err = Card.Archive(false)
			Expect(err).To(BeNil())
		})

		g.It("should delete a card", func() {
			err = Card.Delete()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = Board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
