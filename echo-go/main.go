package main

import (
	"github.com/hashicorp/go-plugin"
	common "github.com/katzenpost/server/common-plugin"
)

type Echo struct{}

func (Echo) OnRequest(payload []byte, hasSURB bool) ([]byte, error) {
	return payload, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.Handshake,
		Plugins: map[string]plugin.Plugin{
			common.KaetzchenService: &common.KaetzchenPlugin{Impl: &Echo{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
