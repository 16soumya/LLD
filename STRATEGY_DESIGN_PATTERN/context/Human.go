package context

import "LLD/LLD/STRATEGY_DESIGN_PATTERN/strategy"

type Human struct {
	haslegsobj strategy.HasLegs
}

func NewHuman(obj strategy.HasLegs) *Human {
	return &Human{haslegsobj: obj}
}

func (h *Human) PrintLegs() {
	h.haslegsobj.Print()
}
