package main

import (
	"fmt"
	"sync"

	factory "github.com/jbiers/go-design-patterns/creational/factory"
	singleton "github.com/jbiers/go-design-patterns/creational/singleton"
)

var lock = &sync.Mutex{}

func main() {
	// FACTORY PATTERN
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	fmt.Printf("Gun: %s\n", ak47.GetName())
	fmt.Printf("Power: %d\n", ak47.GetPower())

	fmt.Println()

	fmt.Printf("Gun: %s\n", musket.GetName())
	fmt.Printf("Power: %d\n", musket.GetPower())

	// ABSTRACT FACTORY PATTERN

	// BUILDER PATTERN

	// SINGLETON PATTERN
	for i := 0; i < 30; i++ {
		go singleton.GetDatabase()
	}

	fmt.Scanln()
}
