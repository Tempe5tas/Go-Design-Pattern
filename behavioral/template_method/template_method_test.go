package template_method

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	fireRifle := NewFireRifle()
	fireRifle.FireGun()

	fmt.Println("--------------------")

	fireMG := NewFireMachineGun()
	fireMG.FireGun()
}
