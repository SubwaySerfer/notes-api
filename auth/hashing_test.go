package security

import "testing"

func TestHashPassword(t *testing.T) {
    password := "mysecurepassword"
    hash, err := HashPassword(password)
    if err != nil {
        t.Fatalf("Error hashing password: %v", err)
    }

    if !VerifyPassword(password, hash) {
        t.Errorf("Password verification failed")
    }
}
