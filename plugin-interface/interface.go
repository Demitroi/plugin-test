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

// FilterEvent llama al servidor rpc para filtra un evento
func (f *FilterEventRPCClient) FilterEvent(event []byte) (filtered []byte, err error) {
	err = f.client.Call("Plugin.FilterEvent", event, &filtered)
	return
}

// FilterEventRPCServer es la implementacion de rpc servidor
type FilterEventRPCServer struct {
	// Impl es la implementacion real de FilterEvent
	Impl FilterEvent
}

// FilterEvent recibe un evento del servidor y lo filtra
func (f *FilterEventRPCServer) FilterEvent(event []byte, filtered *[]byte) error {
	v, err := f.Impl.FilterEvent(event)
	*filtered = v
	return err
}
