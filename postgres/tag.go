package postgres

import (
	"log"

	"github.com/Alireza-Ta/GOASK/model"
)

type Tags []model.Tag

// TagCreate creates a tag.
func (s *Store) TagCreate(tags Tags, qid int) {
	for _, t := range tags {
		t.QuestionId = qid
		err := s.db.Insert(&t)
		if err != nil {
			log.Fatal("Error in inserting tag...", err)
		}
	}
}
