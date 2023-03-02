package castypes

import "fmt"

type Divide struct {
	Expr
	LHS Expr
	RHS Expr
}

func (d Divide) String() string {
	return fmt.Sprintf("( %v ) / ( %v )", d.LHS, d.RHS)
}

func (d Divide) Eval(arg float64) float64 {
	return d.LHS.Eval(arg) / d.RHS.Eval(arg)
}

// D Fraction rule: (f/g)' = (f'g - fg')/(g^2)
func (d Divide) D() Expr {
	return Divide{
		LHS: Sub{
			LHS: Multiply{
				LHS: d.LHS.D(),
				RHS: d.RHS,
			},
			RHS: Multiply{
				LHS: d.LHS,
				RHS: d.RHS.D(),
			},
		},
		RHS: Power{
			Base:  d.RHS,
			Power: Const{Value: 2},
		},
	}
}
