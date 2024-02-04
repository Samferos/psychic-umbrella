package player

import (
	"github.com/Samferos/psychic-umbrella.git/entity/status"
	"github.com/Samferos/psychic-umbrella.git/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	idle = iota
	rolling
)

type Player struct {
	controls map[string]ebiten.Key
	stats    status.Status
	position world.Position
	sprite   *ebiten.Image
	state    int
}
