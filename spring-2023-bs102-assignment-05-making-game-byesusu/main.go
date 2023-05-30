package main

import (
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var storeimg *ebiten.Image
var exitimg *ebiten.Image
var backgroundimg *ebiten.Image
var guideimg *ebiten.Image
var fieldimg *ebiten.Image
var field [5][8]bool

func init() {
	var err error
	storeimg, _, err = ebitenutil.NewImageFromFile("image/button_store.png") // 48 x 28
	if err != nil {
		log.Fatal(err)
	}

	backgroundimg, _, err = ebitenutil.NewImageFromFile("image/background.png")
	if err != nil {
		log.Fatal(err)
	}
	exitimg, _, err = ebitenutil.NewImageFromFile("image/exit.png")
	if err != nil {
		log.Fatal(err)
	}
	guideimg, _, err = ebitenutil.NewImageFromFile("image/guide.png")
	if err != nil {
		log.Fatal(err)
	}
	fieldimg, _, err = ebitenutil.NewImageFromFile("image/field.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	stage int
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if g.stage == 0 && x >= 20 && x < 68 && y >= 10 && y < 38 {
			g.stage = 1
		} else if g.stage == 1 && x >= screenWidth-16 && x < screenWidth && y >= 0 && y < 16 {
			g.stage = 0
		} else if g.stage == 0 && x >= 88 && x < 136 && y >= 10 && y < 38 { // guideimg를 클릭한 경우
			g.stage = 2
		} else if g.stage == 2 && x >= screenWidth-16 && x < screenWidth && y >= 0 && y < 16 {
			g.stage = 0
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(backgroundimg, op)
	if g.stage == 0 {
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(20, 10)
		screen.DrawImage(storeimg, op2)
		op3 := &ebiten.DrawImageOptions{}
		op3.GeoM.Translate(88, 10) // storeimg로부터 20px 옆에 위치하도록 함
		screen.DrawImage(guideimg, op3)

		for i := 0; i < 5; i++ {
			for j := 0; j < 8; j++ {
				if !field[i][j] {
					op4 := &ebiten.DrawImageOptions{}
					op4.GeoM.Translate(float64(j*80), float64(i*80)+48)
					screen.DrawImage(fieldimg, op4)
				}
			}
		}

	} else if g.stage == 1 {
		op4 := &ebiten.DrawImageOptions{}
		op4.GeoM.Translate(screenWidth-48, 0)
		screen.DrawImage(exitimg, op4)
		ebitenutil.DebugPrint(screen, "store")
	} else if g.stage == 2 {
		op4 := &ebiten.DrawImageOptions{}
		op4.GeoM.Translate(screenWidth-48, 0)
		screen.DrawImage(exitimg, op4)
		ebitenutil.DebugPrint(screen, "guide")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update1() error {
	// button corner pixel info
	// left top: screenWidth/2-24, screenHeight/2-14
	// right top: screenWidth/2+24, screenHeight/2-14
	// left bottom: screenWidth/2-24, screenHeight/2+14
	// right bottom: screenWidth/2+24, screenHeight/2+14
	var x, y int
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y = ebiten.CursorPosition()
	}

	if g.stage == 1 || g.stage == 2 {
		if screenWidth/2-24 <= x && x <= screenWidth/2+24 {
			if screenHeight/2-14 <= y && y <= screenHeight/2+14 {
				g.stage = 0 // Set g.stage to 0 when exitimg is clicked and g.stage is already 1
			}
		}
	}

	return nil
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
