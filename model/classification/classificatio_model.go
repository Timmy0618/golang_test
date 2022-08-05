package classification

import (
	"time"
)

type Classification struct {
	UserId    int64     `gorm:"type:int(10) NOT NULL; primary_key;" json:"user_id"`
	GroupId   int64     `gorm:"type:int(10) NOT NULL; primary_key;" json:"group_id"`
	Score     int64     `gorm:"type:int(10) NOT NULL;" json:"score"`
	CreatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `user_classification_groups`
func (Classification) TableName() string {
	return "user_classification"
}
