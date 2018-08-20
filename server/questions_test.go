package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/stretchr/testify/assert"
)

func setup() *Server {
	store := postgres.New()
	err := store.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	server := &Server{}
	server.Store = store
	server.SetupRoute()

	return server
}

func TestGetQuestion(t *testing.T) {
	s := setup()
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
			s.Router.ServeHTTP(res, req)
			if k == "successful" {
				assert.Equal(t, 200, res.Code, "got question.")

			} else {
				assert.Equal(t, 400, res.Code, "didnt get question.")
			}
		})
	}

}
