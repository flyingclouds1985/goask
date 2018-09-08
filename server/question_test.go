package server

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/Alireza-Ta/GOASK/model"

	"github.com/stretchr/testify/assert"
)

var questionTestCases = map[string]model.Question{
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

func TestPostQuestion(t *testing.T) {
	for tc, q := range questionTestCases {
		t.Run(tc, func(t *testing.T) {
			assert := assert.New(t)
			SetupSubTest()
			defer TeardownSubTest()

			body, err := json.Marshal(q)
			checkNil(err, " question: error in json parsing.")

			res := makeRequest("POST", "/questions/", bytes.NewBuffer(body), nil)

			var b model.Question
			err = json.Unmarshal(res.Body.Bytes(), &b)
			checkNil(err, " question: error in json unmarshal.")

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

	res := makeRequest("GET", "/questions/10000", nil, nil)
	assert.Equal(t, 404, res.Code, "Question not found.")
}

func TestGetQuestion(t *testing.T) {
	SetupSubTest()
	defer TeardownSubTest()

	body, err := json.Marshal(questionTestCases["complete"])
	checkNil(err, " question: error in json marshal.")
	oldRes := makeRequest("POST", "/questions/", bytes.NewBuffer(body), nil)
	redirectRes := makeRequest("GET", "/questions/1", nil, nil)
	location := redirectRes.Header().Get("Location")

	newRes := makeRequest("GET", location, nil, nil)

	assert.Equal(t, oldRes.Body.String(), newRes.Body.String(), "got question")
}

func TestPatchQuestion(t *testing.T) {
	assert := assert.New(t)
	SetupSubTest()
	defer TeardownSubTest()

	body, err := json.Marshal(questionTestCases["complete"])
	checkNil(err, " question: error in json marshal.")

	res := makeRequest("POST", "/questions/", bytes.NewBuffer(body), nil)
	res = makeRequest("GET", "/questions/1/this-is-the-question-title", nil, nil)

	var b model.Question
	err = json.Unmarshal(res.Body.Bytes(), &b)
	checkNil(err, " question: error in json unmarshal.")

	b.Id = 1
	b.Title = "this is the question title that is more than 15 words length."
	b.Answered = 1
	b.Body = "Edited : This is the question body that must be more than 50 words till the API let us pass the this test nicely."
	// To have different timestamps.
	time.Sleep(1 * time.Second)

	body, err = json.Marshal(b)
	checkNil(err, " question: error in json marshal.")

	res = makeRequest("PATCH", "/questions/1", bytes.NewBuffer(body), nil)

	var rb model.Question
	err = json.Unmarshal(res.Body.Bytes(), &rb)
	checkNil(err, " question: error in json unmarshal.")

	assert.Equal(b.Title, rb.Title, "got title")
	assert.Equal(b.Body, rb.Body, "got body")
	assert.Equal(b.Answered, rb.Answered, "got answered 1.")
	assert.NotEqual(b.UpdatedAt, rb.UpdatedAt, "different update time.")
}
