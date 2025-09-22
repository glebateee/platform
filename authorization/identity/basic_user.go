package identity

import "strings"

type basicUser struct {
	Id            int
	Name          string
	Roles         []string
	Authenticated bool
}

var UnauthenticatedUser User = &basicUser{}

// GetDisplayName implements User.
func (b *basicUser) GetDisplayName() string {
	return b.Name
}

// GetID implements User.
func (b *basicUser) GetID() int {
	return b.Id
}

// InRole implements User.
func (b *basicUser) InRole(role string) bool {
	for _, r := range b.Roles {
		if strings.EqualFold(r, role) {
			return true
		}
	}
	return false
}

// IsAuthenticated implements User.
func (b *basicUser) IsAuthenticated() bool {
	return b.Authenticated
}

func NewBasicUser(id int, name string, roles ...string) User {
	return &basicUser{Id: id, Name: name, Roles: roles, Authenticated: true}
}
