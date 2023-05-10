package storage

import "github.com/ankeshnirala/go/structure-go-api/types"

type Storage interface {
	Get(int) *types.User
}
