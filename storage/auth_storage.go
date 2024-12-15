package storage

import (
	"errors"
	"gorm.io/gorm"
	"notes-api/models"
	"notes-api/auth"
)

func (d *Database) CreateUser(user models.User) error {
	if err := d.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := d.Conn.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (d *Database) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := d.GetUserByUsername(username)
	if err != nil || user == nil {
		return nil, errors.New("Invalid username or password")
	}

	if !security.VerifyPassword(password, user.PasswordHash) {
		return nil, errors.New("Invalid username or password")
	}

	return &user, nil
}