package humans

import "tamagoshi/food"

//Human is an interface that defines the methods that a human can perform
// - Status() returns the current status of the human
// - Exercise(intensiy) return a string whit the result of the exercise
// - Sleep() returns a string describing the action of sleeping
// - Study() returns a string describing the action of studying
// - Eat(food) returns a string describing the action of eating
type Human interface {
	Status() string
	Exercise(intensity Intensity) string
	Sleep() string
	Study() string
	Eat(food food.Feeder) string
}
