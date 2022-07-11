package word

import "time"

type Word struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Word      string    `gorm:"type:varchar(20) NOT NULL;"`
	Weight    int32     `gorm:"type:varchar(100) NOT NULL;"`
	GroupID   int64     `gorm:"type:int(5);"`
	CreatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:time NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `user_classification_groups`
func (Word) TableName() string {
	return "user_classification_words"
}
