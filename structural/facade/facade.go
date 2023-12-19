package facade

import "fmt"

// Magazine

type Mag struct{}

func (m *Mag) Load() {
	fmt.Println("Loading magazine")
}

func (m *Mag) Unload() {
	fmt.Println("Unloading magazine")
}

// Safety

type Safety struct{}

func (s *Safety) On() {
	fmt.Println("Safety on")
}

func (s *Safety) Off() {
	fmt.Println("Safety off")
}

// Weapon

type Weapon struct {
	safety Safety
	mag    Mag
}

func (w *Weapon) On() {
	fmt.Println("Getting ready to fire")
	w.mag.Load()
	w.safety.Off()
}

func (w *Weapon) Off() {
	fmt.Println("Securing weapon")
	w.mag.Unload()
	w.safety.On()
}
