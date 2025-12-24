package auth

import (
	"strings"

	"github.com/go-rod/rod"
)

func IsLoggedIn(page *rod.Page) bool {
	info, err := page.Info()
	if err != nil {
		return false
	}

	url := strings.ToLower(info.URL)
	title := strings.ToLower(info.Title)

	if strings.Contains(url, "/feed") {
		return true
	}

	if (strings.Contains(title, "linkedin") ||
		strings.Contains(title, "feed")) &&
		!strings.Contains(url, "login") &&
		!strings.Contains(url, "signup") &&
		!strings.Contains(url, "checkpoint") {
		return true
	}

	if _, err := page.Element(".global-nav__me-photo"); err == nil {
		return true
	}

	return false
}
