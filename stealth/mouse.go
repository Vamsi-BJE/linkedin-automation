package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func HumanMove(page *rod.Page, targetX, targetY float64) {
	start := page.Mouse.Position()

	controlX := (start.X+targetX)/2 + rand.Float64()*100 - 50
	controlY := (start.Y+targetY)/2 + rand.Float64()*100 - 50

	steps := 20
	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		x := (1-t)*(1-t)*start.X + 2*(1-t)*t*controlX + t*t*targetX
		y := (1-t)*(1-t)*start.Y + 2*(1-t)*t*controlY + t*t*targetY

		page.Mouse.MoveTo(proto.Point{X: x, Y: y})
		time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)
	}
}
