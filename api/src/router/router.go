package router

import "github.com/gorilla/mux"

//GetRouter : returns a router with all routes
func GetRouter() *mux.Router {
	return mux.NewRouter()
}
