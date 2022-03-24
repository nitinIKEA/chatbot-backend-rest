package main

import (
	"log"
	"net/http"

	"github.com/nitinIKEA/chatbot-backend-rest/internal/config"
	"github.com/nitinIKEA/chatbot-backend-rest/internal/db"
	"github.com/nitinIKEA/chatbot-backend-rest/internal/service"
)

func main() {
	//prepare service
	srv := service.Service{}
	//read config file
	var err error
	srv.Conf, err = config.New("./configs/configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	//create database connections
	srv.DBConns = db.GetConnections(srv.Conf)
	//server
	srv.NewRouter()
	log.Printf("starting application on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
