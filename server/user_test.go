package server

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/stretchr/testify/assert"
)

var userTestCases = map[string]*model.User{
	"complete": &model.User{
		Id:       1,
		Username: "John25",
		Email:    "john25@example.com",
		Password: "secretpassword",
		Bio:      "I'm a new user.",
	},
	// "broken": &model.User{
	// 	Id:       1,
	// 	Username: "John",
	// 	Email:    "John@example.com",
	// 	Password: "secret",
	// 	Bio:      "I'm a new user.",
	// },
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

			assert.Equal(t, 200, res.Code, "user created.")
			assert.Equal(t, u.ExcludeTimes(), b.ExcludeTimes(), "got user.")
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
	newRes := makeRequest("GET", "/users/John", nil)

	assert.Equal(t, oldRes.Body, newRes.Body, "got user.")
}
