package stealth

import "github.com/go-rod/rod"

func HumanType(page *rod.Page, text string) error {
	for _, ch := range text {
		if err := page.InsertText(string(ch)); err != nil {
			return err
		}
		Delay(70, 160)
	}
	return nil
}
