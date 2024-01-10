package context

import "LLD/LLD/STRATEGY_DESIGN_PATTERN/strategy"

type Snake struct {
	haslegsobj strategy.HasLegs
}

func NewSnake(obj strategy.HasLegs) *Snake {
	return &Snake{haslegsobj: obj}
}

func (h *Snake) PrintLegs() {
	h.haslegsobj.Print()
}
