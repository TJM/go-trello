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
	"log"
	"testing"
	"time"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestOrganization(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Organization tests", func() {
		var testBoardName string
		var board *Board
		var member *Member
		var Organization *Organization

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Organization-%v", time.Now().Unix())
			member, err = client.Member("me")
			if err != nil || member == nil {
				log.Fatal("ERROR Retrieving member (me): " + err.Error())
			}
			if len(member.IDOrganizations) < 1 {
				log.Fatalf("ERROR: Organization Tests require at least 1 organization!")
			}
		})

		g.It("should retrieve organization by Id", func() {
			Organization, err = client.Organization(member.IDOrganizations[0])
			Expect(err).To(BeNil())
		})

		g.It("should get members of an organization", func() {
			_, err = Organization.Members()
			Expect(err).To(BeNil())
		})

		g.It("should add a board to an organization", func() {
			board, err = Organization.AddBoard(testBoardName)
			Expect(err).To(BeNil())
		})

		g.It("should get a list of boards for an organization", func() {
			_, err = Organization.Boards()
			Expect(err).To(BeNil())
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			Expect(err).To(BeNil())
		})
	})

}
