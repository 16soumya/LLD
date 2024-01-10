package strategy

import "fmt"

type Has0Legs struct{
}

func (h *Has0Legs) Print(){
	fmt.Println("Has 0 Legs")
}