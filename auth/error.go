package auth

import "github.com/go-rod/rod"

func HasLoginError(page *rod.Page) bool {
	_, err := page.Element("div[role='alert']")
	return err == nil
}
