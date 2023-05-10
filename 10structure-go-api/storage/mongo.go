package storage

import "github.com/ankeshnirala/go/structure-go-api/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(int) *types.User {
	return &types.User{
		ID:   123,
		Name: "Ankesh Kumar",
	}
}
