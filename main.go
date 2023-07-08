package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mhdianrush/go-crud-web/config"
	"github.com/mhdianrush/go-crud-web/controllers/categorycontroller"
	"github.com/mhdianrush/go-crud-web/controllers/homecontroller"
	"github.com/mhdianrush/go-crud-web/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	mux := http.NewServeMux()

	// homepage
	mux.HandleFunc("/", homecontroller.Welcome)

	// Categories
	mux.HandleFunc("/categories", categorycontroller.Index)
	mux.HandleFunc("/categories/add", categorycontroller.Add)
	mux.HandleFunc("/categories/edit", categorycontroller.Edit)
	mux.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	mux.HandleFunc("/products", productcontroller.Index)
	mux.HandleFunc("/products/add", productcontroller.Add)
	mux.HandleFunc("/products/detail", productcontroller.Detail)
	mux.HandleFunc("/products/edit", productcontroller.Edit)
	mux.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
