package main

import (
	"fmt"

	"github.com/bvarberg/wwork/internal/emc"
)

func main() {
	fmt.Printf("The equilibrium moisture content is %v\n", emc.Simpson(70.0, 0.50))
}
