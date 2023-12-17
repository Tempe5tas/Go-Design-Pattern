package abstract_factory

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	var factory AbstractFactory
	var cpu AbstractCPU
	var gpu AbstractGPU

	fmt.Println("Assembling new PC:")
	factory = new(Intel)
	cpu = factory.CreateCPU()
	cpu.Use()
	factory = new(Nvidia)
	gpu = factory.CreateGPU()
	gpu.Use()

	fmt.Println("Assembling another PC:")
	factory = new(AMD)
	cpu = factory.CreateCPU()
	cpu.Use()
	gpu = factory.CreateGPU()
	gpu.Use()
}
