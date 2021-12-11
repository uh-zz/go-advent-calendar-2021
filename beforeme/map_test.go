package beforeme

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var cmds []Command

	m := &MapStatus{
		pX: 1,
		pY: 0,
		bX: 3,
		bY: 1,
		mX: 5,
		mY: 1,
	}

	fmt.Println("before:", m) // before: &{1 0 3 1 5 1}

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
