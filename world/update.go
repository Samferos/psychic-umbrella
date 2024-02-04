package world

func (world *World) Update() {
	for _, entity := range world.entityRegistry {
		entity.Update()
	}
}
