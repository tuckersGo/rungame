package animation

import "github.com/hajimehoshi/ebiten"

type animInfo struct {
	sprites []*ebiten.Image
	speed   int
}

// Handler animation handler
type Handler struct {
	animMap      map[string]animInfo
	currentAnim  animInfo
	lastIdx      int
	remainframes int
}

// New make a new animation handler
func New() *Handler {
	h := &Handler{}
	h.animMap = make(map[string]animInfo)
	return h
}

// Add add new animation
func (h *Handler) Add(name string, sprites []*ebiten.Image, speed int) {
	h.animMap[name] = animInfo{sprites, speed}
}

// Play play the animation
func (h *Handler) Play(name string) {
	h.currentAnim = h.animMap[name]
	h.lastIdx = 0
	h.remainframes = h.currentAnim.speed
}

// Update draw animation frame
func (h *Handler) Update(screen *ebiten.Image, x, y float64) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(h.currentAnim.sprites[h.lastIdx], &op)

	h.remainframes--
	if h.remainframes == 0 {
		h.lastIdx++
		if len(h.currentAnim.sprites) == h.lastIdx {
			h.lastIdx = 0
		}
		h.remainframes = h.currentAnim.speed
	}
}
