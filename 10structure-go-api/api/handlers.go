package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankeshnirala/go/structure-go-api/util"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	val := util.Round2Dec(10.233444)

	fmt.Println("Rounded value ", val)

	json.NewEncoder(w).Encode(user)
}
