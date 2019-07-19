package scenes

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tuckersGo/rungame/rungame2/font"
	"github.com/tuckersGo/rungame/rungame2/global"
	"github.com/tuckersGo/rungame/rungame2/scenemanager"
)

// StartScene first scene
type StartScene struct {
	runnerImg *ebiten.Image
	backImg   *ebiten.Image
}

// Startup StartScene
func (s *StartScene) Startup() {
	var err error
	s.runnerImg, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	s.backImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

var frameCount = 0

// Update StartScene
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	screen.DrawImage(s.backImg, nil)

	// Idle Animation
	frameIdx := (frameCount / global.IdleAnimSpeed) % global.IdleFrames
	sx := global.IdleX + global.FrameWidth*frameIdx
	sy := global.IdleY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, global.ScreenHeight/2)
	subImg := s.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)

	width := font.TextWidth(global.StartSceneText, 2)
	font.DrawTextWithShadow(screen, global.StartSceneText,
		global.ScreenWidth/2-width/2, global.ScreenHeight/2, 2, color.White)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&GameScene{})
	}

	return nil
}
