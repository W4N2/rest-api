package api

import (
	"testing"

	"github.com/w4n2/rest-api/src/store"
)

func TestNewAPIServer(t *testing.T) {
	addr := "localhost:8080"
	store := &store.MockStore{}

	server := NewAPIServer(addr, store)

	if server.addr != addr {
		t.Errorf("Expected address %s, but got %s", addr, server.addr)
	}

	if server.store != store {
		t.Errorf("Expected store %v, but got %v", store, server.store)
	}
}
