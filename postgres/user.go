package postgres

import "github.com/Alireza-Ta/GOASK/model"

func (s *Store) UserCreate(user *model.User) error {
	return s.DB.Insert(user)
}
