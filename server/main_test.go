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
	AppServer *Server
	models    = []string{
		"users",
		"comments",
		"questions",
		"replies",
		"tags",
	}
)

func setup() {
	storeConf := &postgres.Config{Password: "secret", DBname: "GoaskTest"}
	store := postgres.New(storeConf)

	AppServer = NewServer(store, gin.TestMode)
}

func truncateAllTables() {
	for _, model := range models {
		_, err := AppServer.Store.DB.Model(model).Exec(
			fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", model),
		)
		if err != nil {
			log.Fatal("Error in truncating...", err)
		}
	}
}

func SetupSubTest() {
	err := AppServer.Store.CreateSchema()
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

func testMakeRequest(
	method, url string,
	body io.Reader,
	headers map[string]string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	checkNil(err, "error in makeing request.")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	AppServer.Router.ServeHTTP(res, req)

	return res
}
