package humans

//recibe una propiedad y nos da un mensaaje si esta por arriba del top o estar por debajo.

func evaluate(property float64, messageIfTop, messageIfBotton string) string {
	var message string
	switch property {
	case 100:
		message = messageIfTop
	case 0:
		message = messageIfBotton
	}
	return message
}

// para normalizar los rangos de las propiedades de los humanos

func normalizeRange(value float64) float64 {
	if value > 100 {
		return 100
	}
	if value < 0 {
		return 0
	}
	return value
}
