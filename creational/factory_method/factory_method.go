package factory_method

import (
	"fmt"
)

// Abstract product and factory

type AbstractProduct interface {
	Show()
}

type AbstractFactory interface {
	CreateProduct() AbstractProduct
}

// Concrete products

type Rifle struct {
	AbstractProduct
}

func (r *Rifle) Show() {
	fmt.Println("Showing AK-12")
}

type Pistol struct {
	AbstractProduct
}

func (p *Pistol) Show() {
	fmt.Println("Showing MP-433")
}

type Scope struct {
	AbstractProduct
}

func (s *Scope) Show() {
	fmt.Println("Showing PSO-1")
}

// Concrete factories

type RifleFactory struct {
	AbstractFactory
}

func (f *RifleFactory) CreateProduct() (p AbstractProduct) {
	p = new(Rifle)
	return
}

type PistolFactory struct {
	AbstractFactory
}

func (f *PistolFactory) CreateProduct() (p AbstractProduct) {
	p = new(Pistol)
	return
}

type ScopeFactory struct {
	AbstractFactory
}

func (f *ScopeFactory) CreateProduct() (p AbstractProduct) {
	p = new(Scope)
	return
}
