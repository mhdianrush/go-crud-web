package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mhdianrush/go-crud-web/config"
	"github.com/mhdianrush/go-crud-web/controllers/categorycontroller"
	"github.com/mhdianrush/go-crud-web/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
