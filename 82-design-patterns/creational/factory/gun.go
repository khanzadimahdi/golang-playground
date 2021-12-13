package factory

import "log"

type iGun interface {
	Name() string
	SetName(string)
	Power() int
	SetPower(int)
	Shoot()
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) Name() string {
	return g.name
}

func (g *Gun) SetName(n string) {
	g.name = n
}

func (g *Gun) Power() int {
	return g.power
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) Shoot() {
	log.Printf("shoot with power %v", g.power)
}

func getGun() iGun {
	gun := new(Gun)

	gun.SetName("ak47")
	gun.SetPower(100)

	return gun
}
