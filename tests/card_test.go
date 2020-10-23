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
	"os"
	"testing"
	"time"

	"github.com/TJM/go-trello"
	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

//var Client *trello.Client
var Card *trello.Board
var TestCardName string

func init() {
	key := os.Getenv("API_KEY")
	token := os.Getenv("API_TOKEN")
	Client, err = trello.NewAuthClient(key, &token)
	TestCardName = fmt.Sprintf("GoTestTrelloCard-%v", time.Now())
}
func TestCard(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Card tests", func() {

		g.It("should create a card", func() {
			fmt.Printf(TestCardName)
			Expect(err).To(BeNil())
		})

		// g.It("should get a card using two different methods", func() {
		// 	card, err := Board.Card("56cdb3e0f7f4609c2b6f15e4")
		// 	Expect(err).To(BeNil())
		// 	Expect(card.Name).To(Equal("a card"))
		// 	sameCard, err := Client.Card("8sB7wile")
		// 	Expect(err).To(BeNil())
		// 	Expect(sameCard.Name).To(Equal("a card"))
		// 	Expect(sameCard.Desc).To(Equal(card.Desc))
		// })
	})

}
