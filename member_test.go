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

func TestMember(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Member tests", func() {
		var testBoardName string
		var member *Member

		g.Before(func() {
			testBoardName = fmt.Sprintf("GoTestTrello-Member-%v", time.Now().Unix())
		})

		g.It("should retrieve a member (me)", func() {
			member, err = client.Member("me")
			Expect(err).To(BeNil())
		})

		g.It("should retrieve a member (@trello)", func() {
			_, err = client.Member("trello")
			Expect(err).To(BeNil())
		})

		g.It("should retrieve boards for a member", func() {
			_, err = member.Boards()
			Expect(err).To(BeNil())
		})

		g.It("should add a board to a member", func() {
			b, err := member.AddBoard(testBoardName)
			Expect(err).To(BeNil())
			// Cleanup
			err = b.Delete()
			Expect(err).To(BeNil())
		})

		g.It("should get notifications for a member", func() {
			_, err = member.Notifications()
			Expect(err).To(BeNil())
		})

		g.It("should return avatar url for a member", func() {
			url := member.AvatarURL()
			Expect(url).To(ContainSubstring(member.AvatarHash))
		})

	})

}
