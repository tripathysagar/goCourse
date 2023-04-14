package main

import "fmt"

/*
create a weapon with following damage
Axe -> 100
Swoard -> 90
WoodenStick -> 1
Knife -> 40
*/

type WeaponType int

const (
	Axe         WeaponType = iota // assings 0 then ++to next variable
	Swoard                        // 1
	WoodenStick                   // 2
	Knife                         // 3
)

func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "Axe"
	case Swoard:
		return "Swoard"
	case WoodenStick:
		return "WoodenStick"
	case Knife:
		return "knife"
	default:
		panic("wapon not found")
	}
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Swoard:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("wapon does not exit")
	}
}
func main() {
	var Gun WeaponType = 200
	fmt.Printf("damage %s of weapon  %d\n", Axe, getDamage(Axe))
	fmt.Printf("damage %s of weapon  %d\n", Swoard, getDamage(Swoard))
	fmt.Printf("damage %s of weapon  %d\n", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("damage %s of weapon  %d\n", Knife, getDamage(Knife))
	fmt.Printf("damage %s of weapon  %d\n", Gun, getDamage(Gun))
}
