package postgres

import "time"

type Author struct {
	Id             int
	Username       string
	Firstname      string
	Lastname       string
	MembershipDate time.Time `sql:"default:now()"`
	Bio            string
}
