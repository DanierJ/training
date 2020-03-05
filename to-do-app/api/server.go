package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/danierj/training/to-do-app/api/routes"
	"github.com/gorilla/mux"
)

func InitServer(port string) error {
	fmt.Println("\n-------------------- Starting server at port " + port + " --------------------")

	r := routes.NewRouter()
	return validServerPort(port, r)
}

func validServerPort(port string, r *mux.Router) error {
	re := regexp.MustCompile(`^\d{4}$`)
	if port == "" || !re.Match([]byte(port)) {
		port = "6745"
		log.Println("\n\n########################### INVALID PORT NUMBER ######################")
		log.Println("\n\nInvalid port number. Stablishing port to: " + port)
	}

	if err := start(port, r); err != nil {
		log.Println("\n\n########################### PORT ALREADY IN USE ######################")
		err = errors.New("This port is already in use. Please provide another one")
		log.Fatal(err)
		return err
	}
	return nil
}

func start(port string, r *mux.Router) error {
	if err := http.ListenAndServe(":"+port, r); err != nil {
		return err
	}

	return nil
}
