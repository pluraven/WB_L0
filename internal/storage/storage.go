package storage

import (
	"github.com/go-pg/pg/v10"
)

type Storage struct {
	StorageDB    *pg.DB
	StorageCache *Cache
}

func Load(
	usernameDB string,
	passwordDB string,
	addressDB string,
	databaseDB string,
	cacheCapacity int,
) (*Storage, error) {
	cache := NewCache(cacheCapacity)
	db, err := NewDB(
		usernameDB,
		passwordDB,
		addressDB,
		databaseDB,
	)
	if err != nil {
		return nil, err
	}
	s := &Storage{db, cache}
	err = s.Fill(db)
	if err != nil {
		return nil, err
	}
	return s, nil
}
