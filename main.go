package main

import (
	"awesomeProject/src/presenter"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

func handlerRoot(
	writer http.ResponseWriter,
	request *http.Request,
) {
	urlPathElements := strings.Split(request.URL.Path, "/")
	endpoint := urlPathElements[1]

	switch endpoint {
	case "top-films":
		resource := urlPathElements[2]
		presenter.GetTopFilms(resource, writer)
	default:
		writer.WriteHeader(http.StatusBadRequest)
		response := []byte("400 - Bad request, verify your request")
		_, err := writer.Write(response)
		if err != nil {
			return
		}
	}
}

func main() {
	// serverHttp()
	serverHttpRouter()
}

func serverHttpRouter() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", presenter.GetGoVersion)
	router.GET("/api/v1/variable/:variable", presenter.GetVariable)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func serverHttp() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/", handlerRoot)

	err := http.ListenAndServe(":8000", newMux)

	if err != nil {
		fmt.Print("Server not online")
		return
	}
}
