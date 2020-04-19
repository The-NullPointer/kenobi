package router

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"kenobi/app"
	"net/http"
)

type Router struct {
	App    *app.App
	Router *mux.Router
}

func New(a *app.App) *Router {
	router := mux.NewRouter()
	return &Router{a, router}
}

func (r *Router) Init() {
	appRouter := r.Router
	appRouter.Use(r.addConfigToContext)
}

func (r *Router) addConfigToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		context.Set(req, "db", r.App.Database)
		context.Set(req, "config", r.App.Config)
		next.ServeHTTP(w, req)
	})
}
