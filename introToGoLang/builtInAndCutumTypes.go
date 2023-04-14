package main

import "fmt"

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "foo"
	intVar32   int32   = 1
	intVar64   int64   = 1
	intVar     int     = -1
	uintVar    uint    = 1
	uint32Var  uint32  = 1
	uint64Var  uint64  = 1
	uint8Var   uint8   = 0x1
	byteVar    byte    = 0x2
	runVar     rune    = 'a'
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

// function reciver which accpts the struct as ip
func (player Player) getHealth() int {
	return player.health
}

// building custum types from built-in type
type Weapon string

func getWeapon(w Weapon) string {
	//return w the compiler will throw error as the 'w' is of type Weapon not string
	return string(w)
}
func main() {
	player := Player{
		name:        "captain jack",
		health:      10.0,
		attackPower: 100.1,
	}

	fmt.Printf("the palyer value is %+v \n", player)
	fmt.Printf("%s healthe is %d\n", player.name, player.getHealth()) // substituting the "d" and "s" with "v" will get the same op

	users := map[string]int{ // make(map[string]int)
		"tom": 55,
	}

	users["bob"] = 100

	fmt.Printf("our map entry is %+v\n", users)

	fmt.Println("iterating over the map object")
	for key, val := range users {
		fmt.Printf("the key %s and value is %d \n", key, val)
	}

	fmt.Println("iterating over the map's key")
	for key := range users {
		fmt.Printf("the key %s   \n", key)
	}
	fmt.Println("iterating over the map values")
	for _, val := range users {
		fmt.Printf(" value is %d \n", val)
	}
	fmt.Println("for checking the key which is not present in the map")
	fmt.Printf("value for the temp1 key (which is not presnet) : %d\n", users["temp1"])

	val, ok := users["temp2"] // returns 0, false. it intfernce the value as 0
	// but it does not create new entry im the map

	if !ok {
		fmt.Printf("value of the key temp2 is %d\n", val)
		fmt.Printf("our map entry is %+v\n", users)

	}
	delete(users, "bob") // to delete the user["bob"]
	fmt.Printf("our map entry is %+v\n", users)

	numbers := []int{}
	fmt.Println(numbers)

	numbers = append(numbers, 1, 2)
	fmt.Println(numbers)
}
