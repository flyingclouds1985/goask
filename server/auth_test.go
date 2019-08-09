package server

import (
	"bytes"
	"encoding/json"
	"github.com/Alireza-Ta/goask/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostLogin(t *testing.T) {
	SetupSubTest()
	defer TeardownSubTest()

	user := &model.User{
		Username:        "admin",
		Password:        "12345678",
		ConfirmPassword: "12345678",
	}

	// create user
	body, err := json.Marshal(user)
	checkNil(err, "user: err in json parsing.")

	res := testMakeRequest("POST", "/users/", bytes.NewBuffer([]byte(body)), nil)

	var u model.User
	err = json.Unmarshal(res.Body.Bytes(), &u)
	checkNil(err, " user: err in json unmarshal.")
	assert.Equal(t, u.Username, user.Username, "got username.")
	// end create user

	body, err = json.Marshal(user)

	checkNil(err, "error in json marshal.")
	res = testMakeRequest("POST", "/login", bytes.NewBuffer(body), nil)

	var b map[string]interface{}

	err = json.Unmarshal(res.Body.Bytes(), &b)
	checkNil(err, "json unmarshal.")

	headers := map[string]string{
		"Authorization": "Bearer " + b["token"].(string),
	}
	res2 := testMakeRequest("GET", "/auth/hello", nil, headers)
	assert.Equal(t, headers["Authorization"], res2.Header().Get("Authorization"), "got Authorization header.")
}
