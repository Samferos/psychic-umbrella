package world

import "github.com/Samferos/psychic-umbrella.git/entity"

type World struct {
	entityRegistry []entity.Entity
	// terrain TODO
}

type Position struct {
	X, Y int
}
