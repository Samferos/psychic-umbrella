package main

import (
	"log"

	"github.com/Samferos/psychic-umbrella.git/entity"
	"github.com/Samferos/psychic-umbrella.git/player"
	"github.com/Samferos/psychic-umbrella.git/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var controlScheme map[string]ebiten.Key

type Game struct {
	gameWorld world.World
}

func (g *Game) Update() error {
	g.gameWorld.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameWorld.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pyschic Umbrella")
	game := &Game{}
	controlScheme = map[string]ebiten.Key{
		"up":    ebiten.KeyArrowUp,
		"down":  ebiten.KeyArrowDown,
		"right": ebiten.KeyArrowRight,
		"left":  ebiten.KeyArrowLeft,
		"dodge": ebiten.KeyShift,
		"use":   ebiten.KeySpace,
	}
	game.gameWorld.Init([]entity.Entity{player.NewPlayer(controlScheme)})
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
