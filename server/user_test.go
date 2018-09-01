package server

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/stretchr/testify/assert"
)

var userTestCases = map[string]model.User{
	"complete": model.User{
		Username: "John",
		Email:    "john@example.com",
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

			assert.Equal(t, 200, res.Code, "user created.")
		})
	}
}
