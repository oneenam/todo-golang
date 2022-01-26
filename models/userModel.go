package models

import "github.com/jinzhu/gorm"

/**
gorm.Model stores variables such as ID, CreatedAt,
and you don't need to add them separately.
*/

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
}
