package postgres

import (
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

// UserCreate insert new user.
func (s *Store) UserCreate(user *model.User) error {
	return s.DB.Insert(user)
}

// UserFind finds user based on id.
func (s *Store) UserFind(id int) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("id = ?", id).Select()

	return u, err
}

// UserFindByName finds user by username.
func (s *Store) UserFindByName(username string) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("username = ?", username).Select()

	return u, err
}

// UserUpdateExcludePassword updates user but not password.
func (s *Store) UserUpdateExcludePassword(user *model.User) (orm.Result, error) {
	return s.DB.Model(user).ExcludeColumn("password").WherePK().Update()
}
