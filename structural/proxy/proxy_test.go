package proxy

import "testing"

func TestPattern(t *testing.T) {
	var dealer BossDealer

	dealer.smuggler = new(RussianSmuggler)
	dealer.Buy(&Weapon{"rocket launcher", true})

	dealer.smuggler = new(AmericanSmuggler)
	dealer.Buy(&Weapon{"rifle", true})

	dealer.smuggler = new(AmericanSmuggler)
	dealer.Buy(&Weapon{"helicopter", false})
}
