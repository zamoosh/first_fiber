package agency

import (
	"time"

	"first_fiber/models/client/base"
)

type User interface {
	GetId() uint
	GetUsername() string
	GetPassword() string
	GetCellphone() string
	GetFirstname() string
	GetLastname() string
	GetIsStaff() bool
	GetIsSuperuser() bool
}

type Agency struct {
	Id        uint `gorm:"primaryKey;index"`
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	Deleted   *time.Time
	UserID    *uint             `gorm:"index"`
	User      base.AbstractUser `gorm:"foreignKey:UserID"`
}

func (Agency) TableName() string {
	return "agency_agency"
}
