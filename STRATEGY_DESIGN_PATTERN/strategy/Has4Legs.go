package strategy

import "fmt"

type Has4Legs struct{
}

func (h *Has4Legs) Print(){
	fmt.Println("Has 4 Legs")
}