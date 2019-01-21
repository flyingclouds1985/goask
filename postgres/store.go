package postgres

import (
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var models = []interface{}{
	&model.User{},
	&model.Comment{},
	&model.Question{},
	&model.Reply{},
	&model.Tag{},
}

// Config type.Default value for password is empty string.
type Config struct {
	Username string
	Password string
	DBname   string
}

// Store represents postgres store instance.
type Store struct {
	Config Config
	DB     *pg.DB
}

// New makes a new psotgres instance.
func New(config ...Config) *Store {
	conf := initPostgresConfig(config)
	return &Store{
		DB: openDB(conf.Username, conf.Password, conf.DBname),
	}
}

func initPostgresConfig(config []Config) Config {
	defaultConfig := Config{
		Username: "postgres",
		Password: "",
		DBname:   "GOASK",
	}
	switch len(config) {
	case 0:
		return defaultConfig
	case 1:
		conf := config[0]
		if conf.Username == "" {
			conf.Username = defaultConfig.Username
		}
		if conf.DBname == "" {
			conf.DBname = defaultConfig.DBname
		}
		return conf
	default:
		panic("too much argument!")
	}
}

func openDB(username, password, dbname string) *pg.DB {
	return pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: dbname,
	})
}

// CreateSchema create tables.
func (s *Store) CreateSchema() error {
	for _, model := range models {
		err := s.DB.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
