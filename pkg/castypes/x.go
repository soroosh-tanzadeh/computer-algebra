package castypes

type X struct {
	Expr
}

func (x X) String() string {
	return "X"
}

func (x X) Eval(arg float64) float64 {
	return arg
}
