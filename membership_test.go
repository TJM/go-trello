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

func TestMembership(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Membership tests", func() {
		var board *Board
		var member *Member
		var membership *Membership
		var testBoardName string

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Membership-%v", time.Now().Unix())
			board, err = client.CreateBoard(testBoardName)
			Expect(err).To(BeNil())
			Expect(board).NotTo(BeNil())
			member, err = client.Member("trello")
			Expect(err).To(BeNil())
			Expect(member).NotTo(BeNil())
			err = board.AddMember(member, "")
			Expect(err).To(BeNil())
			membership, err = board.GetMembershipForMember(member)
			Expect(err).To(BeNil())
			Expect(membership).NotTo(BeNil())
		})

		// Putting a board test here due to prerequisites
		g.It("should get a membership by id", func() {
			ms, err := board.GetMembership(membership.ID)
			Expect(err).To(BeNil())
			Expect(membership).NotTo(BeNil())
			Expect(ms.ID).To(Equal(membership.ID))
		})

		g.It("should update fields membership for a member on a board", func() {
			err = membership.Update("", "fullName, username, bio")
			Expect(err).To(BeNil())
			// TODO: Check to be sure trello membership has changed?
			// Expect(len(membershipResponse.Memberships)).To(BeNumerically(">", 1))
		})

		g.It("should update type of membership for a member on a board", func() {
			err = membership.Update("admin", "")
			Expect(err).To(BeNil())
			// TODO: Check to be sure trello membership has changed?
			// Expect(len(membershipResponse.Memberships)).To(BeNumerically(">", 1))
		})

		// Keep this test LAST for obvious reasons
		g.After(func() {
			err = board.Delete()
			Expect(err).To(BeNil())
		})

	})

}
