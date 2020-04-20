# Kenobi
A simple golang web template that gives you a high ground

### Libraries Used
* Command Line Interface (CLI): Cobra
* ORM: gorm
* Routing: gorilla/mux, gorilla/mux, gorilla/handlers


## Getting Started

Open your project folder

Initialize your module using ```go mod init MOD_NAME```

Create a file `main.go`. This file will pass all your application configurations to the CLI.

Add the main function


```go

package main

func main(){
	
	
}

``` 

Initialize your app config with ```SecretKey, DatabaseUri, DatabaseDialect,```

More details about the config are in the Config section below

```go

//...
import "github.com/theNullP0inter/kenobi/config"


// func main

myConf := config.Config{
        "SecretKey": "123",
        "DatabaseUri": "/tmp/gorm.db",
        "DatabaseDialect": "sqlite3",
    }
``` 



Pass the config to your Application using 
``` go

import "github.com/theNullP0inter/kenobi/cmd"


// func main
//    ...

cmd.Execute(myConf, cmd.CommandCenter{}) 
/*
    cmd.CommandCenter{} is a list of custom commands you can set. 

    Explained more in detail in ComandCenter section   

*/


``` 


Now, Add a function to handle your routes.

This function receives ```*mux.Router``` as the argument. use this argument to handle your Routes

```go


// func main(){...}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello There"))
}

```


Add the above handler to the App config as `RouterHandler`



```go

// func main

myConf := config.Config{
        "SecretKey": "123",
        "DatabaseUri": "/tmp/gorm.db",
        "DatabaseDialect": "sqlite3",
        "RouterHandler": RegisterRoutes
    }
    
  //...
  
  
  
  //funcRegisterRoutes(){...}
  

```

    
Final main file would look like this:

```go

// main.go

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/theNullP0inter/kenobi/app"
	"github.com/theNullP0inter/kenobi/cmd"
	"github.com/theNullP0inter/kenobi/config"
	"net/http"
)

func main() {

	myConf := config.Config{
		"SecretKey":       "123",
		"DatabaseUri":     "/tmp/gorm.db",
		"DatabaseDialect": "sqlite3",
		"RouterHandler":   RegisterRoutes,
	}

	cmd.Execute(myConf, cmd.CommandCenter{})

}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello There"))
}



``` 

Start the server using the following command.

``` go run main.go serve ```



## Config

Config is map that contains all the information needed for your Application






