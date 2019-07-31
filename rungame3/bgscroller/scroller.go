package bgscroller

import (
	"github.com/hajimehoshi/ebiten"
)

// Scroller horizontal background scroller
type Scroller struct {
	bgimg  *ebiten.Image
	speed  int
	frames int
}

func New(bgimg *ebiten.Image, speed int) *Scroller {
	return &Scroller{bgimg, speed, 0}
}

func (s *Scroller) Update(screen *ebiten.Image) {
	s.frames++

	bgWidth, _ := s.bgimg.Size()

	op := &ebiten.DrawImageOptions{}
	backX := (s.frames / s.speed) % bgWidth
	op.GeoM.Translate(float64(-backX), 0)
	screen.DrawImage(s.bgimg, op)

	op.GeoM.Translate(float64(bgWidth), 0)
	screen.DrawImage(s.bgimg, op)
}
