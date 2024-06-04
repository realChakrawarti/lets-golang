package game

import "fmt"

type Player struct {
	health int
	energy int
	name   string
}

func NewPlayer(name string, health, energy int) Player {
	fmt.Printf("\n%v's Stats | Health: %v | Energy: %v", name, health, energy)
	return Player{health, energy, name}
}

func (p *Player) UpdateHealth(health int) {
	p.health = health
	fmt.Printf("\n%v's Stats | Health: %v | Energy: %v", p.name, p.health, p.energy)
}

func (p *Player) UpdateEnergy(energy int) {
	p.energy = energy
	fmt.Printf("\n%v's Stats | Health: %v | Energy: %v", p.name, p.health, p.energy)
}
