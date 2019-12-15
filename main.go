package main

import (
	"fmt"
	"os"
	"os/exec"

	plugininterface "github.com/Demitroi/plugin-test/plugin-interface"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// handshakeConfig es utilizado como handshake entre el plugin y el host
// si falla se retorna un mensaje amigable para el usuario
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FILTER_EVENT_PLUGIN",
	MagicCookieValue: "INC10",
}

var (
	pluginPath = "./plugin/filterevent.plug"
	// eventFiltering es la interface a la que se va a llamar
	eventFiltering plugininterface.FilterEvent
)

// defaultFilterEvent se usa en caso de que no este disponible el plugin
type defaultFilterEvent struct{}

// FilterEvent no hace nada, simplemente retorna el evento tal cual lo recibio
func (*defaultFilterEvent) FilterEvent(event []byte) (filtered []byte, err error) {
	return event, nil
}

func main() {
	// Validar si existe el archivo
	_, err := os.Stat(pluginPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Si no existe no se carga
			// Se utiliza el filtrador de eventos por default
			eventFiltering = &defaultFilterEvent{}
		} else {
			// Si ocurre algun tipo de error panic!
			panic(err)
		}
	} else {
		// Si existe se carga el plugin
		// Crear mapa de plugins
		var pluginMap = map[string]plugin.Plugin{
			"filterevent": &plugininterface.FilterEventPlugin{},
		}
		// Crear logger
		logger := hclog.New(&hclog.LoggerOptions{
			Name:   "plugin",
			Output: os.Stdout,
			Level:  hclog.Debug,
		})
		// Iniciar el cliente de plugin
		pluginClient := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(pluginPath),
			Logger:          logger,
		})
		defer pluginClient.Kill()
		// Conectarse con el cliente rpc
		rpcClient, err := pluginClient.Client()
		if err != nil {
			panic(err)
		}
		// Solicitar el plugin
		raw, err := rpcClient.Dispense("filterevent")
		if err != nil {
			panic(err)
		}
		// Hacer la asercion para utilizar el plugin
		eventFiltering = raw.(plugininterface.FilterEvent)
	}
	// En este punto ya se tiene eventFiltering listo, ya sea el default o algun plugin
	// Hacer un ejemplo filtrando un evento
	event := []byte(`<articulo nombre="n64"><precio1>10</precio1><precio2>15.25</precio2><precio3>20</precio3></articulo>`)
	eventFiltered, err := eventFiltering.FilterEvent(event)
	if err != nil {
		err = fmt.Errorf("Ha ocurrido un error al filtrar evento: %v", err)
		panic(err)
	}
	fmt.Printf("evento original:\n%s\nevento filtrado:\n%s\n", event, eventFiltered)
}
