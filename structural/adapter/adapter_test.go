package adapter

import "testing"

func TestPattern(t *testing.T) {
	ak := NewAK74M(NewAdapter(new(Picatinny)))
	ak.Use()
}
