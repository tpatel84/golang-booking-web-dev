package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/tpatel84/golang-booking-web-dev/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app *config.AppConfig

	mux := routes(app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing, test pass
	default:
		t.Error(fmt.Sprintf("type isn ot *chi.mux, type is %T", v))
	}
}
