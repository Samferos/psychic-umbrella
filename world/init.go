package world

import "github.com/Samferos/psychic-umbrella.git/entity"

func (world *World) Init(entities []entity.Entity) {
	world.entityRegistry = entities
}
