package tests

import (
	"testing"

	"github.com/danierj/training/to-do-app/api/routes"
)

func TestNewRouter(t *testing.T) {
	r := routes.NewRouter()

	if r == nil {
		t.Errorf("Expected to have an *mux.Router instance. Instead got nil")
	}
}
