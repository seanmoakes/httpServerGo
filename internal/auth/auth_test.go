package auth

import (
	"testing"
)

// TestHashPassword attempts to generate a hash for a
// password and fails if any error exists.
func TestHashPassword(t *testing.T) {
	testPass := "my_Passw0rd!"
	testHash, err := HashPassword(testPass)
	if err != nil {
		t.Fatalf("password hashing failed: %v", err)
	}

	err = CheckPassword(testPass, testHash)
	if err != nil {
		t.Fatalf("password checking failed: %v", err)
	}
}
