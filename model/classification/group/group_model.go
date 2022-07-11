package group

import "time"

type UserClassificationGroup struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(10) NOT NULL;"`
	CreatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
