package auth

import "github.com/go-rod/rod"

// HasLoginError detects visible login errors (wrong email/password).
// It does NOT detect CAPTCHA or checkpoints.
func HasLoginError(page *rod.Page) bool {
	// LinkedIn uses role="alert" for credential errors
	_, err := page.Element("div[role='alert']")
	return err == nil
}
