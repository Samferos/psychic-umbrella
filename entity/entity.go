package entity

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetPosition() (x, y int)
	GetDrawable() (img *ebiten.Image)
	Update()
}
