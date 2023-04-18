package planner

import (
	"strings"

	"github.com/oriser/regroup"
)

var (
	addUserRegexp = regroup.MustCompile(`(?i)^add\s+user\s+(?P<username>\w+)\s+password\s+(?P<password>\w+)$`)
)

type AddUserGroup struct {
	Username string `regroup:"username"`
	Password string `regroup:"password"`
}

func checkAddUser(query string) (username, password string, err error) {
	elem := &AddUserGroup{}
	err = addUserRegexp.MatchToTarget(strings.TrimSpace(query), elem)
	if err != nil {
		return
	}

	username = elem.Username
	password = elem.Password
	err = nil

	return
}
