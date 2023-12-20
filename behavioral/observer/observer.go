package observer

import "fmt"

// Imagine that a group of soldiers marching on the front line

// Abstract subject

type Officer interface {
	AddSoldier(soldier Soldier)
	RemoveSoldier(soldier Soldier)
	Notify()
	Advance()
}

// Abstract observer

type Soldier interface {
	March()
	Hold()
}

// Implementation

type Assault struct{}

func (a *Assault) March() {
	fmt.Println("Assault marching")
}

func (a *Assault) Hold() {
	fmt.Println("Assault guarding rear")
}

type Support struct{}

func (s *Support) March() {
	fmt.Println("Assault marching")
}

func (s *Support) Hold() {
	fmt.Println("Assault guarding front")
}

type Marksman struct{}

func (m *Marksman) March() {
	fmt.Println("Marksman marching")
}

func (m *Marksman) Hold() {
	fmt.Println("Marksman searching for enemies")
}

// Notifier

type Lieutenant struct {
	soldiers []Soldier
}

func (l *Lieutenant) AddSoldier(soldiers ...Soldier) {
	for _, s := range soldiers {
		l.soldiers = append(l.soldiers, s)
	}
}

func (l *Lieutenant) RemoveSoldier(soldier Soldier) {
	for index, s := range l.soldiers {
		if soldier == s {
			l.soldiers = append(l.soldiers[:index], l.soldiers[index+1:]...)
			break
		}
	}
}

func (l *Lieutenant) Notify() {
	fmt.Println("Lieutenant sensed something dangerous")
	fmt.Println("Lieutenant orders his soldiers to hold")
	for _, s := range l.soldiers {
		s.Hold()
	}
}

func (l *Lieutenant) Advance() {
	fmt.Println("Lieutenant grouped his soldiers")
	fmt.Println("Lieutenant orders his soldiers to march")
	for _, s := range l.soldiers {
		s.March()
	}
}
