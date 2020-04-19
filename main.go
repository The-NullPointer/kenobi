package main

import (
	"github.com/gorilla/mux"
	"kenobi/cmd"
	"kenobi/config"
	"net/http"
)

func main() {

	myConf := config.Config{
		"SecretKey":       "123",
		"DatabaseUri":     "host=localhost port=5432 user=admin dbname=zipi password=adminpassword sslmode=disable",
		"DatabaseDialect": "postgres",
		"RouterHandler":   RegisterRoutes,
	}

	cmd.Execute(myConf)

}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", testHandler)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("yoo"))
}
