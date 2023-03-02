package main

import (
	"algebra-system/pkg/castypes"
	"fmt"
)

func main() {
	e := castypes.Add{
		LHS: castypes.Const{Value: 4},
		RHS: castypes.Multiply{
			LHS: castypes.Const{Value: 6},
			RHS: castypes.Power{
				Base:  castypes.X{},
				Power: castypes.Const{Value: 3},
			},
		},
	}

	fmt.Println(e)
	fmt.Println(e.Eval(2))
	fmt.Println(e.D())
}
