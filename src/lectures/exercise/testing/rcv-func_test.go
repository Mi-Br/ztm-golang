// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package player

import "testing"

func TestMaxEnergy(t *testing.T) {

	p := Player{Name: "Mike", Health: 100, MaxHealth: 200, Energy: 100, MaxEnergy: 200}

	p.Consume(500, 500)
	if p.Health > p.MaxHealth || p.Energy > p.MaxEnergy {
		t.Errorf("Player Max health or Energy exceeds")
		t.Errorf("Health: %d, MaxHealth: %d, Energy: %d, MaxEnergy: %d", p.Health, p.MaxHealth, p.Energy, p.MaxEnergy)
	}
}
