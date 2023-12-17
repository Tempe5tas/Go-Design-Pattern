package abstract_factory

import "fmt"

// Abstract entities

type AbstractFactory interface {
	CreateCPU() AbstractCPU
	CreateGPU() AbstractGPU
}

type AbstractCPU interface {
	Use()
}

type AbstractGPU interface {
	Use()
}

// Concrete entities

type Intel struct {
	AbstractFactory
}

func (f *Intel) CreateCPU() (p AbstractCPU) {
	p = new(IntelCPU)
	return
}

func (f *Intel) CreateGPU() (p AbstractGPU) {
	p = new(IntelGPU)
	return
}

type IntelCPU struct {
	AbstractCPU
}

func (p *IntelCPU) Use() {
	fmt.Println("Using Intel Xeon Platinum 8153")
}

type IntelGPU struct {
	AbstractGPU
}

func (p *IntelGPU) Use() {
	fmt.Println("Using Intel Arc A770")
}

type AMD struct {
	AbstractFactory
}

func (f *AMD) CreateCPU() (p AbstractCPU) {
	p = new(AMDCPU)
	return
}

func (f *AMD) CreateGPU() (p AbstractGPU) {
	p = new(AMDGPU)
	return
}

type AMDCPU struct {
	AbstractCPU
}

func (p *AMDCPU) Use() {
	fmt.Println("Using AMD EPYC 9654")
}

type AMDGPU struct {
	AbstractGPU
}

func (p *AMDGPU) Use() {
	fmt.Println("Using AMD Radeon RX 7900 XTX")
}

type Nvidia struct {
	AbstractFactory
}

func (f *Nvidia) CreateGPU() (p AbstractGPU) {
	p = new(NvidiaGPU)
	return
}

type NvidiaGPU struct {
	AbstractGPU
}

func (p *NvidiaGPU) Use() {
	fmt.Println("Using NVIDIA GeForce RTX 4080")
}
