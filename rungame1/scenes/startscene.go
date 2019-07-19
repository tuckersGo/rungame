package scenes

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tuckersGo/rungame/rungame1/global"
)

// StartScene first scene
type StartScene struct {
	runnerImg *ebiten.Image
}

// Startup StartScene
func (s *StartScene) Startup() {
	var err error
	s.runnerImg, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

var frameCount = 0

// Update StartScene
func (s *StartScene) Update(screen *ebiten.Image) error {
	frameCount++

	frameIdx := (frameCount / 5) % global.RunningFrames
	sx := global.RunningX + global.FrameWidth*frameIdx
	sy := global.RunningY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(global.ScreenWidth/2, global.ScreenHeight/2)
	subImg := s.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)
	return nil
}
