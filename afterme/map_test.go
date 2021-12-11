package afterme

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var cmds []Command

	player := NewPoint(1, 0)
	burden := NewPoint(3, 1)
	mark := NewPoint(5, 1)

	m := NewMapStatus(player, burden, mark)

	fmt.Printf("before:%+v\n", m) // before: &{1 0 3 1 5 1}

	cmds = append(cmds, MoveFunc(Right))
	cmds = append(cmds, MoveFunc(Down))
	cmds = append(cmds, MoveFunc(Right))
	cmds = append(cmds, MoveFunc(Right))

	cmds = cmds[:len(cmds)-1] // １つ戻る

	for _, cmd := range cmds {
		switch t := cmd.(type) {
		case MoveFunc:
			t(m)
		}

	}

	fmt.Println("after:", m) // after: &{3 1 4 1 5 1}
}
