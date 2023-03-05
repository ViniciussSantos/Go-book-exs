package main

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.operation, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.operation, b.y.String())
}

func (c call) String() string {

	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%s(", c.function)

	for i, v := range c.args {
		if i != 0 {
			fmt.Fprintf(buf, ", ")
		}
		fmt.Fprintf(buf, "%s", v.String())
	}
	fmt.Fprintf(buf, ")")

	return buf.String()
}
