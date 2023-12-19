package facade

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	ak := new(Weapon)

	fmt.Println("Detected conflict, getting ready")
	ak.On()

	fmt.Println("Conflict resolved, securing weapon")
	ak.Off()
}
