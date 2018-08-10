package postgres

import (
	"fmt"
	"strings"

	"github.com/Alireza-Ta/GOASK/model"
)

type Tags []model.Tag

func (s *Store) TagCreate(tagsString string) (Tags, error) {
	tagsString = strings.TrimSpace(tagsString)
	names := strings.Split(tagsString, ",")
	tags := make(Tags, len(names))

	for k, name := range names {
		var tag model.Tag
		tag.Name = name
		// _, err := s.db.Model((*model.Tag)(nil)).Exec("SELECT setval('tags_id_seq', MAX(id)) FROM tags;")
		// if err != nil {
		// 	fmt.Println(err)

		// 	return tags, err
		// }
		fmt.Println(tag)
		res, err := s.db.Model(&tag).OnConflict("(name) DO NOTHING").SelectOrInsert()
		fmt.Println(res)
		if err != nil {
			fmt.Println(err)
			return tags, err
		}

		tags[k] = tag
	}
	fmt.Println(tags)
	return tags, nil
}
