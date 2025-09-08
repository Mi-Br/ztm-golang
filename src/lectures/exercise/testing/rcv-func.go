//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package player

import (
	"fmt"
)

type Player struct {
	Name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

func (p *Player) Consume(h, e int) {
	p.Health += h
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	p.Energy += e
	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}
}

func (p Player) String() string {
	return fmt.Sprintf("%s stats changed to %d %d \n", p.Name, p.Health, p.Energy)
}
func main() {
	var p1 Player = Player{Name: "Johny", Health: 100, MaxHealth: 100, Energy: 300, MaxEnergy: 300}
	p1.Consume(-10, 10)
	fmt.Print(p1)
	p1.Consume(-30, -100)
	fmt.Print(p1)

}
