package postgres

import (
	"github.com/Alireza-Ta/GOASK/model"
)

// CreateUser insert new user.
func (s *Store) CreateUser(user *model.User) error {
	return s.DB.Insert(user)
}

// FindUser finds user based on id.
func (s *Store) FindUser(id int) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("id = ?", id).Select()

	return u, err
}

// FindUserByName finds user by username.
func (s *Store) FindUserByName(username string) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("username = ?", username).Select()

	return u, err
}

// UpdateUserExcludePassword updates user but not password.
func (s *Store) UpdateUserExcludePassword(user *model.User) (int, error) {
	res, err := s.DB.Model(user).ExcludeColumn("password").WherePK().Update()
	return res.RowsAffected(), err
}
