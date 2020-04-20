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

	cmd.Execute(myConf, []cmd.ExecCommandFunction{testCommand, testCommand2})

}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", testHandler)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("yoo"))
}

func testCommand(app *app.App) *cobra.Command {

	return &cobra.Command{
		Use:   "custom2",
		Short: "Print the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app)
		},
	}
}

func testCommand2(app *app.App) *cobra.Command {

	return &cobra.Command{
		Use:   "custom",
		Short: "Print the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app)
		},
	}
}
