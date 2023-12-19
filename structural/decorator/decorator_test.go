package decorator

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	var weapon Weapon

	weapon = new(AR)
	weapon.Show()
	fmt.Println("-----------------")

	weapon = WithRedDot(weapon)
	weapon.Show()
	fmt.Println("-----------------")

	weapon = WithGrip(weapon)
	weapon.Show()
	fmt.Println("-----------------")
}
