package service

import "github.com/kercylan98/vivid/src/vivid"

type Context interface {
	vivid.ActorContext

	Select(serviceType string) (vivid.ActorRef, error)

	SelectAll(serviceType string) []vivid.ActorRef
}
