package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	rodstealth "github.com/go-rod/stealth"

	"linkedin-automation/config"
	"linkedin-automation/logx"
	"linkedin-automation/state"
	"linkedin-automation/stealth"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.Load()
	if cfg.Email == "" || cfg.Password == "" {
		log.Fatal("Missing credentials")
	}

	st := state.Load()
	if st.Runs >= cfg.MaxRuns {
		log.Fatal("Rate limit reached")
	}

	logx.Logger.Println("Starting browser...")
	u := launcher.New().
		Headless(false).
		Leakless(false).
		Set("remote-debugging-port", "0").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.Close()

	page := rodstealth.MustPage(browser)

	logx.Logger.Println("Opening LinkedIn login page...")
	page.MustNavigate("https://www.linkedin.com/login")
	time.Sleep(3 * time.Second)

	page.MustElement("#username").MustClick()
	stealth.HumanType(page, cfg.Email)
	stealth.Delay(300, 500)

	page.MustElement("#password").MustClick()
	stealth.HumanType(page, cfg.Password)
	stealth.Delay(400, 700)

	btn := page.MustElement("button[type=submit]")
	box := btn.MustShape().Box()
	stealth.HumanMove(page, box.X+10, box.Y+10)
	btn.MustClick()

	logx.Logger.Println("Login clicked. Continuing demo flow...")
	time.Sleep(4 * time.Second)

	logx.Logger.Println("Navigating to Search page (same tab)...")
	page.MustNavigate("https://www.linkedin.com/search/results/people/?keywords=Recruiter")
	time.Sleep(4 * time.Second)

	stealth.RandomScroll(page)
	time.Sleep(1 * time.Second)

	logx.Logger.Println("Locating profile...")
	profiles, err := page.Elements("a[href*='/in/']")
	if err != nil || len(profiles) == 0 {
		logx.Logger.Println("No profiles found. Ending demo.")
		return
	}

	profile := profiles[0]
	pBox := profile.MustShape().Box()
	stealth.HumanMove(page, pBox.X+10, pBox.Y+10)
	time.Sleep(500 * time.Millisecond)
	profile.MustClick()

	time.Sleep(4 * time.Second)

	logx.Logger.Println("Checking for Connect button...")
	buttons, _ := page.Elements("button")
	foundConnect := false

	for _, b := range buttons {
		txt, _ := b.Text()
		if strings.Contains(txt, "Connect") {
			bBox := b.MustShape().Box()
			stealth.HumanMove(page, bBox.X+5, bBox.Y+5)
			fmt.Println("Simulating Connect click (NOT sent)")
			foundConnect = true
			break
		}
	}

	if !foundConnect {
		logx.Logger.Println("Connect not available (Follow / Message only profile).")
	}

	logx.Logger.Println("Navigating to Messaging (same tab)...")
	page.MustNavigate("https://www.linkedin.com/messaging/")
	time.Sleep(4 * time.Second)

	threads, err := page.Elements(".msg-conversation-listitem")
	if err == nil && len(threads) > 0 {
		thread := threads[0]
		tBox := thread.MustShape().Box()
		stealth.HumanMove(page, tBox.X+10, tBox.Y+10)
		thread.MustClick()
		time.Sleep(2 * time.Second)

		if input, err := page.Element("div[contenteditable='true']"); err == nil {
			input.MustClick()
			stealth.HumanType(page, "Hello! Just following up.")
			fmt.Println("Follow-up message typed (NOT sent)")
			time.Sleep(1 * time.Second)
		}
	} else {
		logx.Logger.Println("No message threads found (New Account).")
	}

	st.Runs++
	state.Save(st)

	logx.Logger.Println("Demo complete. Closing in 5 seconds.")
	time.Sleep(5 * time.Second)
}
