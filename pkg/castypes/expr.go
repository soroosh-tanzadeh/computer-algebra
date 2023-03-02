package castypes

type Expr interface {
	String() string
	Eval(float64) float64
	D() Expr
}
