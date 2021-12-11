package afterme

// MapStatus マップ状態
type MapStatus struct {
	player point // 人の座標
	burden point // 荷物の座標
	mark   point // マーク座標
}

// point 座標
type point struct {
	X, Y int
}

// Newpoint pointコンストラクタ
func NewPoint(x, y int) *point {
	return &point{X: x, Y: y}
}

// NewMapStatus MapStatusコンストラクタ
func NewMapStatus(p, b, m *point) *MapStatus {
	return &MapStatus{
		player: *p,
		burden: *b,
		mark:   *m,
	}
}

// MoveFunc func(*MapStatus)型
type MoveFunc func(*MapStatus)

// Union 空レシーバ
func (m MoveFunc) Union() {}

// Command 抽象コマンド
type Command interface {
	Union()
}

// move 人の移動
func (m *MapStatus) move(x, y int) {

	m.player.X += x
	m.player.Y += y

	// 人とかぶったら荷物をずらす
	if m.player.X == m.burden.X && m.player.Y == m.burden.Y {
		m.burden.X += x
		m.burden.Y += y
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
