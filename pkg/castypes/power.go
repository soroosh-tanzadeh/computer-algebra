package castypes

import (
	"fmt"
	"math"
)

type Power struct {
	Expr
	Base  Expr
	Power Expr
}

func (p Power) String() string {
	if (p.Power == Const{Value: 1}) {
		return fmt.Sprintf("%v", p.Base)
	}
	return fmt.Sprintf("(%v^ (%v))", p.Base, p.Power)
}

func (p Power) Eval(args float64) float64 {
	var baseEval = p.Base.Eval(args)
	var powerEval = p.Power.Eval(args)

	return math.Pow(baseEval, powerEval)
}

// Sin(cos(x)) = cos(cos(x)) * -sin(x)
// D power rule: (x^b)' = bx^(b-1)
func (p Power) D() Expr {
	return Power{
		Base: Multiply{
			LHS: p.Power,
			RHS: p.Base,
		},
		Power: Sub{
			LHS: p.Power,
			RHS: Const{Value: 1},
		},
	}
}
