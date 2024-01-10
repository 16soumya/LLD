package context

import "LLD/LLD/STRATEGY_DESIGN_PATTERN/strategy"

type Tiger struct {
	haslegsobj strategy.HasLegs
}

func NewTiger(obj strategy.HasLegs) *Tiger {
	return &Tiger{haslegsobj: obj}
}

func (h *Tiger) PrintLegs() {
	h.haslegsobj.Print()
}
