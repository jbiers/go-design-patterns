package factory

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// this is the factory itself
func getGun(gunType string) (iGun, error) {
	if gunType == "musket" {
		return newMusket(), nil
	}

	if gunType == "ak47" {
		return newAK47(), nil
	}

	return nil, fmt.Errorf("Invalid gunType provided: %s", gunType)
}
