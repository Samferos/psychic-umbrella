package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Player) GetPosition() (int, int) {
	return p.position.X, p.position.Y
}

func (p *Player) GetDrawable() *ebiten.Image {
	return p.sprite
}

func NewPlayer(controls map[string]ebiten.Key) *Player {
	sprite := ebiten.NewImage(5, 5)
	sprite.Fill(color.White)
	return &Player{controls: controls, sprite: sprite}
}
