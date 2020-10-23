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

func TestBoard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Board tests", func() {
		var Board *trello.Board
		var TestBoardName string

		g.Before(func() {
			TestBoardName = fmt.Sprintf("GoTestTrello-Board-%v", time.Now())

		})

		g.It("should create a board", func() {
			Board, err = Client.CreateBoard(TestBoardName)
			Expect(err).To(BeNil())
			Expect(Board.Name).To(Equal(TestBoardName))
		})

		g.It("should get a board by ID", func() {
			Board, err = Client.Board(Board.ID)
			Expect(err).To(BeNil())
			Expect(Board.Name).To(Equal(TestBoardName))
		})

		g.It("should change the board background to red", func() {
			err = Board.SetBackground("red")
			Expect(err).To(BeNil())
			Expect(Board.Prefs.Background).To(Equal("red"))
		})

		g.It("should change the description to something", func() {
			err = Board.SetDescription("something")
			Expect(err).To(BeNil())
			Expect(Board.Desc).To(Equal("something"))
		})

		g.It("should get the lists in a board", func() {
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			// This part is somewhat dangerous if Trello changes the default Board Template
			Expect(lists[0].Name).To(Equal("To Do"))
			Expect(lists[1].Name).To(Equal("Doing"))
			Expect(lists[2].Name).To(Equal("Done"))
		})

		g.It("should get the members of a board", func() {
			_, err := Board.Members()
			Expect(err).To(BeNil())
		})

		g.It("should get the cards in a board", func() {
			_, err := Board.Cards()
			Expect(err).To(BeNil())
		})

		g.It("should get the checklists in a board", func() {
			_, err := Board.Checklists()
			Expect(err).To(BeNil())
		})

		g.It("should get the membercards in a board", func() {
			// Retrieve member for "me"
			member, err := Client.Member("me")
			Expect(err).To(BeNil())
			// Get "my" cards on this board (probably none)
			_, err = Board.MemberCards(member.ID)
			Expect(err).To(BeNil())
		})

		g.It("should get the actions in a board", func() {
			_, err := Board.Actions()
			Expect(err).To(BeNil())
		})

		g.It("should add a list to a board", func() {
			list, err := Board.AddList(trello.List{
				Name: "go-test",
			})
			Expect(err).To(BeNil())
			Expect(list.Name).To(BeEquivalentTo("go-test"))
			Expect(list.IDBoard).To(Equal(Board.ID))
		})

		g.It("should get the labels in a board", func() {
			_, err := Board.Labels()
			Expect(err).To(BeNil())
		})

		g.It("should add a label to a board", func() {
			label, err := Board.AddLabel("go-testing", "orange")
			Expect(err).To(BeNil())
			Expect(label.Name).To(Equal("go-testing"))
			Expect(label.Color).To(Equal("orange"))
		})

		g.It("should duplicate (copy) the board", func() {
			new, err := Board.Duplicate("DUP-"+TestBoardName, true)
			Expect(err).To(BeNil())
			Expect(new.ID).NotTo(Equal(Board.ID))
			// and cleanup
			err = new.Delete()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = Board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
