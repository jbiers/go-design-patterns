package main

import (
	"fmt"
)

func main() {
	ak47, _ := factory.getGun("ak47")
	musket, _ := factory.getGun("musket")

	fmt.Printf("Gun: %s", ak47.getName())
	fmt.Printf("Power: %d", ak47.getPower())

	fmt.Println()

	fmt.Printf("Gun: %s", musket.getName())
	fmt.Printf("Power: %d", musket.getPower())
}
