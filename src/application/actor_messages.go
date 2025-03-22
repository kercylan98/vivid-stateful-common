package application

import "github.com/kercylan98/vivid-stateful-common/src/application/service"

// Request：XXXAsk
// Response：XXXReply
// Notify：XXXNotify

type (
	registerServiceAsk struct {
		Provider service.Provider
	}

	registerServiceReply struct {
	}
)

type (
	launchAsk int8

	launchReply int8
)
