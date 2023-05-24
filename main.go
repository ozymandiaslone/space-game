package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	maxRadius = 50
)

var (
	button       = startButton{w: 100, h: 36, x: (screenWidth / 2) - 50, y: (screenHeight / 2) - 18}
	screenWidth  = float64(1024)
	screenHeight = float64(768)
	gameStarted  = false
	system       = solarSystem{
		planets: []planet{
			planet{
				x: rand.Float64() * screenWidth,
				y: rand.Float64() * screenHeight,
				r: rand.Float64() * maxRadius,
			},
		},
		stars: []Star{
			Star{
				x: rand.Float64() * screenWidth,
				y: rand.Float64() * screenHeight,
				r: rand.Float64() * maxRadius,
			},
		},
	}
	camera = camera{
		mapX:  0,
		mapY:  0,
		scale: 1,
	}
)

type Game struct{}

type camera struct {
	mapX, mapY float64
	scale      float64
}

func (c *camera) Translate(dx, dy float64) {
	c.mapX += dx
	c.mapY += dy
}

func (c *camera) Scale(s float64) {
	c.scale *= s
}

func (c *camera) Draw(loadedEntities []entity) {

}

type startButton struct {
	w, h float64
	x, y float64
}

func (b *startButton) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x, b.y, b.w, b.h, color.RGBA{255, 255, 255, 255})
}

func (b *startButton) IsClicked(x, y int) bool {
	fx := float64(x)
	fy := float64(y)
	return fx > b.x && fx < b.x+b.w && fy > b.y && fy < b.y+b.h
}

type Star struct {
	x, y float64
	r    float64
}

type planet struct {
	x, y float64
	r    float64
}

type solarSystem struct {
	planets []planet
	stars   []Star
}

func (s *solarSystem) Draw(screen *ebiten.Image) {
	for _, p := range s.planets {
		drawCircle(screen, p.x, p.y, p.r)
	}
	for _, s := range s.stars {
		drawCircle(screen, s.x, s.y, s.r)
	}
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	if !gameStarted && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if button.IsClicked(mouseX, mouseY) {
			// Do things to start game
			gameStarted = true
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !gameStarted {
		button.Draw(screen)
	} else {
		system.Draw(screen)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func drawCircle(screen *ebiten.Image, x, y, radius float64) {
	minAngle := math.Acos(1 - 1/radius)
	for angle := float64(0); angle < 360; angle += minAngle {
		xDelta := radius * math.Cos(angle)
		yDelta := radius * math.Sin(angle)
		x1 := int(math.Round(x + xDelta))
		y1 := int(math.Round(y + yDelta))
		screen.Set(x1, y1, color.RGBA{255, 255, 255, 255})
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Test - Space Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
