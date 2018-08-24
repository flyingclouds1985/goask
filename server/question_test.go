package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alireza-Ta/GOASK/model"

	"github.com/stretchr/testify/assert"
)

var testCases = map[string]model.Question{
	"complete": model.Question{
		Post:  model.Post{Body: "This is the question body that must be more than 50 words till the API let us pass the this test nicely."},
		Title: "This is the question title.",
		Tags: []*model.Tag{
			&model.Tag{
				Name: "Go",
			},
			&model.Tag{
				Name: "Test",
			},
		},
	},
	"withoutTitle": model.Question{
		Post:  model.Post{Body: "This is the question body."},
		Title: "",
	},
	"minLengthTitle": model.Question{
		Post:  model.Post{Body: "This is the question body."},
		Title: "title min:15",
	},
	"withoutBody": model.Question{
		Title: "This is the question title.",
	},
	"minLengthBody": model.Question{
		Post:  model.Post{Body: "This is the question body less than 50 words."},
		Title: "This is the question title.",
	},
}

func TestCreateQuestion(t *testing.T) {
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
			if tc == "complete" {
				assert.Equal(200, res.Code, "got question.")
				assert.Equal(1, b.Id, "got id 1.")
				assert.Equal(q.Post, b.Post, "got body.")
				assert.Equal(q.Title, b.Title, "got title.")
				assert.Equal(q.Tags[0].Name, b.Tags[0].Name, "got first tag.")
				assert.Equal(q.Tags[1].Name, b.Tags[1].Name, "got second tag.")
			}
			if tc == "withoutTitle" || tc == "minLengthTitle" {
				assert.Equal(400, res.Code, "got question title validation error.")
			}
			if tc == "withoutBody" || tc == "minLengthBody" {
				assert.Equal(400, res.Code, "got question body validation error.")
			}
		})

	}

}

func TestQuestionNotFound(t *testing.T) {
	defer TeardownSubTest()

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/questions/10000", nil)
	if err != nil {
		t.Error("Error in sending request.")
	}
	TestServer.Router.ServeHTTP(res, req)

	assert.Equal(t, 404, res.Code, "Question not found.")
}

func TestGetQuestion(t *testing.T) {
	SetupSubTest()
	defer TeardownSubTest()

	oldRes := httptest.NewRecorder()
	body, err := json.Marshal(testCases["complete"])
	if err != nil {
		t.Error("Error in json marshal...")
	}
	req, err := http.NewRequest("POST", "/questions/", bytes.NewBuffer(body))
	if err != nil {
		t.Error("Error in posting question...")
	}
	TestServer.Router.ServeHTTP(oldRes, req)

	newRes := httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/questions/1/this-is-the-question-title", nil)
	TestServer.Router.ServeHTTP(newRes, req)

	fmt.Println(oldRes.Body)
	fmt.Println(newRes.Body)
}
