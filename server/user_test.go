package server

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/stretchr/testify/assert"
)

var userTestCases = map[string]*model.User{
	"complete": &model.User{
		Id:       1,
		Username: "Tommy",
		Email:    "Tommy25@example.com",
		Password: "secretpassword",
		Bio:      "I'm a new user.",
	},
	"broken": &model.User{
		Id:       1,
		Username: "John",
		Email:    "John@example.com",
		Password: "secret",
		Bio:      "I'm a new user.",
	},
}

func TestPostUser(t *testing.T) {
	for tc, u := range userTestCases {
		t.Run(tc, func(t *testing.T) {
			SetupSubTest()
			defer TeardownSubTest()

			body, err := json.Marshal(u)
			checkNil(err, "user: err in json parsing.")

			res := makeRequest("POST", "/users/", bytes.NewBuffer([]byte(body)))

			var b model.User
			err = json.Unmarshal(res.Body.Bytes(), &b)
			checkNil(err, " user: err in json unmarshal.")

			if res.Code == 400 {
				var e map[string]map[string]interface{}
				err = json.Unmarshal(res.Body.Bytes(), &e)
				checkNil(err, " user: err in json unmarshal.")

				assert.Equal(t, "400", e["errors"]["status"], "got errors.")
			} else {
				assert.Equal(t, 200, res.Code, "user created.")
				assert.Equal(t, u.ExcludeTimes(), b.ExcludeTimes(), "got user.")
			}
		})
	}
}

func TestUserNotFound(t *testing.T) {
	defer TeardownSubTest()

	res := makeRequest("GET", "/users/russ", nil)
	assert.Equal(t, 404, res.Code, "User not found.")
}

func TestGetUser(t *testing.T) {
	SetupSubTest()
	defer TeardownSubTest()

	body, err := json.Marshal(userTestCases["complete"])
	checkNil(err, " user: error in json marshal.")
	oldRes := makeRequest("POST", "/users/", bytes.NewBuffer([]byte(body)))
	newRes := makeRequest("GET", "/users/Tommy", nil)

	assert.Equal(t, oldRes.Body, newRes.Body, "got user.")
}

func PatchUser(t *testing.T) {
	SetupSubTest()
	defer TeardownSubTest()

	body, err := json.Marshal(questionTestCases["complete"])
	checkNil(err, " question: error in json marshal.")
	oldRes := makeRequest("POST", "/questions/", bytes.NewBuffer(body))
	redirectRes := makeRequest("GET", "/questions/1", nil)
	location := redirectRes.Header().Get("Location")

	newRes := makeRequest("GET", location, nil)

	assert.Equal(t, oldRes.Body.String(), newRes.Body.String(), "got question")
}

func TestPatchUser(t *testing.T) {
	assert := assert.New(t)
	SetupSubTest()
	defer TeardownSubTest()

	body, err := json.Marshal(userTestCases["complete"])
	checkNil(err, " user: error in json marshal.")

	res := makeRequest("POST", "/users/", bytes.NewBuffer(body))
	res = makeRequest("GET", "/users/Tommy", nil)

	var b model.User
	err = json.Unmarshal(res.Body.Bytes(), &b)
	checkNil(err, " user: error in json unmarshal.")

	b.Id = 1
	b.Username = "Tommy2"
	b.Email = "Tommy2@example.com"
	b.Bio = "This is my new bio."
	// To have different timestamps.
	time.Sleep(1 * time.Second)

	body, err = json.Marshal(b)
	checkNil(err, " user: error in json marshal.")

	res = makeRequest("PATCH", "/users/1", bytes.NewBuffer(body))

	var rb model.User
	err = json.Unmarshal(res.Body.Bytes(), &rb)
	checkNil(err, " user: error in json unmarshal.")

	assert.Equal(b.Username, rb.Username, "got username")
	assert.Equal(b.Email, rb.Email, "got email")
	assert.Equal(b.Bio, rb.Bio, "got bio.")
	assert.Equal(b.CreatedAt, rb.CreatedAt, "got created_at.")
	assert.NotEqual(b.UpdatedAt, rb.UpdatedAt, "different update time.")
}
