package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

/**
gorm.Model stores variables such as ID, CreatedAt,
and you don't need to add them separately.
*/

type Task struct {
	gorm.Model
	Ttile       string
	Completed   bool
	CompletedAt time.Time
}
