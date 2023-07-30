package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/mhdianrush/go-crud-web/config"
	"github.com/mhdianrush/go-crud-web/controllers/categorycontroller"
	"github.com/mhdianrush/go-crud-web/controllers/productcontroller"
)

var logger = logrus.New()

func main() {
	config.ConnectDB()

	routes := mux.NewRouter()

	routes.HandleFunc("/categories", categorycontroller.Index)
	routes.HandleFunc("/categories/add", categorycontroller.Add)
	routes.HandleFunc("/categories/edit", categorycontroller.Edit)
	routes.HandleFunc("/categories/delete", categorycontroller.Delete)

	routes.HandleFunc("/products", productcontroller.Index)
	routes.HandleFunc("/products/add", productcontroller.Add)
	routes.HandleFunc("/products/detail", productcontroller.Detail)
	routes.HandleFunc("/products/edit", productcontroller.Edit)
	routes.HandleFunc("/products/delete", productcontroller.Delete)

	err := godotenv.Load()
	if err != nil {
		logger.Printf("failed loaded env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: routes,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}

	logger.Printf("Server Running on Port %s", os.Getenv("SERVER_PORT"))
}
