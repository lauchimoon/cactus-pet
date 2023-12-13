package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"

	e "github.com/hajimehoshi/ebiten/v2"
	"github.com/lauchimoon/cactus-pet/resources/images"
)

var (
	cactusImage *e.Image
)

type Game struct {
	CactusX   int
	CactusY   int
	CactusDir int

	CactusAnimX            int
	CactusAnimY            int
	CactusAnimFrameWidth   int
	CactusAnimFrameHeight  int
	CactusAnimFrameCounter int

	MonitorWidth  int
	MonitorHeight int
}

func init() {
	var err error
	img, err := png.Decode(bytes.NewReader(images.Cactus_png))
	if err != nil {
		log.Fatal(err)
	}
	cactusImage = e.NewImageFromImage(img)
}

func (g *Game) Update() error {
	g.CactusX += 2 * g.CactusDir
	if g.CactusX >= g.MonitorWidth-g.CactusAnimFrameWidth || g.CactusX <= 0 {
		g.CactusDir *= -1
	}

	// animation
	g.CactusAnimFrameCounter++
	if g.CactusAnimFrameCounter >= 10 {
		g.CactusAnimX = 128
	} else {
		g.CactusAnimX = 0
	}

	if g.CactusAnimFrameCounter >= 20 {
		g.CactusAnimFrameCounter = 0
	}

	e.SetWindowPosition(g.CactusX, g.CactusY)

	return nil
}

func (g *Game) Draw(screen *e.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0})
	rec := image.Rect(g.CactusAnimX, g.CactusAnimY, g.CactusAnimX+g.CactusAnimFrameWidth, g.CactusAnimFrameHeight)

	screen.DrawImage(cactusImage.SubImage(rec).(*e.Image), nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 128, 128
}

func main() {
	game := &Game{
		CactusX:                0,
		CactusDir:              1,
		CactusAnimX:            0,
		CactusAnimY:            0,
		CactusAnimFrameWidth:   128,
		CactusAnimFrameHeight:  128,
		CactusAnimFrameCounter: 0,
	}

	e.SetWindowSize(128, 128)
	e.SetWindowTitle("Hello world")
	e.SetWindowDecorated(false)
	e.SetWindowFloating(true)

	game.CactusDir = 1
	game.MonitorWidth, game.MonitorHeight = e.ScreenSizeInFullscreen()
	game.CactusY = game.MonitorHeight - 200
	if err := e.RunGameWithOptions(game, &e.RunGameOptions{ScreenTransparent: true}); err != nil {
		log.Fatal(err)
	}
}
