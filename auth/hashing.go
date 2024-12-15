package auth

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword создает хэш пароля
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}

// VerifyPassword проверяет пароль на совпадение с хэшем
func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
