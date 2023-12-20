package observer

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	var o Lieutenant
	var a1, a2, a3 Assault
	var m1, m2 Marksman
	var s1, s2 Support

	o.AddSoldier(&a1, &a2, &m1, &a3, &m2, &s1, &s2)
	o.Advance()

	fmt.Println("-------------------")

	o.Notify()
}
