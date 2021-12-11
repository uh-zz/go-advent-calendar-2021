package afterme

import (
	"fmt"
)

func ExampleNewMapStatus() {
	player := NewPoint(0, 0)
	burden := NewPoint(1, 1)
	mark := NewPoint(1, 3)

	mapStatus := NewMapStatus(player, burden, mark)

	fmt.Printf("%+v", mapStatus)
	// Output: &{player:{X:0 Y:0} burden:{X:1 Y:1} mark:{X:1 Y:3}}
}

func ExampleNewMapStatusWithOption() {
	player := NewPoint(0, 0)
	burden := NewPoint(1, 1)
	mark := NewPoint(1, 3)

	beforeM := NewMapStatus(player, burden, mark)

	afterM := NewMapStatusWithOption(
		beforeM,
		Right(),
		Down(),
		Down())

	fmt.Printf("%+v", afterM)
	// Output: &{player:{X:1 Y:2} burden:{X:1 Y:3} mark:{X:1 Y:3}}
}
