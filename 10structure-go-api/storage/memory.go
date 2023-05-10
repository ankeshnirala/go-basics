package storage

import "github.com/ankeshnirala/go/structure-go-api/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(int) *types.User {
	return &types.User{
		ID:   123,
		Name: "Ankesh Kumar",
	}
}
