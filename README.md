# Kenobi
A simple golang web template that gives you a high ground

### Libraries Used
* Command Line Interface (CLI): Cobra
* ORM: gorm
* Routing: gorilla/mux, gorilla/mux, gorilla/handlers


## Getting Started

create a file `main.go` in your project folder


```go

package main

import (
	"github.com/gorilla/mux"
	"github.com/theNullP0inter/kenobi/cmd"
	"github.com/theNullP0inter/kenobi/config"
	"net/http"
)

func main() {

	myConf := config.Config{
		"SecretKey":       "123",
		"DatabaseUri":     "host=localhost port=5432 user=admin dbname=test password=adminpassword sslmode=disable",
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



``` 

## Config

