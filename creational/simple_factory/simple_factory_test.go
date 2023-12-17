package simple_factory

import "testing"

func TestPattern(t *testing.T) {
	factory := new(Factory)

	p1 := factory.CreateProduct("CPU")
	p1.Show()

	p2 := factory.CreateProduct("GPU")
	p2.Show()

	p3 := factory.CreateProduct("SSD")
	p3.Show()
}
