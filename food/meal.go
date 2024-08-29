package food

type Meal struct {
	*Food
}

func NewMeal(name, taste string) *Meal {
	return &Meal{
		Food: &Food{
			foodAtributes: foodAtributes{
				Name:  name,
				Taste: taste,
			},
			foodEffects: foodEffects{
				Hunger:     -50,
				Sleepyness: 10,
				Stamina:    25,
			},
		},
	}
}
