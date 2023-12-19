package adapter

import "fmt"

// Adaptive goal

type Dovetail interface {
	UseDovetail()
}

// Business

type AK74M struct {
	rail Dovetail
}

func NewAK74M(r Dovetail) *AK74M {
	return &AK74M{r}
}

func (w *AK74M) Use() {
	fmt.Println("Using AK-74M with accessories")
	w.rail.UseDovetail()
}

// Role to be adapted

type Picatinny struct{}

func (p *Picatinny) UsePic() {
	fmt.Println("Using picatinny rail")
}

type Adapter struct {
	pic *Picatinny
}

func (a *Adapter) UseDovetail() {
	fmt.Println("Using dovetail")
	a.pic.UsePic()
}

func NewAdapter(pic *Picatinny) *Adapter {
	return &Adapter{pic}
}
