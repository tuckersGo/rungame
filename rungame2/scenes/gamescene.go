package scenes

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tuckersGo/rungame/rungame2/global"
)

// GameScene scene of game
type GameScene struct {
	runnerImg *ebiten.Image
	backImg   *ebiten.Image
}

// Startup initialize GameScene
func (g *GameScene) Startup() {
	frameCount = 0

	var err error
	g.runnerImg, _, err = ebitenutil.NewImageFromFile("./images/runner.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	g.backImg, _, err = ebitenutil.NewImageFromFile("./images/background.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

// Update GameScene
func (g *GameScene) Update(screen *ebiten.Image) error {
	frameCount++

	bgWidth, _ := g.backImg.Size()

	op := &ebiten.DrawImageOptions{}
	backX := (frameCount / 2) % bgWidth
	op.GeoM.Translate(float64(-backX), 0)
	screen.DrawImage(g.backImg, op)

	op.GeoM.Translate(float64(bgWidth), 0)
	screen.DrawImage(g.backImg, op)

	// Running Animation
	frameIdx := (frameCount / global.RunningAnimSpeed) % global.RunningFrames
	sx := global.RunningX + global.FrameWidth*frameIdx
	sy := global.RunningY

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, global.ScreenHeight/2)
	subImg := g.runnerImg.SubImage(image.Rect(sx, sy, sx+global.FrameWidth, sy+global.FrameHeight)).(*ebiten.Image)
	screen.DrawImage(subImg, op)
	return nil
}
