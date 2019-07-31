package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tuckersGo/rungame/rungame3/actor"
	"github.com/tuckersGo/rungame/rungame3/font"
	"github.com/tuckersGo/rungame/rungame3/global"
	"github.com/tuckersGo/rungame/rungame3/scenemanager"
)

// StartScene first scene
type StartScene struct {
	backImg *ebiten.Image
	runner  *actor.Runner
}

// Startup StartScene
func (s *StartScene) Startup() {
	s.runner = actor.NewRunner(0, global.ScreenHeight/2)

	var err error
	s.backImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	s.runner.SetState(actor.Idle)
}

var frameCount = 0

// Update StartScene
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	screen.DrawImage(s.backImg, nil)

	// Idle Animation
	s.runner.Update(screen)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-width/2, global.ScreenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}

	return nil
}
