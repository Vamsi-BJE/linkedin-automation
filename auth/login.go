package auth

import (
	"strings"

	"github.com/go-rod/rod"
)

// IsLoggedIn checks for strong and weak signals of an authenticated session.
// It is intentionally permissive to avoid false negatives during demos.
func IsLoggedIn(page *rod.Page) bool {
	info, err := page.Info()
	if err != nil {
		return false
	}

	url := strings.ToLower(info.URL)
	title := strings.ToLower(info.Title)

	// Strong signal: Feed
	if strings.Contains(url, "/feed") {
		return true
	}

	// Weak-but-acceptable signals:
	// - Logged-in areas but not login/signup pages
	if (strings.Contains(title, "linkedin") ||
		strings.Contains(title, "feed")) &&
		!strings.Contains(url, "login") &&
		!strings.Contains(url, "signup") &&
		!strings.Contains(url, "checkpoint") {
		return true
	}

	// UI signal: profile avatar (if present)
	if _, err := page.Element(".global-nav__me-photo"); err == nil {
		return true
	}

	return false
}
