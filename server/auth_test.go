package server

import (
	"bytes"
	"encoding/json"
	"testing"
	"github.com/Alireza-Ta/goask/model"
	"github.com/stretchr/testify/assert"
)

func TestPostLogin(t *testing.T) {
	user := &model.User{
		Username: "admin",
		Password: "admin",
	}
	body, err := json.Marshal(user)

	checkNil(err, "error in json marshal.")
	res := testMakeRequest("POST", "/login", bytes.NewBuffer(body), nil)

	var b map[string]interface{}

	err = json.Unmarshal(res.Body.Bytes(), &b)
	checkNil(err, "json unmarshal.")

	headers := map[string]string{
		"Authorization": "Bearer " + b["token"].(string),
	}
	res2 := testMakeRequest("GET", "/auth/hello", nil, headers)
	assert.Equal(t, headers["Authorization"], res2.Header().Get("Authorization"), "got Authorization header.")
}