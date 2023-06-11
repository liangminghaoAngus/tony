package main

import (
	"image"
	"log"
	"sync"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Game struct {
	view *furex.View
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
	// return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	g.view.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.view.Draw(screen)
}

//go:embed onBoard.html
var html string

func main() {

	once := sync.Once{}

	view := furex.Parse(html, &furex.ParseOptions{
		Width:  1280,
		Height: 720,
		Components: map[string]furex.Component{
			"button": func() *furex.View {
				return &furex.View{}
			},
		},
		Handler: furex.NewHandler(furex.HandlerOpts{
			Update: func(v *furex.View) {
				once.Do(func() {
					println("call update")
				})
			},
			Draw: func(screen *ebiten.Image, frame image.Rectangle, v *furex.View) {
				println("call draw")
			},
		}),
	})

	gameItem := &Game{
		view: view,
	}

	if err := ebiten.RunGame(gameItem); err != nil {
		log.Fatal(err)
	}
}
