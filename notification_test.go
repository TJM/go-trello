package trello

import (
	"testing"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestNotification(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Notification tests", func() {
		var member *Member
		var Notification *Notification

		g.Before(func() {
			member, err = client.Member("me")
			Expect(err).To(BeNil())
			notifications, err := member.Notifications()
			Expect(err).To(BeNil())
			if len(notifications) > 0 {
				Notification = &notifications[0]
			} else {
				t.Skip("No notifications to get")
			}
		})

		g.It("should retrieve a notification by id", func() {
			_, err = client.Notification(Notification.ID)
			Expect(err).To(BeNil())
		})

	})

}
