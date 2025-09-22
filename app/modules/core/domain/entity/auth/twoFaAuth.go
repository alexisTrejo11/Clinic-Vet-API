package auth

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"strings"
	"time"
)

// TwoFactorAuth represents two-factor authentication settings for a user
type TwoFactorAuth struct {
	enabled     bool
	method      TwoFactorMethod
	secret      string
	backupCodes []string
	enabledAt   *time.Time
	lastUsedAt  *time.Time
}

// TwoFactorMethod represents the type of 2FA method
type TwoFactorMethod string

const (
	TwoFactorMethodTOTP    TwoFactorMethod = "totp"   // Time-based OTP
	TwoFactorMethodSMS     TwoFactorMethod = "sms"    // SMS codes
	TwoFactorMethodEmail   TwoFactorMethod = "email"  // Email codes
	TwoFactorMethodBackup  TwoFactorMethod = "backup" // Backup codes
	TwoFactorMethodUnknown TwoFactorMethod = "unknown"
)

// NewTwoFactorAuth creates a new TwoFactorAuth value object
func NewTwoFactorAuth(enabled bool, method TwoFactorMethod, secret string,
	backupCodes []string, enabledAt, lastUsedAt *time.Time,
) TwoFactorAuth {
	return TwoFactorAuth{
		enabled:     enabled,
		method:      method,
		secret:      secret,
		backupCodes: backupCodes,
		enabledAt:   enabledAt,
		lastUsedAt:  lastUsedAt,
	}
}

// NewDisabledTwoFactorAuth creates a disabled 2FA instance
func NewDisabledTwoFactorAuth() TwoFactorAuth {
	return TwoFactorAuth{
		enabled:     false,
		method:      TwoFactorMethodUnknown,
		secret:      "",
		backupCodes: nil,
		enabledAt:   nil,
		lastUsedAt:  nil,
	}
}

// GenerateNewTwoFactorAuth generates a new TOTP 2FA setup
func GenerateNewTwoFactorAuth() (TwoFactorAuth, error) {
	secret, err := generateTOTPSecret()
	if err != nil {
		return NewDisabledTwoFactorAuth(), err
	}

	backupCodes := generateBackupCodes(8)
	now := time.Now()

	return TwoFactorAuth{
		enabled:     true,
		method:      TwoFactorMethodTOTP,
		secret:      secret,
		backupCodes: backupCodes,
		enabledAt:   &now,
		lastUsedAt:  nil,
	}, nil
}

func (t TwoFactorAuth) Enabled() bool {
	return t.enabled
}

func (t TwoFactorAuth) Method() TwoFactorMethod {
	return t.method
}

func (t TwoFactorAuth) Secret() string {
	return t.secret
}

func (t TwoFactorAuth) BackupCodes() []string {
	return t.backupCodes
}

func (t TwoFactorAuth) EnabledAt() *time.Time {
	return t.enabledAt
}

func (t TwoFactorAuth) LastUsedAt() *time.Time {
	return t.lastUsedAt
}

// ValidateCode validates a 2FA code
func (t TwoFactorAuth) ValidateCode(code string) (bool, error) {
	if !t.enabled {
		return false, errors.New("2FA is not enabled")
	}

	if code == "" {
		return false, errors.New("code cannot be empty")
	}

	// Check if it's a backup code
	if t.isBackupCode(code) {
		return true, t.consumeBackupCode(code)
	}

	// Validate based on method
	switch t.method {
	case TwoFactorMethodTOTP:
		return t.validateTOTP(code)
	case TwoFactorMethodSMS:
		return t.validateSMS(code)
	case TwoFactorMethodEmail:
		return t.validateEmail(code)
	default:
		return false, fmt.Errorf("unsupported 2FA method: %s", t.method)
	}
}

// IsBackupCode checks if the code is a valid backup code
func (t TwoFactorAuth) IsBackupCode(code string) bool {
	return t.isBackupCode(code)
}

// HasBackupCodesRemaining checks if there are backup codes available
func (t TwoFactorAuth) HasBackupCodesRemaining() bool {
	return len(t.backupCodes) > 0
}

// RemainingBackupCodes returns the number of remaining backup codes
func (t TwoFactorAuth) RemainingBackupCodes() int {
	return len(t.backupCodes)
}

// GenerateNewBackupCodes generates new backup codes and replaces old ones
func (t *TwoFactorAuth) GenerateNewBackupCodes(count int) error {
	if !t.enabled {
		return errors.New("cannot generate backup codes for disabled 2FA")
	}

	if count < 1 || count > 20 {
		return errors.New("backup code count must be between 1 and 20")
	}

	t.backupCodes = generateBackupCodes(count)
	return nil
}

// Disable disables 2FA and clears all sensitive data
func (t *TwoFactorAuth) Disable() {
	t.enabled = false
	t.method = TwoFactorMethodUnknown
	t.secret = ""
	t.backupCodes = nil
	t.enabledAt = nil
	t.lastUsedAt = nil
}

// RecordUsage records when 2FA was last used
func (t *TwoFactorAuth) RecordUsage() {
	now := time.Now()
	t.lastUsedAt = &now
}

// Private helper methods

func (t TwoFactorAuth) isBackupCode(code string) bool {
	for _, backupCode := range t.backupCodes {
		if backupCode == code {
			return true
		}
	}
	return false
}

func (t *TwoFactorAuth) consumeBackupCode(code string) error {
	for i, backupCode := range t.backupCodes {
		if backupCode == code {
			// Remove the used backup code
			t.backupCodes = append(t.backupCodes[:i], t.backupCodes[i+1:]...)
			return nil
		}
	}
	return errors.New("backup code not found")
}

func (t TwoFactorAuth) validateTOTP(code string) (bool, error) {
	// In a real implementation, you would use a TOTP library like:
	// github.com/pquerna/otp/totp
	// This is a simplified version

	if len(code) != 6 {
		return false, errors.New("TOTP code must be 6 digits")
	}

	// Basic validation - in real implementation, use proper TOTP validation
	// For now, just return true for demonstration
	// actual validation would be: totp.Validate(code, t.secret)

	return true, nil
}

func (t TwoFactorAuth) validateSMS(code string) (bool, error) {
	// SMS code validation logic
	if len(code) != 6 {
		return false, errors.New("SMS code must be 6 digits")
	}

	// In real implementation, you would verify against stored/expected code
	return true, nil
}

func (t TwoFactorAuth) validateEmail(code string) (bool, error) {
	// Email code validation logic
	if len(code) != 6 {
		return false, errors.New("email code must be 6 digits")
	}

	// In real implementation, you would verify against stored/expected code
	return true, nil
}

// Utility functions

func generateTOTPSecret() (string, error) {
	// Generate 20 random bytes for TOTP secret
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate TOTP secret: %w", err)
	}

	// Base32 encode without padding
	secret := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
	return secret, nil
}

func generateBackupCodes(count int) []string {
	codes := make([]string, count)
	for i := 0; i < count; i++ {
		// Generate 8-character backup codes in format XXXX-XXXX
		part1 := randomString(4, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		part2 := randomString(4, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		codes[i] = fmt.Sprintf("%s-%s", part1, part2)
	}
	return codes
}

func randomString(length int, charset string) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // Should not happen in practice
	}

	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}

// TwoFactorMethod utilities

func (m TwoFactorMethod) IsValid() bool {
	switch m {
	case TwoFactorMethodTOTP, TwoFactorMethodSMS, TwoFactorMethodEmail, TwoFactorMethodBackup:
		return true
	default:
		return false
	}
}

func (m TwoFactorMethod) String() string {
	return string(m)
}

func ParseTwoFactorMethod(method string) (TwoFactorMethod, error) {
	normalized := strings.ToLower(strings.TrimSpace(method))
	switch normalized {
	case "totp":
		return TwoFactorMethodTOTP, nil
	case "sms":
		return TwoFactorMethodSMS, nil
	case "email":
		return TwoFactorMethodEmail, nil
	case "backup":
		return TwoFactorMethodBackup, nil
	default:
		return TwoFactorMethodUnknown, fmt.Errorf("invalid 2FA method: %s", method)
	}
}

func (t TwoFactorAuth) Validate() error {
	if t.enabled {
		if !t.method.IsValid() {
			return errors.New("invalid 2FA method when enabled")
		}
		if t.secret == "" && t.method != TwoFactorMethodBackup {
			return errors.New("secret is required for non-backup methods")
		}
		if t.enabledAt == nil {
			return errors.New("enabledAt timestamp is required when enabled")
		}
		if t.enabledAt.After(time.Now()) {
			return errors.New("enabledAt cannot be in the future")
		}
		if t.lastUsedAt != nil && t.lastUsedAt.After(time.Now()) {
			return errors.New("lastUsedAt cannot be in the future")
		}
	}
	return nil
}
