package afterme

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	player := NewPoint(0, 0)
	burden := NewPoint(1, 1)
	mark := NewPoint(1, 3)

	beforeM := NewMapStatus(player, burden, mark)

	fmt.Printf("before:%+v\n", beforeM) // before:&{player:{X:0 Y:0} burden:{X:1 Y:1} mark:{X:1 Y:3}}

	afterM := NewMapStatusWithOption(
		beforeM,
		Right(),
		Down(),
		Down())

	fmt.Printf("after:%+v\n", afterM) // after:&{player:{X:1 Y:2} burden:{X:1 Y:3} mark:{X:1 Y:3}}
}
