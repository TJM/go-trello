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
	"net/http"
	"os"
	"testing"

	"github.com/TJM/go-trello"
	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Client tests", func() {

		g.It("NewClient should create a default client", func() {
			_, err = trello.NewClient()
			Expect(err).To(BeNil())
		})

		g.It("NewCustomClient should create a custom client", func() {
			_, err = trello.NewCustomClient(http.DefaultClient)
			Expect(err).To(BeNil())
		})

		g.It("NewAuthClient should create a client", func() {
			key := os.Getenv("API_KEY")
			token := os.Getenv("API_TOKEN")
			_, err = trello.NewAuthClient(key, &token)
			Expect(err).To(BeNil())
		})

		// NOTE: Other methods will be tested as part of other tests
	})

}
