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

func init() {
}
func TestBoard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Board tests", func() {
		var Board *trello.Board
		var TestBoardName string

		g.Before(func() {
			TestBoardName = fmt.Sprintf("GoTestTrello-%v", time.Now())

		})

		g.It("should create a board", func() {
			Board, err = Client.CreateBoard(TestBoardName)
			Expect(err).To(BeNil())
			Expect(Board.Name).To(Equal(TestBoardName))
		})

		g.It("should get a board", func() {
			Board, err = Client.Board(Board.ID)
			Expect(err).To(BeNil())
			Expect(Board.Name).To(Equal(TestBoardName))
		})

		g.It("should get the lists", func() {
			lists, err := Board.Lists()
			Expect(err).To(BeNil())
			Expect(lists[0].Name).To(Equal("To Do"))
			Expect(lists[1].Name).To(Equal("Doing"))
			Expect(lists[2].Name).To(Equal("Done"))
		})

		g.It("should change the background to red", func() {
			err = Board.SetBackground("red")
			Expect(err).To(BeNil())
			Expect(Board.Prefs.Background).To(Equal("red"))
		})

		g.It("should change the description to something", func() {
			err = Board.SetDescription("something")
			Expect(err).To(BeNil())
			Expect(Board.Desc).To(Equal("something"))
		})

	})

}
