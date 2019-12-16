package inc20

import (
	"encoding/xml"

	"github.com/hashicorp/go-hclog"
)

// EventoXML estructura usada para el xml
type EventoXML struct {
	Nombre  string   `xml:"nombre,attr"`
	Precio1 float64  `xml:"precio1"`
	Precio2 float64  `xml:"precio2"`
	Precio3 float64  `xml:"precio3"`
	XMLName xml.Name `xml:"articulo"`
}

// Incrementar20 es una implementacion real de la interface de plugin
// Incrementa un 20% los precios
type Incrementar20 struct {
	Logger hclog.Logger
}

// FilterEvent implementa el metodo para incrementar 20% los precios
func (i *Incrementar20) FilterEvent(event []byte) (filtered []byte, err error) {
	// Hacer unmarshal al evento original
	var eventoOriginal EventoXML
	err = xml.Unmarshal(event, &eventoOriginal)
	if err != nil {
		return nil, err
	}
	// Incrementar los precios un 20%
	eventoOriginal.Precio1 += eventoOriginal.Precio1 * 0.20
	eventoOriginal.Precio2 += eventoOriginal.Precio2 * 0.20
	eventoOriginal.Precio3 += eventoOriginal.Precio3 * 0.20
	// Volver a formar el xml
	return xml.Marshal(eventoOriginal)
}
