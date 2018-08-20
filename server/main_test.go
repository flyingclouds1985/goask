package server

import (
	"log"
	"os"
	"testing"

	"github.com/Alireza-Ta/GOASK/postgres"
)

var TestServer *Server

func setup() {
	store := postgres.New()
	err := store.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	TestServer = &Server{}
	TestServer.Store = store
	TestServer.SetupRoute()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
