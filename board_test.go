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

func TestBoard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Board tests", func() {
		var board *Board
		var label *Label
		var testBoardName string

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Board-%v", time.Now().Unix())

		})

		g.It("should create a board", func() {
			board, err = client.CreateBoard(testBoardName)
			Expect(err).To(BeNil())
			Expect(board).NotTo(BeNil())
			Expect(board.Name).To(Equal(testBoardName))
		})

		g.It("should get a board by ID", func() {
			board, err = client.Board(board.ID)
			Expect(err).To(BeNil())
			Expect(board).NotTo(BeNil())
			Expect(board.Name).To(Equal(testBoardName))
		})

		g.It("should change the board background to red", func() {
			err = board.SetBackground("red")
			Expect(err).To(BeNil())
			Expect(board.Prefs.Background).To(Equal("red"))
		})

		g.It("should change the description to something", func() {
			err = board.SetDescription("something")
			Expect(err).To(BeNil())
			Expect(board.Desc).To(Equal("something"))
		})

		g.It("should get the lists in a board", func() {
			lists, err := board.Lists()
			Expect(err).To(BeNil())
			// This part is somewhat dangerous if Trello changes the default Board Template
			Expect(lists[0].Name).To(Equal("To Do"))
			Expect(lists[1].Name).To(Equal("Doing"))
			Expect(lists[2].Name).To(Equal("Done"))
		})

		g.It("should get the members of a board", func() {
			_, err := board.Members()
			Expect(err).To(BeNil())
		})

		g.It("should get all the actions in a board", func() {
			_, err := board.Actions()
			Expect(err).To(BeNil())
		})

		g.It("should get filtered actions in a board", func() {
			arg := NewArgument("filter", "addMemberToCard")
			_, err := board.Actions(arg)
			Expect(err).To(BeNil())
		})

		g.It("should add a list to a board", func() {
			list, err := board.AddList(List{
				Name: "go-test",
			})
			Expect(err).To(BeNil())
			Expect(list.Name).To(BeEquivalentTo("go-test"))
			Expect(list.IDBoard).To(Equal(board.ID))
		})

		g.It("should get the labels in a board", func() {
			_, err := board.Labels()
			Expect(err).To(BeNil())
		})

		g.It("should add a label to a board", func() {
			label, err = board.AddLabel("go-testing", "orange")
			Expect(err).To(BeNil())
			Expect(label.Name).To(Equal("go-testing"))
			Expect(label.Color).To(Equal("orange"))
		})

		g.It("should update a label name", func() {
			err = label.SetName("super-go-testing")
			Expect(err).To(BeNil())
		})

		g.It("should update a label color", func() {
			err = label.SetColor("purple")
			Expect(err).To(BeNil())
		})

		g.It("should delete a label", func() {
			err = label.Delete()
			Expect(err).To(BeNil())
		})

		g.It("should duplicate (copy) the board", func() {
			new, err := board.Duplicate("DUP-"+testBoardName, true)
			Expect(err).To(BeNil())
			Expect(new.ID).NotTo(Equal(board.ID))
			// and cleanup
			err = new.Delete()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
