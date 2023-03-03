package route

import (
	"encoding/json"
	globalhandler "example/server/handler/global_handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Port = "8001"
var URL = "http://127.0.0.1"

func HandlerFunc() {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := globalhandler.ResultRouteError{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := globalhandler.ResultRouteError{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	router.HandleFunc("/", HomeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+Port, router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
