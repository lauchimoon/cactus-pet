package main

import (
    "log"
	e "github.com/hajimehoshi/ebiten/v2"
    "image/color"
    "image/png"
    "image"
    "github.com/lauchimoon/cactus-pet/resources/images"
    "bytes"
)

var (
    cactusImage *e.Image
    x int
    y int
    dir int

    animX int = 0
    animY int = 0
    frameWidth int = 128
    frameHeight int = 128
    frameCounter int = 0

    flipped bool = false

    monitorWidth int
    monitorHeight int
)

type Game struct {}

func init() {
    var err error
    img, err := png.Decode(bytes.NewReader(images.Cactus_png))
    if err != nil {
        log.Fatal(err)
    }
    cactusImage = e.NewImageFromImage(img)

    dir = 1
}

func (g *Game) Update() error {
    x += 2*dir
    if x >= monitorWidth - frameWidth || x <= 0 {
        dir *= -1
        flipped = !flipped
    }

    // animation
    frameCounter++
    if frameCounter >= 10 {
        animX = 128
    } else {
        animX = 0
    }

    if frameCounter >= 20 {
        frameCounter = 0
    }

    e.SetWindowPosition(x, y)

    return nil
}

func (g *Game) Draw(screen *e.Image) {
    screen.Fill(color.RGBA{0, 0, 0, 0})
    rec := image.Rect(animX, animY, animX + frameWidth, frameHeight)

    screen.DrawImage(cactusImage.SubImage(rec).(*e.Image), nil)
    //screen.DrawImage(cactusImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 128, 128
}

func main() {
    game := &Game{}
    e.SetWindowSize(128, 128)
    e.SetWindowTitle("Hello world")
    e.SetWindowDecorated(false)

    monitorWidth, monitorHeight = e.ScreenSizeInFullscreen()
    y = monitorHeight - 200
    if err := e.RunGameWithOptions(game, &e.RunGameOptions{ScreenTransparent: true, }); err != nil {
        log.Fatal(err)
    }
}
