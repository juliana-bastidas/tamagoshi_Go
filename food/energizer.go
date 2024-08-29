package food

type Energizer struct {
	*Food
}

func NewEnergizer(name, taste string) *Energizer {
	return &Energizer{
		Food: &Food{
			foodAtributes: foodAtributes{
				Name:  name,
				Taste: taste,
			},
			foodEffects: foodEffects{
				Hunger:     -10,
				Sleepyness: -50,
				Stamina:    50,
			},
		},
	}
}
