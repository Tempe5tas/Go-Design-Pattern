package template_method

import "fmt"

// Abstract

type Gun interface {
	Load()      // Load weapon
	SafetyOff() // Disable safety
	Aim()       // Aim at target
	Fire()      // Fire at target
}

// Template

type template struct {
	g Gun
}

// Fixed template for encapsulation

func (t *template) FireGun() {
	if t == nil {
		return
	}
	t.g.Load()
	t.g.SafetyOff()
	t.g.Aim()
	t.g.Fire()
}

// Specific templates: firing rifle

type FireRifle struct {
	template
}

func (f *FireRifle) Load() {
	fmt.Println("Loading magazine")
}

func (f *FireRifle) SafetyOff() {
	fmt.Println("Switching to semi-auto")
}

func (f *FireRifle) Aim() {
	fmt.Println("Aiming with a holographic sight")
}

func (f *FireRifle) Fire() {
	fmt.Println("Pulling trigger and fire")
}

func NewFireRifle() *FireRifle {
	fire := new(FireRifle)
	fire.g = fire
	return fire
}

// Specific templates: firing machine gun

type FireMachineGun struct {
	template
}

func (f *FireMachineGun) Load() {
	fmt.Println("Loading cartridge chain")
}

func (f *FireMachineGun) SafetyOff() {
	fmt.Println("Switching to full-auto")
}

func (f *FireMachineGun) Aim() {
	fmt.Println("Aiming with a optical sight")
}

func (f *FireMachineGun) Fire() {
	fmt.Println("Pulling trigger and fire")
}

func NewFireMachineGun() *FireMachineGun {
	fire := new(FireMachineGun)
	fire.g = fire
	return fire
}
