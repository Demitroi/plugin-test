package inc20_test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Demitroi/plugin-test/plugins/inc20/inc20"
	"github.com/hashicorp/go-hclog"
)

func TestFilterEvent(t *testing.T) {
	// Crear logger
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	// Crear incrementador
	incrementador := &inc20.Incrementar20{Logger: logger}
	// Abrir el evento de prueba
	eventByte, err := ioutil.ReadFile("testdata/example.xml")
	if err != nil {
		t.Fatalf("Falla al abrir evento de prueba: %v", err)
	}
	// Filtrar el evento
	eventFilteredByte, err := incrementador.FilterEvent(eventByte)
	if err != nil {
		t.Error(err)
		return
	}
	// Abrir el evento esperado
	eventExpectedByte, err := ioutil.ReadFile("testdata/example_expected.xml")
	if err != nil {
		t.Fatalf("Falla al abrir el evento esperado: %v", err)
	}
	// Decodificar el evento esperado y el filtrado
	var eventFiltered, eventExpected inc20.EventoXML
	err = xml.Unmarshal(eventFilteredByte, &eventFiltered)
	if err != nil {
		t.Fatalf("Falla al decodificar evento filtrado: %b", err)
	}
	err = xml.Unmarshal(eventExpectedByte, &eventExpected)
	if err != nil {
		t.Fatalf("Falla al decodificar evento esperado: %b", err)
	}
	// Comparar el evento esperado con el filtrado
	if eventFiltered != eventExpected {
		t.Error("El evento filtrado no es igual al esperado")
		return
	}
}

func TestFilterEventError(t *testing.T) {
	// Crear logger
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	// Crear incrementador
	incrementador := inc20.Incrementar20{Logger: logger}
	// Enviar un evento invalido para que retorne error
	eventFilteredByte, err := incrementador.FilterEvent([]byte(`:(`))
	if err == nil {
		t.Error("Error no debe ser nil")
		return
	}
	if eventFilteredByte != nil {
		t.Error("eventFilteredByte debe ser nil")
	}
}
