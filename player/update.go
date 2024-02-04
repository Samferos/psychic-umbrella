package player

import "github.com/hajimehoshi/ebiten/v2"

func (p *Player) Update() {
	for action, key := range p.controls {
		if ebiten.IsKeyPressed(key) {
			p.action(action)
		}
	}
}
