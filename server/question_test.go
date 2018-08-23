package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alireza-Ta/GOASK/model"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuestion(t *testing.T) {

	testCases := map[string]model.Question{
		"BodyTitleTags": model.Question{
			Post:  model.Post{Body: "first question body"},
			Title: "First question title.",
			Tags: []*model.Tag{
				&model.Tag{
					Name: "Go",
				},
				&model.Tag{
					Name: "Test",
				},
			},
		},
	}

	for tc, q := range testCases {
		t.Run(tc, func(t *testing.T) {
			assert := assert.New(t)
			SetupSubTest()
			defer TeardownSubTest()

			res := httptest.NewRecorder()
			body, err := json.Marshal(q)
			if err != nil {
				t.Error("error in json parsing.")
			}
			req, err := http.NewRequest("POST", "/questions/", bytes.NewBuffer(body))
			if err != nil {
				t.Error("error in getting queestion from endpoint!!!")
			}

			TestServer.Router.ServeHTTP(res, req)

			var b model.Question
			err = json.Unmarshal(res.Body.Bytes(), &b)
			if err != nil {
				t.Error("error in json unmarshal.")
			}

			assert.Equal(200, res.Code, "got question.")
			assert.Equal(1, b.Id, "got id 1.")
			assert.Equal(q.Post, b.Post, "got body.")
			assert.Equal(q.Title, b.Title, "got title.")
			assert.Equal(q.Tags[0].Name, b.Tags[0].Name, "got first tag.")
			assert.Equal(q.Tags[1].Name, b.Tags[1].Name, "got second tag.")
		})

	}

}

func TestQuestionNotFound(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/questions/10000", nil)
	if err != nil {
		t.Error("Error in sending request.")
	}
	TestServer.Router.ServeHTTP(res, req)

	assert.Equal(t, 404, res.Code, "Question not found.")
}
