package postgres

import (
	"github.com/Alireza-Ta/goask/model"
	"github.com/go-pg/pg"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type DBError struct {
	message string
}

func (e DBError) Error() string {
	return e.message
}

// CreateUser insert new user.
func (s *Store) CreateUser(user *model.User) error {
	err := s.DB.Insert(user)
	pgErr, ok := err.(pg.Error)
	// check for duplicate values.
	if ok && pgErr.IntegrityViolation() {
		column := pgErr.Field('n')
		switch column {
		case "users_username_key":
			return DBError{"This username exists. Please choose another one!"}
		case "users_email_key":
			return DBError{"This email exists. Please choose another one!"}
		}
	}

	return err
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
