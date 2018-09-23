package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/L-oris/yabb/controller/postcontroller"
	"github.com/L-oris/yabb/controller/rootcontroller"
	"github.com/L-oris/yabb/httperror"
	"github.com/L-oris/yabb/inject"
	"github.com/L-oris/yabb/inject/types"
	"github.com/L-oris/yabb/models/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/post").Handler(
		negroni.New(negroni.Wrap(
			inject.Container.Get(types.PostController.String()).(postcontroller.Controller).Router),
		))

	router.PathPrefix("/").Handler(
		negroni.New(negroni.Wrap(
			inject.Container.Get(types.RootController.String()).(rootcontroller.Controller).Router),
		))

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		httperror.NotFound(w, "Route Not Found")
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	server := &http.Server{
		Addr:         ":" + env.Vars.Port,
		Handler:      loggedRouter,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(server.ListenAndServe())
}
