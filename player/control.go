package player

func (p *Player) action(verb string) {
	if p.state == rolling {
		p.roll()
	}
	// Movement controls
	if p.state == idle {
		var x, y int
		switch verb {
		case "up":
			x, y = 0, -1
		case "down":
			x, y = 0, 1
		case "right":
			x, y = 1, 0
		case "left":
			x, y = -1, 0
		}
		p.position.Translate(x, y)
	}
}

func (p *Player) roll() {

}
