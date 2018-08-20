package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuestion(t *testing.T) {

	tests := map[string]struct {
		id string
	}{
		"successful": {
			id: "1",
		},
		"unsuccessful": {
			id: "100",
		},
	}

	for k, question := range tests {
		t.Run(k, func(t *testing.T) {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/questions/"+question.id, nil)
			if err != nil {
				t.Error("error in getting queestion from endpoint!!!")
			}
			TestServer.Router.ServeHTTP(res, req)
			if k == "successful" {
				assert.Equal(t, 200, res.Code, "got question.")

			} else {
				assert.Equal(t, 404, res.Code, "didnt get question.")
			}
		})
	}

}
