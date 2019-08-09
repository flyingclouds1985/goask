package postgres

import (
	"github.com/Alireza-Ta/goask/model"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (s *Store) FindUserByLoginCredentials(username, password string) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return nil, err
	}
	return u.Copy(), nil
}

// UpdateUserExcludePassword updates user but not password.
func (s *Store) UpdateUserExcludePassword(user *model.User) (int, error) {
	user.UpdatedAt = time.Now()
	res, err := s.DB.Model(user).ExcludeColumn("password").WherePK().Update()
	return res.RowsAffected(), err
}
