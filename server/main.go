package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"incidentreport/app/incident"
	"incidentreport/app/index"
	"incidentreport/db/database"
	"incidentreport/pkg/middleware"
	"incidentreport/pkg/route"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	port := 9090

	var routes = route.Routes{
		route.Route{
			Name:            "Index",
			Method:          http.MethodGet,
			Pattern:         "/",
			HandlerFunction: index.Index,
		},
		route.Route{
			Name:            "CreateIncident",
			Method:          http.MethodPost,
			Pattern:         "/incidents",
			HandlerFunction: incident.HandleImageUpload,
		},
		route.Route{
			Name:            "AddIncidentMetaData",
			Method:          http.MethodPost,
			Pattern:         "/incidents/{id}",
			HandlerFunction: incident.HandleAddMetaData,
		},
		route.Route{
			Name:            "GetOneIncident",
			Method:          http.MethodGet,
			Pattern:         "/incidents/{id}",
			HandlerFunction: incident.HandleGetOneIncident,
		},
	}

	router := mux.NewRouter()
	for _, oneRoute := range routes {
		var handler http.Handler

		handler = oneRoute.HandlerFunction

		router.
			Methods(oneRoute.Method).
			Path(oneRoute.Pattern).
			Name(oneRoute.Name).
			Handler(handler)
	}

	// database connection
	err := database.Connect()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database connection established")

	defer func() {
		err := database.Disconnect()
		if err != nil {
			log.Println(err)
		}
		log.Println("Database disconnected")
	}()

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
		Addr:    fmt.Sprintf(":%v", port),
		Handler: handlers.CORS(origins, headers, methods)(router),
	}

	router.Use(middleware.JSONMiddleware)

	log.Fatal(server.ListenAndServe())
}
