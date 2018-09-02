package postgres

import "github.com/Alireza-Ta/GOASK/model"

func (s *Store) UserCreate(user *model.User) error {
	return s.DB.Insert(user)
}

func (s *Store) UserFind(username string) (*model.User, error) {
	u := new(model.User)
	err := s.DB.Model(u).Where("username = ?", username).Select()

	return u, err
}
