package server

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Alireza-Ta/GOASK/model"
)

func TestPostLogin(t *testing.T) {
	user := &model.User{
		Username: "admin",
		Password: "admin",
	}
	body, err := json.Marshal(user)
	checkNil(err, "error in json marshal.")
	res := makeRequest("POST", "/login", bytes.NewBuffer(body), nil)

	var b map[string]interface{}
	err = json.Unmarshal(res.Body.Bytes(), &b)
	checkNil(err, "json unmarshal.")

	// headers := map[string]string{
	// 	"Authorization": "Bearer " + b["token"].(string),
	// }

	// res2 := makeRequest("GET", "/auth/hello", nil, headers)

}
