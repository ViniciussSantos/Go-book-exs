package main

import (
	"bufio"
	"fmt"
	"log"
	eval "main/ex7.14"
	"os"
	"strconv"
)

func parseAndCheck(s string) (eval.Expr, map[eval.Var]bool, error) {
	if s == "" {
		return nil, nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, vars, err
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var env = make(eval.Env)

	fmt.Printf("Enter the expr:")
	input.Scan()
	s := input.Text()
	expr, vars, err := parseAndCheck(s)

	if err != nil {
		log.Fatalf("Errorparsing %s\n%v\n", s, err)
	}

	for k := range vars {
		if _, ok := env[k]; !ok {
			fmt.Printf("Enter the value for %v:", k)
			input.Scan()
			digit := input.Text()

			n, err := strconv.ParseFloat(digit, 64)

			if err != nil {
				fmt.Printf("error parsing value %v\n%v\n", digit, err)
			}
			env[k] = n
		}
	}

	fmt.Printf("%v = %v\n", expr, expr.Eval(env))

}
