package route

import (
	"encoding/json"
	bookhandler "example/server/handler/book_handler"
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

	// buku
	router.HandleFunc("/api/buku", bookhandler.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/buku", bookhandler.Store).Methods("POST")
	router.HandleFunc("/api/buku/detail/{book_id}", bookhandler.Detail).Methods("GET")
	router.HandleFunc("/api/buku/update/{book_id}", bookhandler.Update).Methods("PUT")
	router.HandleFunc("/api/buku/delete/{book_id}", bookhandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+Port, router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}
