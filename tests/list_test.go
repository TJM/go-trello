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

func TestList(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("List tests", func() {
		var Board *trello.Board
		var List *trello.List
		var TestBoardName string

		g.Before(func() {
			TestBoardName = fmt.Sprintf("GoTestTrello-List-%v", time.Now())
			Board, err = Client.CreateBoard(TestBoardName)
			Expect(err).To(BeNil())
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			List = &lists[0]
		})

		g.It("should retrieve a list by ID", func() {
			_, err = Client.List(List.ID)
			Expect(err).To(BeNil())
		})

		g.It("should retrieve cards in a list", func() {
			_, err = List.Cards()
			Expect(err).To(BeNil())
		})

		g.It("should retrieve actions for a list", func() {
			_, err = List.Actions()
			Expect(err).To(BeNil())
		})

		g.It("should add a card to a list", func() {
			_, err = List.AddCard(trello.Card{
				Name: "Testing 123",
				Desc: "Does this thing work?",
			})
			Expect(err).To(BeNil())
		})

		g.It("should archive a list", func() {
			err = List.Archive(true)
			Expect(err).To(BeNil())
		})

		g.It("should unarchive a list", func() {
			err = List.Archive(false)
			Expect(err).To(BeNil())
		})

		g.It("should move a list", func() {
			_, err = List.Move("bottom")
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = Board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
