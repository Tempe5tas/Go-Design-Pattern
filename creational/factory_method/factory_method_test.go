package factory_method

import "testing"

func TestPattern(t *testing.T) {
	// Define factory and product variables
	var factory AbstractFactory
	var product AbstractProduct

	// Create product: rifle
	factory = new(RifleFactory)
	product = factory.CreateProduct()
	product.Show()

	// Create product: pistol
	factory = new(PistolFactory)
	product = factory.CreateProduct()
	product.Show()

	// Create product: Scope
	factory = new(ScopeFactory)
	product = factory.CreateProduct()
	product.Show()
}
