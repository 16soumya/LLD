package main

import (
	"LLD/LLD/STRATEGY_DESIGN_PATTERN/context"
	"LLD/LLD/STRATEGY_DESIGN_PATTERN/strategy"
)

func main() {
	has0legs := &strategy.Has0Legs{}
	has2legs := &strategy.Has2Legs{}
	has4legs := &strategy.Has4Legs{}

	var animal context.Animal

	animal = context.NewHuman(has2legs)
	animal.PrintLegs()

	animal = context.NewTiger(has4legs)
	animal.PrintLegs()

	animal = context.NewSnake(has0legs)
	animal.PrintLegs()
}
