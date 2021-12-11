package beforeme

type MapStatus struct {
	pX, pY int // プレイヤー座標
	bX, bY int // 荷物座標
	mX, mY int // マーク座標
}

// MoveFunc func(*MapStatus)型
type MoveFunc func(*MapStatus)

// Union 空レシーバ
func (m MoveFunc) Union() {}

// Command 抽象コマンド
type Command interface {
	Union()
}

// move プレイヤーの移動
func (m *MapStatus) move(x, y int) {

	m.pX += x
	m.pY += y

	// プレイヤーとかぶったら荷物をずらす
	if m.pX == m.bX && m.pY == m.bY {
		m.bX += x
		m.bY += y
	}
}

// Right 右へ進む
func Right(m *MapStatus) {
	m.move(1, 0)
}

// Down 下へ進む
func Down(m *MapStatus) {
	m.move(0, 1)
}
