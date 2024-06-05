package storage

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

type JsonData struct {
	ID         string
	Data       string
	DataString string
}

func NewDB(user string, password string, address string, database string) (*pg.DB, error) {
	log.Println("Connecting to database...")
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Addr:     address,
		Database: database,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Storage) WriteToDB(model *JsonData) error {
	_, err := s.StorageDB.Model(model).Insert()
	if err != nil {
		return err
	}
	return nil
}
