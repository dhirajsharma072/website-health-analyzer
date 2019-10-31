package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dhirajsharma072/website-health-analyzer/src/controller"
	. "github.com/dhirajsharma072/website-health-analyzer/src/dao"
	"github.com/gorilla/mux"
)

var dao = WebsiteDAO{}
var defaultPort = "9000"

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	dao.HostName = "localhost"
	dao.DatabaseName = "website-analyzer"
	dao.Connect()
}

// Define HTTP request routes
func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := mux.NewRouter()
	r.HandleFunc("/websites", controller.AllWebsites).Methods("GET")
	r.HandleFunc("/websites", controller.CreateWebsite).Methods("POST")
	r.HandleFunc("/websites", controller.UpdateWebsite).Methods("PUT")
	r.HandleFunc("/websites/{id}", controller.DeleteWebsite).Methods("DELETE")
	r.HandleFunc("/websites/{id}", controller.PatchWebsite).Methods("PATCH")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
