package main

import (
	"fmt"
	"log"
	eval "main/ex7.14"
	"net/http"
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

func calculate(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("expr")
	expr, vars, err := parseAndCheck(s)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error parsing expression: %v\n", err)
		return
	}

	var env = make(eval.Env)

	for k := range vars {
		if _, ok := env[k]; !ok {
			value := r.URL.Query().Get(string(k))

			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "please provide the value for %s\n", k)
				return
			}

			n, err := strconv.ParseFloat(value, 64)

			if err != nil {
				fmt.Fprintf(w, "error parsing value %v\n%v\n", value, err)
				return
			}
			env[k] = n
		}
	}

	fmt.Fprintf(w, "%v = %v\n", expr, expr.Eval(env))

}

func main() {
	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
