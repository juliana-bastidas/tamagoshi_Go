package screen

import (
	"fmt"
	"tamagoshi/humans"
)

//se le colocara una descripcion de la funcion para que aparezca despues

// HumanNamePrompt propts the suer for a human's name
// and returns it as a string value, it will keep asking
// until a valid name is provided
func HumanNamePrompt() string {
	name := ""
	for name == "" {
		fmt.Printf("What is the human's name? \n")
		// el _ ignora el valor de retorno
		if _, err := fmt.Scan(&name); err != nil {
			fmt.Scanln()
			fmt.Printf("Please enter a valid name\n")
		}
	}
	return name
}

type humanChoice int

const (
	HUMAN_UNSELECTED humanChoice = iota
	PERSON
	CHILD
)

var humanChoiceLabels = map[humanChoice]string{
	PERSON: "Person",
	CHILD:  "Child",
}

func HumanTypePrompt() humanChoice {
	return choicePrompt[humanChoice]("What did you want to create?", humanChoiceLabels, HUMAN_UNSELECTED, PERSON, CHILD)
}

//las funciones HumanNamePrompt y HumanTypePrompt son espeficicas para la creacion de humanos

type activityChoice int

const (
	ACTIVITY_UNSELECTED activityChoice = iota
	EXERCISE
	SLEEP
	STUDY
	EAT
	STATUS
	EXIT
)

var activityChoiceLabels = map[activityChoice]string{
	EXERCISE: "Exercise",
	SLEEP:    "Sleep",
	STUDY:    "Study",
	EAT:      "Eat",
	STATUS:   "Status",
	EXIT:     "Exit",
}

func ActivityPrompt() activityChoice {
	return choicePrompt[activityChoice]("What did you want me to do?", activityChoiceLabels, ACTIVITY_UNSELECTED, EXERCISE, SLEEP, STUDY, EAT, STATUS, EXIT)
}

type foodChoice int

const (
	FOOD_UNSELECTED foodChoice = iota
	DESSERT
	MEAL
	ENERGIZER
)

var foodChoiceLabels = map[foodChoice]string{
	DESSERT:   "Dessert",
	MEAL:      "Meal",
	ENERGIZER: "Energizer",
}

func FoodPrompt() foodChoice {
	return choicePrompt[foodChoice]("Which thing should I eat?", foodChoiceLabels, FOOD_UNSELECTED, DESSERT, MEAL, ENERGIZER)
}

type intensityChoice = humans.Intensity

const (
	UNSELECTED_INTENSITY intensityChoice = iota
	LOW_INTENSITY
	MEDIUM_INTENSITY
	HIGH_INTENSITY
)

var intenstyChoiceLabels = map[intensityChoice]string{
	humans.LOW_INTENSITY:    "Low",
	humans.MEDIUM_INTENSITY: "Medium",
	humans.HIGH_INTENSITY:   "High",
}

func IntensityPrompt() intensityChoice {
	return choicePrompt[intensityChoice]("What intensity should I exercise?", intenstyChoiceLabels, UNSELECTED_INTENSITY, humans.LOW_INTENSITY, humans.MEDIUM_INTENSITY, humans.HIGH_INTENSITY)
}
