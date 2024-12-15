package auth

import (
	"errors"
	"time"
	
	"github.com/golang-jwt/jwt/v5" // JWT библиотека
)

// JWTSecret — секрет для подписи токенов (рекомендуется хранить в переменной окружения)
var JWTSecret = []byte("your-secret-key")

// AuthenticatedUser — структура данных авторизованного пользователя
type AuthenticatedUser struct {
	ID       int
	Username string
	Token    string
}

// GenerateJWT генерирует JWT токен для авторизованного пользователя
func GenerateJWT(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Токен действует 24 часа
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// ValidateJWT проверяет JWT токен и возвращает данные пользователя
func ValidateJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// RegisterUser регистрирует нового пользователя
func RegisterUser(username, password string) (*AuthenticatedUser, error) {
	// Хэшируем пароль с использованием HashPassword
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Сохраните username и hashedPassword в базу данных (пример)
	userID := 1 // Получите реальный userID после сохранения

	// Генерируем JWT токен
	token, err := GenerateJWT(userID, username)
	if err != nil {
		return nil, err
	}

	return &AuthenticatedUser{
		ID:       userID,
		Username: username,
		Token:    token,
	}, nil
}

// LoginUser аутентифицирует пользователя
func LoginUser(username, password string) (*AuthenticatedUser, error) {
	// Здесь получите хэш пароля из базы данных (пример)
	storedHashedPassword := "$2b$10$..." // Замените на реальный хэш из базы данных

	// Проверяем пароль с использованием VerifyPassword
	if !VerifyPassword(password, storedHashedPassword) {
		return nil, errors.New("invalid username or password")
	}

	// Получите userID из базы данных
	userID := 1 // Замените на реальный ID

	// Генерируем JWT токен
	token, err := GenerateJWT(userID, username)
	if err != nil {
		return nil, err
	}

	return &AuthenticatedUser{
		ID:       userID,
		Username: username,
		Token:    token,
	}, nil
}
