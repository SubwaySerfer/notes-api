package storage

import (
	_ "errors"
	_ "gorm.io/gorm"
	"notes-api/models"
)

func (d *Database) CreateUser(user models.User) error {
	if err := d.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}