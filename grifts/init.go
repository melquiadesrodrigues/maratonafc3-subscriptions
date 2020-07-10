package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/melquiadesr/maratonafc3_subscriptions/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
