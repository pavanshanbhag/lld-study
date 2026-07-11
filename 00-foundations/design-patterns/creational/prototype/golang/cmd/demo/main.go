package main

import "prototype"

func main() {
	registry := prototype.NewEnemyRegistry()

	registry.Register("flying", prototype.NewEnemy("FlyingEnemy", 100, 12.0, false, "Laser"))
	registry.Register("armored", prototype.NewEnemy("ArmoredEnemy", 300, 6.0, true, "Cannon"))

	e1 := registry.Get("flying")
	e2 := registry.Get("flying")
	e2.SetHealth(80)

	e3 := registry.Get("armored")

	e1.PrintStats()
	e2.PrintStats()
	e3.PrintStats()
}
