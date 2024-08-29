package food

type Food struct {
	foodAtributes
	foodEffects
}

//foodAtributes y foodEffects estan privados por lo que no se pueden acceder desde otro paquete en cambio Food es publico y puede acceder
//esto quiere decir que se puede ingresar a Food pero no a foodAtributes y foodEffects, no puedo cambiar cosas de los metodos

type foodAtributes struct {
	Name  string
	Taste string
}

type foodEffects struct {
	Hunger     float64
	Sleepyness float64
	Stamina    float64
}

type Feeder interface {
	GetEffects() foodEffects
	GetAttributes() foodAtributes
}

func (f Food) GetAttributes() foodAtributes {
	return f.foodAtributes
}

func (f Food) GetEffects() foodEffects {
	return f.foodEffects
}
