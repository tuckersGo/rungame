package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/tuckersGo/rungame/rungame1/global"
	"github.com/tuckersGo/rungame/rungame1/scenemanager"
	"github.com/tuckersGo/rungame/rungame1/scenes"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 2.0, "Run game")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
