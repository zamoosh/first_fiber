package client

import (
	"fmt"
	"strings"

	"first_fiber/models/agency"
	"first_fiber/models/client/base"
)

type User struct {
	base.AbstractUser
	AgencySet []agency.Agency
}

func (User) TableName() string {
	return "client_user"
}

func (u User) String() (fullName string) {
	fullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return strings.Replace(fullName, " ", "", 1)
}
