package plugininterface

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
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

// FilterEventPlugin es la implementacion de la interface Plugin
type FilterEventPlugin struct {
	// Impl injeccion de la implementacion
	Impl FilterEvent
}

// Server retorna un servidor RPC, en este caso la estructura FilterEventRPCServer
func (p *FilterEventPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &FilterEventRPCServer{Impl: p.Impl}, nil
}

// Client regresa una implementacion de nuestra interface que se comunica por RFC, en este caso la estructura FilterEventRPCClient
func (p *FilterEventPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &FilterEventRPCClient{client: c}, nil
}
