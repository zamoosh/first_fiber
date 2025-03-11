package agency

import "time"

type Agency struct {
	Id      uint `gorm:"primaryKey;index"`
	Name    string
	Active  bool
	Deleted *time.Time
	UserID  uint `gorm:"index"`
}

func (Agency) TableName() string {
	return "agency_agency"
}
