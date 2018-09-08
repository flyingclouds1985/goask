package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
)

var (
	TestServer *Server
	models     = []string{
		"users",
		"comments",
		"questions",
		"replies",
		"tags",
	}
)

func setup() {
	store := postgres.New("postgres", "secret", "GoaskTest")

	TestServer = &Server{}
	TestServer.Store = store
	TestServer.SetupRoute(gin.TestMode)
}

func truncateAllTables() {
	for _, model := range models {
		_, err := TestServer.Store.DB.Model(model).Exec(
			fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", model),
		)
		if err != nil {
			log.Fatal("Error in truncating...", err)
		}
	}
}

func SetupSubTest() {
	err := TestServer.Store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}
}

func TeardownSubTest() {
	truncateAllTables()
}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func makeRequest(method string, url string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	checkNil(err, "error in makeing request.")
	for k, v := range headers {
		fmt.Println("-------------", k, v)
		req.Header.Set(k, v)
	}
	TestServer.Router.ServeHTTP(res, req)

	return res
}
