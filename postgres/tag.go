package postgres

import (
	"log"

	"github.com/Alireza-Ta/goask/model"
)

// CreateTag creates a tag.
func (s *Store) CreateTag(tags []*model.Tag, qid int) {
	for _, t := range tags {
		t.QuestionId = qid
		err := s.DB.Insert(t)
		if err != nil {
			log.Fatal("Error in inserting tag...", err)
		}
	}
}
