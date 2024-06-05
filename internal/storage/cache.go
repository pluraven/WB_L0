package storage

import (
	"container/list"
	"github.com/go-pg/pg/v10"
	"strconv"
)

type Cache struct {
	Capacity int
	Storage  map[string]string
	List     *list.List
}

func NewCache(capacity int) *Cache {
	return &Cache{
		Capacity: capacity,
		Storage:  make(map[string]string),
		List:     list.New(),
	}
}

func (s *Storage) PutInCache(key string, value string) {
	s.StorageCache.Storage[key] = value
	s.StorageCache.List.PushFront(key)
	if s.StorageCache.List.Len() > s.StorageCache.Capacity {
		delete(s.StorageCache.Storage, s.StorageCache.List.Back().Value.(string))
		s.StorageCache.List.Remove(s.StorageCache.List.Back())
	}
}

func (s *Storage) Fill(db *pg.DB) error {
	var dbSize int
	_, err := db.QueryOne(pg.Scan(&dbSize), "SELECT COUNT(*) FROM json_data")
	if err != nil {
		return err
	}

	var orders []JsonData

	_, err = db.Query(&orders, "SELECT * FROM json_data LIMIT "+strconv.Itoa(s.StorageCache.Capacity))
	if err != nil {
		return err
	}
	for i, _ := range orders {
		s.PutInCache(orders[i].ID, orders[i].DataString)
	}
	return nil
}
