package food

type Dessert struct {
	*Food
}

func NewDessert(name, taste string) *Dessert {
	return &Dessert{
		Food: &Food{
			foodAtributes: foodAtributes{
				Name:  name,
				Taste: taste,
			},
			foodEffects: foodEffects{
				Hunger:     -20,
				Sleepyness: 0,
				Stamina:    10,
			},
		},
	}
}
