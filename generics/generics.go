package generics

// MapStatus マップ状態
type MapStatus[T Addable] struct {
	player point[T] // 人の座標
	burden point[T] // 荷物の座標
	mark   point[T] // マーク座標
}

type Addable interface {
	~int
}

// point 座標
type point[T Addable] struct {
	X, Y T
}

func (p *point[T]) Add(x, y T) {
	p.X += x
	p.Y += y
}

// Newpoint pointコンストラクタ
func NewPoint[T Addable](x, y T) *point[T] {
	// return &point{X: x, Y: y}
	return new(point[T])
}

// NewMapStatus MapStatusコンストラクタ
func NewMapStatus[T Addable](p, b, m *point[T]) *MapStatus[T] {
	return new(MapStatus[T])
}

// Option func(*MapStatus)型
type Option[T Addable] func(*MapStatus[T])

// Right 右に１つ進む
func Right[T Addable]() Option[T] {
	return func(m *MapStatus[T]) {
		m.move(1, 0)
	}
}

// Down 下に１つ進む
func Down[T Addable]() Option[T] {
	return func(m *MapStatus[T]) {
		m.move(0, 1)
	}
}

// move 人の移動
func (m *MapStatus[T]) move(x, y T) {

	m.player.X += x
	m.player.Y += y

	// 人とかぶったら荷物をずらす
	if m.player.X == m.burden.X && m.player.Y == m.burden.Y {
		m.burden.X += x
		m.burden.Y += y
	}
}

// NewMapStatusWithOption アクション後のマップ状態
func NewMapStatusWithOption[T Addable](m *MapStatus[T], options ...Option[T]) *MapStatus[T] {
	for _, option := range options {
		option(m)
	}
	return m
}
