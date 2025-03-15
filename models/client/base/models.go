package base

import (
	"fmt"
	"strings"
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

type AbstractUser struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique;not null"`
	Password    string
	Cellphone   string `gorm:"unique;not null"`
	FirstName   string
	LastName    string
	IsStaff     bool
	IsSuperuser bool
}

func (AbstractUser) TableName() string {
	return "client_user"
}

func (u AbstractUser) String() (fullName string) {
	fullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return strings.Replace(fullName, " ", "", 1)
}

func (u AbstractUser) GetId() uint {
	return u.Id
}

func (u AbstractUser) GetUsername() string {
	return u.Username
}

func (u AbstractUser) GetPassword() string {
	return u.Password
}

func (u AbstractUser) GetCellphone() string {
	return u.Cellphone
}

func (u AbstractUser) GetFirstname() string {
	return u.FirstName
}

func (u AbstractUser) GetLastname() string {
	return u.LastName
}

func (u AbstractUser) GetIsStaff() bool {
	return u.IsStaff
}

func (u AbstractUser) GetIsSuperuser() bool {
	return u.IsSuperuser
}
