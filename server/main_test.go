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
	"path/filepath"
	"github.com/Alireza-Ta/goask/postgres"
	"github.com/Alireza-Ta/goask/pkg/config"
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
	storeConf := postgres.Config{Password: "secret", DBname: "GoaskTest"}
	store := postgres.New(storeConf)

	currentDir, err := os.Getwd()
	root := filepath.Dir(currentDir)

	c, err := config.Load(root + "/configuration.json")
	if err != nil {
		panic(err)
	}

	TestServer = NewServer(store, gin.TestMode, c)
}

func truncateAllTables() {
	for _, model := range models {
		_, err := TestServer.Store.DB.Model(model).Exec(
			fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", model),
			// drop all tables
			//fmt.Sprintf("DROP SCHEMA public CASCADE; CREATE SCHEMA public;"),
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
	TestServer.Router.ServeHTTP(res, req)

	return res
}
