package world

// Position util
func (pos *Position) Translate(x, y int) {
	pos.X, pos.Y = pos.X+x, pos.Y+y
}
