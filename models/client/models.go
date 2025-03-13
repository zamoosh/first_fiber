package client

import (
	"fmt"
	"strings"

	"first_fiber/models/agency"
)

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique;not null"`
	Password    string
	Cellphone   string `gorm:"unique;not null"`
	FirstName   string
	LastName    string
	IsStaff     string
	IsSuperuser string
	AgencySet   []agency.Agency
}

func (User) TableName() string {
	return "client_user"
}

func (u User) String() (fullName string) {
	fullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return strings.Replace(fullName, " ", "", 1)
}
