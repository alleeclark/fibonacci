package routers

import (
	"fibonacci/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetFibRoutes registers routes for fibonacci sequence.
func SetFibRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/fib", controllers.PostFibSequenceNums).Methods("POST")
	http.Handle("/", router)
	return router
}
