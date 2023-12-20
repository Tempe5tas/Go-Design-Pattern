package strategy

import "fmt"

// Imagine that you're playing HOI4, and selecting land doctrines
// Abstract strategy

type Doctrine interface {
	Tip()
}

// Concrete strategies

type MobileWarfare struct{}

func (m *MobileWarfare) Tip() {
	fmt.Println("Selected mobile warfare.")
	fmt.Println("Should build more trucks and tanks for war.")
}

type SuperiorFirepower struct{}

func (s *SuperiorFirepower) Tip() {
	fmt.Println("Selected superior firepower.")
	fmt.Println("Should build more cannons for war.")
}

type GrandBattleplan struct{}

func (g *GrandBattleplan) Tip() {
	fmt.Println("Selected grand battleplan.")
	fmt.Println("Be careful about your strategic deployment.")
}

type MassAssault struct{}

func (m *MassAssault) Tip() {
	fmt.Println("Selected mass assault.")
	fmt.Println("Be sure that you have enough manpower.")
}

// Context

type General struct {
	doctrine Doctrine
}

// Set a strategy

func (g *General) SetDoctrine(d Doctrine) {
	g.doctrine = d
}

func (g *General) Tip() {
	g.doctrine.Tip()
}
