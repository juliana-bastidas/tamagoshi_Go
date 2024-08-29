package screen

import "fmt"

// interfaz choice , contiene otras interfaces que se usaran para la seleccion de opciones
type choice interface {
	humanChoice | activityChoice | foodChoice | intensityChoice

	//humanChoice | activityChoice | intensiryChoice | foodChoice
}

//la funcion choicePrompt permite al usuario hacer una seleccion a partir de una lista de opciones.

func choicePrompt[T choice](question string, labelsMap map[T]string, unselected T, options ...T) T {
	userChoice := unselected
	validChoice := false

	for !validChoice {
		fmt.Println(question)
		//para todo lo que este en options, se va a imprimir el valor (person 1, child 2) y luego el string que viene de labelsMap
		//el for siempre va a imprimir la pregunta y las opciones
		for _, value := range options {
			fmt.Printf("%d. %s\n", value, labelsMap[value])
		}
		if _, err := fmt.Scan(&userChoice); err != nil || userChoice == unselected || userChoice > options[len(options)-1] {
			fmt.Printf("Invalid choice, please select a valid one\n")
		} else {
			validChoice = true
		}
	}
	return userChoice
}
