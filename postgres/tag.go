package postgres

import (
	"log"

	"github.com/Alireza-Ta/GOASK/model"
)

// TagCreate creates a tag.
func (s *Store) TagCreate(tags []*model.Tag, qid int) {
	for _, t := range tags {
		t.QuestionId = qid
		err := s.DB.Insert(t)
		if err != nil {
			log.Fatal("Error in inserting tag...", err)
		}
	}
}
