package proxy

import "fmt"

// Abstract, not necessary

type Smuggler interface {
	Buy(w *Weapon)
}

// Concrete

type Weapon struct {
	Kind    string
	Genuine bool
}

type RussianSmuggler struct {
}

func (s *RussianSmuggler) Buy(w *Weapon) {
	fmt.Println("Bought", w.Kind, "from Russia.")
}

type AmericanSmuggler struct {
}

func (s *AmericanSmuggler) Buy(w *Weapon) {
	fmt.Println("Bought", w.Kind, "from America.")
}

// Proxy

type BossDealer struct {
	smuggler Smuggler
}

func (d *BossDealer) Buy(w *Weapon) {
	fmt.Println("Order received:", w.Kind)
	if !w.Genuine {
		fmt.Println("Warning: not genuine weapon:", w.Kind)
	} else {
		d.smuggler.Buy(w)
	}
	fmt.Println("Order done:", w.Kind)
	fmt.Println("-------------------------")
}
