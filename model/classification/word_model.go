package classification

import "time"

type UserClassificationWord struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Word      string    `gorm:"type:varchar(20) NOT NULL;"`
	Weight    int32     `gorm:"type:varchar(100) NOT NULL;"`
	GroupID   int64     `gorm:"type:int(5);"`
	CreatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
