package user

import (
	"fmt"
	"strings"
)

type ClientUser struct {
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Cellphone string `gorm:"unique;not null"`
	FirstName string
	LastName  string
}

func (ClientUser) TableName() string {
	return "client_user"
}

func (u ClientUser) String() (fullName string) {
	fullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return strings.Replace(fullName, " ", "", 1)
}
