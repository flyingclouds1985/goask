package server

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Alireza-Ta/GOASK/postgres"
)

var TestServer *Server
var models []string = []string{
	"users",
	"comments",
	"questions",
	"replies",
	"tags",
}

func setup() {
	store := postgres.New("postgres", "secret", "GoaskTest")

	TestServer = &Server{}
	TestServer.Store = store
	TestServer.SetupRoute()
}

func truncateAllTables() {
	for _, model := range models {
		_, err := TestServer.Store.DB.Model(model).Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", model))
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
