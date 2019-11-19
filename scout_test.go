package scout_test

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zy4/scout"
)

func ExampleSentryRecovery() {
	r := chi.NewRouter()

	// Apply the middleware to the router
	r.Use(scout.SentryRecovery)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("catched")
	})
}
