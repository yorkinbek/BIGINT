package main

import (
	bigint "BIGINT/bigInt"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("-113253253")

	if err != nil {
		panic(err)
	}
	b, err := bigint.NewInt("0000011231231231")

	if err != nil {
		panic(err)
	}

	err = a.Set("2")
	if err != nil {
		panic(err)
	}
	x := bigint.Bigint{Value: "+9999999999"}
	y := bigint.Bigint{Value: "-1111119999999"}

	m := bigint.Bigint{Value: "+888999998899"}
	n := bigint.Bigint{Value: "-9999999999"}

	i := bigint.Bigint{Value: "12325554545"}
	j := bigint.Bigint{Value: "1234566789"}
	c := bigint.Add(x, y)
	d := bigint.Sub(m, n)
	e := bigint.Multiply(i, j)
	f := bigint.Mod(i, j)
	fmt.Println(" after Set :", a)
	fmt.Println(b)

	fmt.Println("a + b = ", c)
	fmt.Println("a - b = ", d)
	fmt.Println("a * b = ", e)
	fmt.Println("a mod b = ", f)
}
