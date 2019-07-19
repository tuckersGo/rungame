// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package font

import (
	"image/color"
	"log"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

const (
	arcadeFontBaseSize = 8
)

var (
	arcadeFonts map[int]font.Face
)

// getArcadeFonts returns Arcade font
func getArcadeFonts(scale int) font.Face {
	if arcadeFonts == nil {
		tt, err := truetype.Parse(fonts.ArcadeN_ttf)
		if err != nil {
			log.Fatal(err)
		}

		arcadeFonts = map[int]font.Face{}
		for i := 1; i <= 4; i++ {
			const dpi = 72
			arcadeFonts[i] = truetype.NewFace(tt, &truetype.Options{
				Size:    float64(arcadeFontBaseSize * i),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
		}
	}
	return arcadeFonts[scale]
}

// TextWidth returns width of text
func TextWidth(str string, scale int) int {
	maxW := 0
	for _, line := range strings.Split(str, "\n") {
		b, _ := font.BoundString(getArcadeFonts(scale), line)
		w := (b.Max.X - b.Min.X).Ceil()
		if maxW < w {
			maxW = w
		}
	}
	return maxW
}

var (
	shadowColor = color.NRGBA{0, 0, 0, 0x80}
)

// DrawTextWithShadow draws Text with Shadow
func DrawTextWithShadow(rt *ebiten.Image, str string, x, y, scale int, clr color.Color) {
	offsetY := arcadeFontBaseSize * scale
	for _, line := range strings.Split(str, "\n") {
		y += offsetY
		text.Draw(rt, line, getArcadeFonts(scale), x+1, y+1, shadowColor)
		text.Draw(rt, line, getArcadeFonts(scale), x, y, clr)
	}
}

// DrawTextWithShadowCenter draws Text with Shadow Center
func DrawTextWithShadowCenter(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str, scale)
	x += (width - w) / 2
	DrawTextWithShadow(rt, str, x, y, scale, clr)
}

// DrawTextWithShadowRight draws Text with Shadow right
func DrawTextWithShadowRight(rt *ebiten.Image, str string, x, y, scale int, clr color.Color, width int) {
	w := TextWidth(str, scale)
	x += width - w
	DrawTextWithShadow(rt, str, x, y, scale, clr)
}
