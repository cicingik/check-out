package database

import (
	"fmt"
	"time"
)

type (
	Model struct {
		ID        int        `gorm:"Column:id;primary_key" json:"id"`
		CreatedAt JSONTime   `gorm:"Column:created_at;default=NOW()" json:"created_at"`
		UpdatedAt JSONTime   `gorm:"Column:updated_at;default=NOW()" json:"updated_at"`
		DeletedAt *time.Time `gorm:"Column:deleted_at;default:NULL" sql:"index" json:"deleted_at"`
	}

	JSONTime time.Time
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05Z"))
	return []byte(stamp), nil
}
