package simple_factory

import (
	"fmt"
)

// Abstract product

type AbstractProduct interface {
	Show()
}

// Concrete products

type CPU struct {
	AbstractProduct
}

func (p *CPU) Show() {
	fmt.Println("Showing CPU")
}

type GPU struct {
	AbstractProduct
}

func (p *GPU) Show() {
	fmt.Println("Showing graphic card")
}

type Memory struct {
	AbstractProduct
}

func (p *Memory) Show() {
	fmt.Println("Showing memory module")
}

type SSD struct {
	AbstractProduct
}

func (p *SSD) Show() {
	fmt.Println("Showing solid state drive")
}

// Factory

type Factory struct{}

func (f *Factory) CreateProduct(kind string) (p AbstractProduct) {
	switch kind {
	case "CPU":
		p = new(CPU)
	case "GPU":
		p = new(GPU)
	case "Memory":
		p = new(Memory)
	case "SSD":
		p = new(SSD)
	default:
		panic("invalid product type")
	}
	return
}
