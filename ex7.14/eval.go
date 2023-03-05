package eval

import (
	"fmt"
	"math"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.operation {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)

	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.operation))
}

func (b binary) Eval(env Env) float64 {
	switch b.operation {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)

	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.operation))
}

func (c call) Eval(env Env) float64 {
	switch c.function {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.function))
}

func (m min) Eval(env Env) float64 {

	findMin := func(args []float64) float64 {

		if len(args) == 1 {
			return args[0]
		}

		min := args[0]

		for i := 1; i < len(args); i++ {
			if args[i] < min {
				min = args[i]
			}
		}

		return min
	}

	var args []float64

	for _, v := range m.args {
		args = append(args, v.Eval(env))
	}

	return findMin(args)

}
