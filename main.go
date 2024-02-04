package main

import (
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	x, y float64
	gs GameStage
}

func (g *Game) init(width int, height int) {
	g.gs.init(width, height)
}

func (g *Game) Update() error {
	g.gs.clearField()

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.y -= 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.y += 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.gs.MoveLeft()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.gs.MoveRight()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gs.operatingTetrimino.Rotate()
	}

	g.gs.putTetrimino()
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DrawRect(screen, g.x, g.y, 16, 16, color.RGBA{0x80, 0x80, 0x80, 0xff})

	g.gs.DrawField(screen)
}

func (gs *GameStage) DrawField(screen *ebiten.Image) {
	for i := 0; i < gs.height; i++ {
		for j := 0; j < gs.width; j++ {
			if gs.field[i][j] == 0 {
				// draw stroke rect
				ebitenutil.DrawRect(screen, float64(j * 16), float64(i * 16), 16, 16, color.RGBA{0x00, 0x00, 0x00, 0xff})
				ebitenutil.DrawRect(screen, float64(j * 16), float64(i * 16), 16, 1, color.RGBA{0xff, 0xff, 0xff, 0xff})
				ebitenutil.DrawRect(screen, float64(j * 16), float64(i * 16), 1, 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
			} else {
				ebitenutil.DrawRect(screen, float64(j * 16), float64(i * 16), 16, 16, color.RGBA{0xff, 0xff, 0xff, 0xff})
			}
		}
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}


func main() {
	game := &Game{
		x: 0,
		y: 0,
	}

	game.init(15, 15)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}