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

func TestList(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("List tests", func() {
		var board *Board
		var list *List
		var testBoardName string

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-List-%v", time.Now().Unix())
			board, err = client.CreateBoard(testBoardName)
			Expect(err).To(BeNil())
			lists, err := board.Lists()
			Expect(err).To(BeNil())
			list = &lists[0]
		})

		g.It("should retrieve a list by ID", func() {
			_, err = client.List(list.ID)
			Expect(err).To(BeNil())
		})

		g.It("should retrieve actions for a list", func() {
			_, err = list.Actions()
			Expect(err).To(BeNil())
		})

		g.It("should add a card to a list", func() {
			c, err := list.AddCard(Card{
				Name: "Testing 123",
				Desc: "Does this thing work?",
			})
			Expect(err).To(BeNil())
			Expect(c).NotTo(BeNil())
		})

		g.It("should retrieve cards in a list", func() {
			_, err = list.Cards()
			Expect(err).To(BeNil())
		})

		g.It("should archive a list", func() {
			err = list.Archive(true)
			Expect(err).To(BeNil())
		})

		g.It("should unarchive a list", func() {
			err = list.Archive(false)
			Expect(err).To(BeNil())
		})

		g.It("should move a list", func() {
			err = list.Move("bottom")
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
