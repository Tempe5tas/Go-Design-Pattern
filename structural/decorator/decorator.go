package decorator

import "fmt"

// Abstract

type Weapon interface {
	Show()
}

type Accessory struct { // Should be interfaced, but can't use due to member variables.
	weapon Weapon
}

func (a *Accessory) Show() {}

// Concrete components

type AR struct {
	Weapon
	Name string
}

func (r *AR) Show() {
	fmt.Println("Showing HK416 assault rifle.")
}

type Pistol struct {
	Weapon
	Name string
}

func (r *Pistol) Show() {
	fmt.Println("Showing G17 pistol.")
}

// Concrete decorators

type RedDot struct {
	Accessory
}

func (a *RedDot) Show() {
	a.weapon.Show()
	fmt.Println("With red dot scope.")
}

func WithRedDot(w Weapon) Weapon {
	return &RedDot{Accessory{w}}
}

type Grip struct {
	Accessory
}

func (a *Grip) Show() {
	a.weapon.Show()
	fmt.Println("With foregrip.")
}

func WithGrip(w Weapon) Weapon {
	return &Grip{Accessory{w}}
}
