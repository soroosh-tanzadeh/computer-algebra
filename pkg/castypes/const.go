package castypes

import "fmt"

type Const struct {
	Expr
	Value float64
}

func (c Const) String() string {
	return fmt.Sprint(c.Value)
}

func (c Const) Eval(arg float64) float64 {
	return c.Value
}

func (c Const) D() Expr {
	return Const{Value: 0}
}
