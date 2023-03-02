package castypes

import "fmt"

type Add struct {
	Expr
	LHS Expr
	RHS Expr
}

func (a Add) String() string {
	return fmt.Sprintf("( %v ) + ( %v )", a.LHS, a.RHS)
}

func (a Add) Eval(arg float64) float64 {
	return a.LHS.Eval(arg) + a.RHS.Eval(arg)
}

func (a Add) D() Expr {
	return Add{LHS: a.LHS.D(), RHS: a.RHS.D()}
}
