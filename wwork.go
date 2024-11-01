package main

import (
	"fmt"

	"github.com/bvarberg/wwork/internal/emc"
)

func main() {
	fmt.Printf("The equilibrium moisture content is %v\n", emc.Equilibrium(70.0, 50.0))
}
