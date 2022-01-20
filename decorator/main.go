package main

import (
	"fmt"

	"github.com/kilianstallz/go-patterns/decorator/drinks"
)

func main() {

	es := drinks.NewEspresso()
	print(es)

	var b drinks.Beverage
	b = drinks.NewEspresso()
	b = drinks.NewWhip(b)
	print(b)
}

func print(b drinks.Beverage) {
	fmt.Println(b.GetDescription(), "= \u20ac", b.Cost())
}
