package strategy

import "fmt"

type Has2Legs struct{
}

func (h *Has2Legs) Print(){
	fmt.Println("Has 2 Legs")
}