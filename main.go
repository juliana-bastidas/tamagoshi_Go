package main

import (
	"fmt"
	"tamagoshi/food"
	"tamagoshi/humans"
	"tamagoshi/screen"
)

func main() {

	// Functionality example
	//person := humans.NewPerson("Wences")
	//fmt.Println("person", person)
	//child := humans.NewChild("Wences Jr")
	//fmt.Println("child", child)
	// dessert := food.NewDessert("pay","sweet")
	// meal := food.NewMeal("meat","tasty")
	// fmt.Println(person.Status())
	// fmt.Println(person.Exercise(humans.HIGH_INTENSITY))
	// fmt.Println(person.Status())
	// fmt.Println(person.Eat(meal))
	// fmt.Println(person.Eat(dessert))
	// fmt.Println(person.Status())
	// fmt.Println(child.Status())
	// fmt.Println(child.Eat(meal))
	// fmt.Println(child.Eat(dessert))
	// fmt.Println(child.Status())

	// PRUEBAS
	var human humans.Human
	humanType := screen.HumanTypePrompt()
	name := screen.HumanNamePrompt()

	switch humanType {
	case screen.PERSON:
		human = humans.NewPerson(name)
	case screen.CHILD:
		human = humans.NewChild(name)
	}

	fmt.Println(human.Status())

	selectedActivity := screen.ACTIVITY_UNSELECTED
	for selectedActivity != screen.EXIT {
		selectedActivity = screen.ActivityPrompt()
		switch selectedActivity {
		case screen.EXERCISE:
			selectedIntensity := humans.UNSELECTED_INTENSITY
			selectedIntensity = screen.IntensityPrompt()
			fmt.Println(human.Exercise(selectedIntensity))
		case screen.SLEEP:
			fmt.Println(human.Sleep())
		case screen.STUDY:
			fmt.Println(human.Study())
		case screen.EAT:
			var selectedFood food.Feeder
			foodChoice := screen.FoodPrompt()
			switch foodChoice {
			case screen.DESSERT:
				selectedFood = food.NewDessert("pay", "tasty")
			case screen.MEAL:
				selectedFood = food.NewMeal("steak", "smoked")
			case screen.ENERGIZER:
				selectedFood = food.NewEnergizer("boost", "energizer")
			}
			fmt.Println(human.Eat(selectedFood))
		case screen.STATUS:
			fmt.Println(human.Status())

		}
	}

}
