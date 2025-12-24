package stealth

import "github.com/go-rod/rod"

func RandomScroll(page *rod.Page) {
	page.Mouse.Scroll(0, 300, 5)
	Delay(300, 700)
	page.Mouse.Scroll(0, -150, 3)
}
