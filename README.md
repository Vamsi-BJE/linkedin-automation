# LinkedIn Automation System (Go + Rod)

## Overview

This project is a **LinkedIn automation system** built using **Golang** and the **Rod browser automation framework**.  
It demonstrates how to automate LinkedIn workflows such as **login, people search, profile interaction, connection logic, and messaging flow** while prioritizing **stealth, human-like behavior, and account safety**.

The automation is intentionally designed to **simulate actions without performing irreversible operations**, making it suitable for demonstrations, testing, and technical evaluation.

---

## Key Objectives

- Automate LinkedIn workflows using Go
- Avoid bot detection via stealth techniques
- Simulate realistic human behavior
- Enforce strict rate-limiting and safety constraints
- Demonstrate clean, extensible automation architecture

---

## Features Implemented

### 1. Secure Login Automation
- Credentials loaded from environment variables (`.env`)
- Human-like typing with randomized delays
- Bézier curve–based mouse movement (no straight-line motion)
- Stealth browser configuration to suppress automation fingerprints

**Why this matters:**  
Reduces bot detection risk and avoids hard-coded credentials.

---

### 2. Stealth & Anti-Detection Techniques
- Integration with `go-rod/stealth`
- Disabled automation indicators (e.g., `navigator.webdriver`)
- Natural mouse trajectories using Bézier curves
- Randomized delays between all actions

**Why this matters:**  
LinkedIn actively detects automation; stealth is mandatory.

---

### 3. People Search Automation
- Navigates to LinkedIn People Search
- Searches using keyword-based queries (e.g., “Recruiter”)
- Scrolls results to trigger lazy loading
- Mimics natural browsing behavior

**Why this matters:**  
Search and targeting are core functional requirements.

---

### 4. Profile Identification & Interaction
- Detects candidate profiles using `/in/` profile URLs
- Moves mouse to profile cards to demonstrate targeting logic
- Opens profiles safely for demonstration purposes

**Why this matters:**  
Shows the ability to identify and interact with real profiles.

---

### 5. Connection Request Logic (Simulation)
- Detects availability of **Connect / Follow** buttons
- Moves mouse toward actionable buttons
- **Does NOT send real connection requests**

**Why simulation only:**  
Prevents account risk while proving technical capability.

---

### 6. Messaging Workflow (Safe Demonstration)
- Navigates to LinkedIn Messaging
- Detects existing conversation threads
- Types follow-up messages using human-like input
- Gracefully skips messaging if no threads exist (new account)

**Why this matters:**  
Demonstrates follow-up and response-handling logic.

---

### 7. Rate Limiting & Execution Control
- Tracks execution runs using a persistent `state.json`
- Enforces a maximum number of runs per account
- Prevents aggressive or repeated automation

**Why this matters:**  
Protects accounts from throttling and bans.

---

## Project Structure

```text 
linkedin-automation/
│
├── main.go # Core automation flow
│
├── config/
│ └── config.go # Environment-based configuration
│
├── auth/
│ ├── login.go # Login state detection
│ └── captcha.go # CAPTCHA detection logic
│
├── stealth/
│ ├── mouse.go # Bézier mouse movement
│ ├── typing.go # Human typing simulation
│ ├── scroll.go # Natural scrolling
│ └── timing.go # Randomized delays
│
├── state/
│ └── store.go # Rate-limiting state
│
├── logx/
│ └── logger.go # Structured logging
│
├── .env.example # Environment variable template
├── state.json # Execution tracking
└── README.md


---

## Technology Stack

- **Language:** Go (Golang)
- **Browser Automation:** go-rod
- **Stealth Handling:** go-rod/stealth
- **Configuration:** godotenv
- **Browser:** Chromium (Rod-managed)

---

## Environment Setup

### 1. Install Dependencies
```bash
go mod tidy
```

**3. Configure Credentials**
Create a file named `.env` in the root directory:
```env
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_secure_password
```

## ▶️ Usage

**Run the automation:**
```bash
go run main.go
```

**Reset Rate Limit:**
If you hit the "Rate limit reached" error, simply reset the counter:
* **Windows:** `del state.json`
* **Mac/Linux:** `rm state.json`

## ⚠️ Disclaimer

This tool is for **educational purposes only**.
* **Simulation Mode:** The bot is currently configured to *simulate* clicks (hovering and typing) without actually sending connection requests or messages to prevent spam.
* **Account Safety:** Always use with caution. LinkedIn strictly prohibits aggressive automation. The built-in rate limiter is designed to help, but use at your own risk.
