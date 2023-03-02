package castypes

import "fmt"

type Multiply struct {
	Expr
	LHS Expr
	RHS Expr
}

func (m Multiply) String() string {
	return fmt.Sprintf("( %v ) * ( %v )", m.LHS, m.RHS)
}

func (m Multiply) Eval(arg float64) float64 {
	return m.LHS.Eval(arg) * m.RHS.Eval(arg)
}

// Product rule: (fg)' = f'g + fg'
func (m Multiply) D() Expr {
	return Add{
		LHS: Multiply{
			LHS: m.LHS.D(),
			RHS: m.RHS,
		},
		RHS: Multiply{
			LHS: m.LHS,
			RHS: m.RHS.D(),
		},
	}
}
