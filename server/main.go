package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"incidentreport/app/index"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	port := 9090

	router := mux.NewRouter()
	router.HandleFunc("/", index.Index)

	// cors
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-EventType"})
	methods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
	})


	log.Println("starting http server on ", port)
	log.Println("starting http server on ", port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
		Handler: handlers.CORS(origins, headers, methods)(router),
	}

	log.Fatal(server.ListenAndServe())
}
