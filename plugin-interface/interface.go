package plugininterface

import (
	"net/rpc"
)

// FilterEvent es la interface que deben implementar los plugins que filtran eventos
type FilterEvent interface {
	FilterEvent(event []byte) (filtered []byte, err error)
}

// FilterEventRPCClient implementacion del cliente RPC
type FilterEventRPCClient struct {
	client *rpc.Client
}

// FilterEvent filtra un evento
func (f *FilterEventRPCClient) FilterEvent(event []byte) (filtered []byte, err error) {
	err = f.client.Call("Plugin.FilterEvent", event, &filtered)
	return
}
