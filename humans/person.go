package humans

import (
	"fmt"
	"tamagoshi/food"
)

// misma estructura tanto para persona como para niño
type Person struct {
	name       string
	hunger     float64
	sleepyness float64
	stamina    float64
}

//creates a new persona

//como siempre vamos a trabajar con la misma persona, no vamos a crear y eliminar podemos dejar el puntero, siempre vamos a estar editando el mismo objeto en memoria

func NewPerson(name string) *Person {
	return &Person{name: name, hunger: 80, sleepyness: 10, stamina: 80}
}

func (p Person) Status() string {
	message := fmt.Sprintf(" My names is %s, this is my status: \n Hunger: %f \n Sleepyness: %f \n Stamina: %f \n", p.name, p.hunger, p.sleepyness, p.stamina)
	message += evaluate(p.hunger, "I'm very hungry\n", "I'm satiated\n")
	message += evaluate(p.sleepyness, "I need to sleep\n", "I'm totally rested\n")
	message += evaluate(p.stamina, "I have a lot of energy\n", "I'm tired\n")
	return message

}

func (p *Person) Sleep() string {
	p.stamina = normalizeRange(100)
	p.sleepyness = normalizeRange(0)
	p.hunger = normalizeRange(p.hunger + 20)
	return "I have slept"
}

func (p *Person) Study() string {
	p.stamina = normalizeRange(p.stamina - 10)
	p.sleepyness = normalizeRange(p.sleepyness + 30)
	p.hunger = normalizeRange(p.hunger + 25)
	return "I have studied"
}

func (p *Person) Exercise(intensity Intensity) string {
	message := ""
	switch intensity {
	case LOW_INTENSITY:
		p.stamina = normalizeRange(p.stamina - 10)
		p.hunger = normalizeRange(p.hunger + 10)
		message = "I have exercise at low intensity"
	case MEDIUM_INTENSITY:
		p.stamina = normalizeRange(p.stamina - 25)
		p.hunger = normalizeRange(p.hunger + 30)
		message = "I have exercise at medium intensity"
	case HIGH_INTENSITY:
		p.stamina = normalizeRange(p.stamina - 50)
		p.hunger = normalizeRange(p.hunger + 60)
		message = "I have exercise at high intensity"
	}
	return message
}

func (p *Person) Eat(f food.Feeder) string {
	effects := f.GetEffects()  // los efectos sobre la persona de lo que estuviera comiendo
	attrs := f.GetAttributes() // el nombre de lo que se estaba comiendo y el sabor
	p.stamina = normalizeRange(p.stamina + effects.Stamina)
	p.sleepyness = normalizeRange(p.sleepyness + effects.Sleepyness)
	p.hunger = normalizeRange(p.hunger + effects.Hunger)
	return fmt.Sprintf("I have eaten\n That %s, tasted %s", attrs.Name, attrs.Taste)
}
