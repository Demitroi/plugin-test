package main

import (
	"os"

	plugininterface "github.com/Demitroi/plugin-test/plugin-interface"
	"github.com/Demitroi/plugin-test/plugins/inc20/inc20"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// handshakeConfig es utilizado como handshake entre el plugin y el host
// si falla se retorna un mensaje amigable para el usuario
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FILTER_EVENT_PLUGIN",
	MagicCookieValue: "INC10",
}

func main() {
	// Crear logger
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	// Crear incrementador
	incrementador := &inc20.Incrementar20{
		Logger: logger,
	}
	pluginFilterEvent := &plugininterface.FilterEventPlugin{
		Impl: incrementador,
	}
	// pluginMap es el mapa de plugins que se puede dispensar
	var pluginMap = map[string]plugin.Plugin{
		"filterevent": pluginFilterEvent,
	}
	// Servir el plugin
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
