package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"guardian-grid-api/internal/platform/database"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func IsFirstRun() bool {
	db := database.GetSQLite()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		panic(err)
	}

	return count == 0
}

func RunInitialSetup() {
	fmt.Println("🔐 GuardianGrid First-Time Setup")
	fmt.Println("--------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Generate TOTP
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "GuardianGrid",
		AccountName: username,
	})
	if err != nil {
		panic(err)
	}

	secret := key.Secret()
	otpURL := key.URL()

	fmt.Println("📲 Scan this QR code with Google Authenticator:")

	// Generate QR in terminal
	qr, err := qrcode.New(otpURL, qrcode.Medium)
	if err != nil {
		panic(err)
	}

	fmt.Println(qr.ToString(false)) // prints QR in terminal

	fmt.Println("\nOr manually enter secret:", secret)

	// Verify once
	fmt.Print("\nEnter TOTP code to confirm setup: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	if !totp.Validate(code, secret) {
		fmt.Println("❌ Invalid TOTP. Setup failed.")
		os.Exit(1)
	}

	// Save user
	saveUser(username, secret)

	fmt.Println("✅ Setup complete! You can now log in.")
}

func saveUser(username, secret string) {
	db := database.GetSQLite()

	_, err := db.Exec(
		"INSERT INTO users (username, totp_secret) VALUES (?, ?)",
		username,
		secret,
	)

	if err != nil {
		panic(err)
	}
}
