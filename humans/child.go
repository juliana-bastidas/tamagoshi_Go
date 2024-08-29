package humans

import (
	"fmt"
	"tamagoshi/food"
)

// type Child struct {
// 	name       string
// 	hunger     int
// 	sleepyness int
// 	stamina    int
// }

//child toma como a un campo oculto a person , simplemente child extiende a person y hereda sus campos

type Child struct {
	*Person
}

// creates a new child
// func NewChild(name string) Child {
// 	return Child{
// 		name:       name,
// 		hunger:     100,
// 		sleepyness: rand.Intn(100) + 1,
// 		stamina:    rand.Intn(100) + 1,
// 	}
// }

func NewChild(name string) *Child {
	return &Child{
		Person: &Person{
			name:       name,
			hunger:     100,
			sleepyness: 10,
			stamina:    80,
		},
	}
}

//como child extiende a person deben tener los mismos metodos a menos que alguno cambie como eat por eso aca no se crean los metodos de exercise, study,sleep.

func (p *Child) Eat(f food.Feeder) string {
	effects := f.GetEffects()  // los efectos sobre la persona de lo que estuviera comiendo
	attrs := f.GetAttributes() // el nombre de lo que se estaba comiendo y el sabor
	p.stamina = normalizeRange(p.stamina + effects.Stamina*1.2)
	p.sleepyness = normalizeRange(p.sleepyness + effects.Sleepyness*1.1)
	p.hunger = normalizeRange(p.hunger + effects.Hunger*0.8)
	return fmt.Sprintf("I have eaten\n That %s, tasted %s", attrs.Name, attrs.Taste)
}
