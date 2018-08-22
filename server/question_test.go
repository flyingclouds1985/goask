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

	for k, question := range testCases {
		assert := assert.New(t)
		t.Run(k, func(t *testing.T) {
			res := httptest.NewRecorder()
			body, err := json.Marshal(question)
			if err != nil {
				t.Error("error in json parsing.")
			}
			req, err := http.NewRequest("POST", "/questions/", bytes.NewBuffer(body))
			if err != nil {
				t.Error("error in getting queestion from endpoint!!!")
			}

			TestServer.Router.ServeHTTP(res, req)

			if k == "BodyTitleTags" {
				var b model.Question
				err := json.Unmarshal(res.Body.Bytes(), &b)
				if err != nil {
					t.Error("error in json unmarshal.")
				}

				assert.Equal(200, res.Code, "got question.")
				assert.Equal(testCases["BodyTitleTags"].Post, b.Post, "got body.")
				assert.Equal(testCases["BodyTitleTags"].Title, b.Title, "got title.")
				assert.Equal(testCases["BodyTitleTags"].Tags[0].Name, b.Tags[0].Name, "got first tag.")
				assert.Equal(testCases["BodyTitleTags"].Tags[1].Name, b.Tags[1].Name, "got second tag.")
			}
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
