package movement

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// go get github.com/hajimehoshi/ebiten
var storeimg *ebiten.Image

func init() {
	var err error
	storeimg, _, err = ebitenutil.NewImageFromFile("image/button_store.png") // 48 x 28
	if err != nil {
		log.Fatal(err)
	}
}

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	playerX float64 = 0
	playerY float64 = 0
)

func update(screen *ebiten.Image) error {
	// Move the player
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerY += 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(playerX), float64(playerY))
	screen.DrawImage(playerImg, op)
}
