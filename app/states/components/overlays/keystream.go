package overlays

import (
	"math"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/wieku/danser-go/app/settings"
	"github.com/wieku/danser-go/framework/graphics/batch"
	"github.com/wieku/danser-go/framework/graphics/font"
	"github.com/wieku/danser-go/framework/graphics/shape"
	"github.com/wieku/danser-go/framework/math/mutils"
	"github.com/wieku/danser-go/framework/math/vector"
)

const (
	streamBarAlpha    = 0.4  // Blondazz's barColor: 255,255,255,100
	streamFadeHeight  = 0.36 // part of the canvas where bars fade out (Blondazz: 255px of 700)
	streamSegmentSize = 16.0 // bars are drawn in segments so the fade can be a gradient
)

type streamBar struct {
	top    float64
	height float64
}

type streamKey struct {
	letter  string
	pressed bool
	count   int
	bars    []streamBar
}

// keyStream reproduces the classic osu! streaming key overlay (Blondazz/KeyOverlay):
// a row of bordered squares with the key letter, where every press spawns a bar that
// travels upwards at a constant speed. While the key is held the bar keeps stretching
// downwards, so its bottom edge stays pinned to the key and its top edge recedes.
type keyStream struct {
	keys   []*streamKey
	width  float64
	height float64
}

func newKeyStream(letters []string, width, height float64) *keyStream {
	stream := &keyStream{
		keys:   make([]*streamKey, 0, len(letters)),
		width:  width,
		height: height,
	}

	for _, letter := range letters {
		stream.keys = append(stream.keys, &streamKey{letter: letter})
	}

	return stream
}

func (stream *keyStream) boxMetrics() (size, outline, margin, boxWidth float64) {
	cfg := settings.Gameplay.KeyOverlay
	scale := cfg.Scale

	size = cfg.KeySize * scale
	outline = cfg.Outline * scale
	margin = cfg.Margin * scale
	boxWidth = size + outline*2

	return
}

func (stream *keyStream) boxTop() float64 {
	_, _, margin, boxWidth := stream.boxMetrics()

	return stream.height - margin - boxWidth + settings.Gameplay.KeyOverlay.YOffset
}

func (stream *keyStream) boxLeft(index int) float64 {
	_, _, margin, boxWidth := stream.boxMetrics()

	return stream.width - margin - boxWidth - float64(len(stream.keys)-1-index)*(boxWidth+margin) + settings.Gameplay.KeyOverlay.XOffset
}

func (stream *keyStream) press(index int) {
	if index < 0 || index >= len(stream.keys) {
		return
	}

	key := stream.keys[index]

	if key.pressed {
		return
	}

	key.pressed = true
	key.count++

	// every press spawns a new bar at the top edge of the key
	key.bars = append(key.bars, streamBar{top: stream.boxTop()})
}

func (stream *keyStream) release(index int) {
	if index < 0 || index >= len(stream.keys) {
		return
	}

	stream.keys[index].pressed = false
}

func (stream *keyStream) update(delta float64) {
	if delta <= 0 {
		return
	}

	dist := settings.Gameplay.KeyOverlay.BarSpeed * delta / 1000

	if dist <= 0 {
		return
	}

	for _, key := range stream.keys {
		if key.pressed {
			if len(key.bars) == 0 {
				key.bars = append(key.bars, streamBar{top: stream.boxTop()})
			}

			// only the newest bar stretches, its bottom edge stays pinned to the key
			key.bars[len(key.bars)-1].height += dist
		}

		// every bar travels upwards, held or not (this is what keeps older trails moving)
		for i := range key.bars {
			key.bars[i].top -= dist
		}

		// A configured bar length acts like a fixed-length lane: whatever would pass
		// the far end is clipped there, so even a long hold stops stretching at the
		// limit instead of running to the top of the screen.
		if barLength := settings.Gameplay.KeyOverlay.BarLength; barLength > 0 {
			minTop := stream.boxTop() - barLength

			for i := range key.bars {
				bar := &key.bars[i]

				if bar.top < minTop {
					bottom := bar.top + bar.height

					bar.top = minTop
					bar.height = bottom - minTop
				}
			}

			for len(key.bars) > 0 && key.bars[0].height <= 0 {
				key.bars = key.bars[1:]
			}
		} else {
			// unlimited: drop bars that left the canvas through the top edge
			cut := 0
			for cut < len(key.bars) && key.bars[cut].top+key.bars[cut].height < 0 {
				cut++
			}

			if cut > 0 {
				key.bars = key.bars[cut:]
			}
		}
	}
}

func (stream *keyStream) draw(batch *batch.QuadBatch, shapes *shape.Renderer, camera mgl32.Mat4, keyFont *font.Font, alpha float64) {
	cfg := settings.Gameplay.KeyOverlay

	size, outline, margin, boxWidth := stream.boxMetrics()

	bottom := stream.height - margin
	top := bottom - boxWidth

	fadeZone := stream.height * streamFadeHeight

	// alpha of a bar pixel at canvas y; when a bar length is configured the fade
	// ends exactly where the bar disappears, so nothing pops out of existence
	alphaAt := func(y float64) float64 {
		a := streamBarAlpha * alpha

		if !cfg.Fade {
			return a
		}

		if cfg.BarLength > 0 {
			return a * mutils.Clamp(1-(top-y)/cfg.BarLength, 0, 1)
		}

		return a * mutils.Clamp(y/fadeZone, 0, 1)
	}

	batch.Flush()

	shapes.SetCamera(camera)
	shapes.Begin()

	// bars first, so the key squares sit on top of their own trail
	for i, key := range stream.keys {
		left := stream.boxLeft(i)
		right := left + boxWidth

		for _, bar := range key.bars {
			segments := max(1, int(math.Ceil(bar.height/streamSegmentSize)))

			for s := 0; s < segments; s++ {
				y0 := bar.top + bar.height*float64(s)/float64(segments)
				y1 := bar.top + bar.height*float64(s+1)/float64(segments)

				barAlpha := alphaAt(y0)

				if barAlpha < 0.004 {
					continue
				}

				shapes.SetColor(1, 1, 1, barAlpha)
				shapes.DrawQuad(float32(left), float32(y0), float32(right), float32(y0), float32(right), float32(y1), float32(left), float32(y1))
			}
		}
	}

	// key squares: a border drawn as four strips so the inside stays transparent,
	// plus a translucent fill while the key is held
	for i, key := range stream.keys {
		left := stream.boxLeft(i)
		right := left + boxWidth

		shapes.SetColor(1, 1, 1, alpha)

		// top / bottom
		shapes.DrawQuad(float32(left), float32(top), float32(right), float32(top), float32(right), float32(top+outline), float32(left), float32(top+outline))
		shapes.DrawQuad(float32(left), float32(bottom-outline), float32(right), float32(bottom-outline), float32(right), float32(bottom), float32(left), float32(bottom))

		// left / right
		shapes.DrawQuad(float32(left), float32(top), float32(left+outline), float32(top), float32(left+outline), float32(bottom), float32(left), float32(bottom))
		shapes.DrawQuad(float32(right-outline), float32(top), float32(right), float32(top), float32(right), float32(bottom), float32(right-outline), float32(bottom))

		if key.pressed {
			shapes.SetColor(1, 1, 1, streamBarAlpha*alpha)
			shapes.DrawQuad(
				float32(left+outline), float32(top+outline),
				float32(right-outline), float32(top+outline),
				float32(right-outline), float32(bottom-outline),
				float32(left+outline), float32(bottom-outline),
			)
		}
	}

	shapes.End()

	batch.ResetTransform()

	for i, key := range stream.keys {
		left := stream.boxLeft(i)
		centerX := left + boxWidth/2

		batch.SetColor(1, 1, 1, alpha)
		keyFont.DrawOrigin(batch, centerX, top+boxWidth/2, vector.Centre, size*0.42, true, key.letter)

		if cfg.Counter {
			batch.SetColor(1, 1, 1, alpha*0.85)
			keyFont.DrawOrigin(batch, centerX, bottom+size*0.3, vector.Centre, size*0.3, true, strconv.Itoa(key.count))
		}
	}

	batch.ResetTransform()
}
