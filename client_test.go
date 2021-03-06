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
	"net/http"
	"os"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("client tests", func() {
		var board *Board
		var testBoardName string

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Board-%v", time.Now().Unix())
		})

		g.It("should create a default client", func() {
			_, err = NewClient()
			Expect(err).To(BeNil())
		})

		g.It("NewCustomclient should create a custom client", func() {
			_, err = NewCustomClient(http.DefaultClient)
			Expect(err).To(BeNil())
		})

		g.It("NewAuthclient should create a client", func() {
			key := os.Getenv("API_KEY")
			token := os.Getenv("API_TOKEN")
			_, err = NewAuthClient(key, &token)
			Expect(err).To(BeNil())
		})

		g.It("should return version", func() {
			ver := client.Version()
			Expect(ver).To(Equal("1"))
		})

		g.It("should create a board", func() {
			board, err = client.CreateBoard(testBoardName)
			Expect(err).To(BeNil())
			Expect(board).NotTo(BeNil())
			Expect(board.Name).To(Equal(testBoardName))
		})

		g.After(func() {
			Expect(board).NotTo(BeNil())
			err = board.Delete()
			Expect(err).To(BeNil())
		})
	})

}
