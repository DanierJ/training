package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/danierj/training/to-do-app/api/routes"
)

func initServer(port int) {
	fmt.Println("\n-------------------- Starting server at port " + strconv.Itoa(port) + " --------------------")

	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
