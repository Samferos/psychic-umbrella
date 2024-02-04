package world

import "github.com/hajimehoshi/ebiten/v2"

func (world World) Draw(screen *ebiten.Image) {
	for _, entity := range world.entityRegistry {
		op := ebiten.DrawImageOptions{}
		x, y := entity.GetPosition()
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(entity.GetDrawable(), &op)
	}
}
