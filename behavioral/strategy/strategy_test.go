package strategy

import "testing"

func TestPattern(t *testing.T) {
	you := General{}

	you.SetDoctrine(new(MassAssault))
	you.Tip()

	you.SetDoctrine(new(MobileWarfare))
	you.Tip()
}
