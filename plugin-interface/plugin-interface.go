package plugininterface

// FilterEvent es la interface que deben implementar los plugins que filtran eventos
type FilterEvent interface {
	FilterEvent(event []byte) (filtered []byte, err error)
}
