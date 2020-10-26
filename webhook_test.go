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
	"math/rand"
	"os"
	"testing"

	goblin "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestWebhook(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Webhook tests", func() {
		var webhook *Webhook
		var token string
		var callbackURL string

		g.Before(func() {
			token = os.Getenv("API_TOKEN")
			Expect(token).NotTo(BeEmpty())
			callbackURL = fmt.Sprintf("https://www.google.com/?q=go-trello-test-%v", rand.Intn(65536))
		})

		g.It("should create a webhook", func() {
			member, err := client.Member("me")
			Expect(err).To(BeNil())
			webhook, err = client.CreateWebhook(Webhook{
				Description: "Testing Webhook",
				CallbackURL: callbackURL,
				IDModel:     member.ID,
			})
			Expect(err).To(BeNil())
		})

		g.It("should deactivate a webhook", func() {
			err = webhook.SetActive(false)
			Expect(err).To(BeNil())
			Expect(webhook.Active).To(BeFalse())
		})

		g.It("should set a callback url for a webhook", func() {
			err = webhook.SetCallbackURL(callbackURL + "-test")
			Expect(err).To(BeNil())
			Expect(webhook.CallbackURL).To(Equal(callbackURL + "-test"))
		})

		g.It("should set a callback url for a webhook", func() {
			err = webhook.SetDescription("Go-Trello Testing")
			Expect(err).To(BeNil())
			Expect(webhook.Description).To(Equal("Go-Trello Testing"))
		})

		g.It("should set a idModel for a webhook", func() {
			err = webhook.SetIDModel(webhook.IDModel) // kinda cheating
			Expect(err).To(BeNil())
		})

		g.It("should retrieve a webhook by id", func() {
			_, err = client.Webhook(webhook.ID)
			Expect(err).To(BeNil())
		})

		g.It("should get webhooks for a token", func() {
			_, err = client.Webhooks(token)
		})

		// Destructive Action - Should be last
		g.It("should delete a webhook", func() {
			err = webhook.Delete()
			Expect(err).To(BeNil())
		})

	})

}
