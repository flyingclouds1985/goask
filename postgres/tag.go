package postgres

import (
	"fmt"
	"strings"

	"github.com/Alireza-Ta/GOASK/model"
)

type Tags []*model.Tag

func (s *Store) TagCreate(tagsString string) (Tags, error) {
	tagsString = strings.Trim(tagsString, " ")
	names := strings.Split(tagsString, ",")
	tags := make(Tags, len(names))

	for k, name := range names {
		tag := new(model.Tag)
		tag.Name = name
		_, err := s.db.Model((*model.Tag)(nil)).Exec("SELECT setval('tags_id_seq', MAX(id)) FROM tags;")
		if err != nil {
			fmt.Println(err)

			return tags, err
		}

		_, err = s.db.Model(tag).OnConflict("(name) DO NOTHING").Insert()
		if err != nil {
			fmt.Println(err)
			return tags, err
		}

		tags[k] = tag
	}

	return tags, nil
}
