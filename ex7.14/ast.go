package main

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

type literal float64

type unary struct {
	operation rune
	x         Expr
}

type binary struct {
	operation rune
	x, y      Expr
}

type call struct {
	function string
	args     []Expr
}
