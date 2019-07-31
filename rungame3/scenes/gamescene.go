package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tuckersGo/rungame/rungame3/actor"
	"github.com/tuckersGo/rungame/rungame3/bgscroller"
	"github.com/tuckersGo/rungame/rungame3/global"
)

// GameScene scene of game
type GameScene struct {
	bgscroller *bgscroller.Scroller
	runner     *actor.Runner
}

// Startup initialize GameScene
func (g *GameScene) Startup() {
	frameCount = 0

	g.runner = actor.NewRunner(0, global.ScreenHeight/2)

	backImg, _, err := ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.bgscroller = bgscroller.New(backImg, global.BGScrollSpeed)
	g.runner.SetState(actor.Running)
}

// Update GameScene
func (g *GameScene) Update(screen *ebiten.Image) error {
	g.bgscroller.Update(screen)
	// Running Animation
	g.runner.Update(screen)
	return nil
}
