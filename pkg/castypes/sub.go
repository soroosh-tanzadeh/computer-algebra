package castypes

import "fmt"

type Sub struct {
	Expr
	LHS Expr
	RHS Expr
}

func (s Sub) String() string {
	return fmt.Sprintf("( %v ) - ( %v )", s.LHS, s.RHS)
}

func (s Sub) Eval(arg float64) float64 {
	return s.LHS.Eval(arg) - s.RHS.Eval(arg)
}

func (a Sub) D() Expr {
	return Sub{LHS: a.LHS.D(), RHS: a.RHS.D()}
}
