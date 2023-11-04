package main

import (
	"awesomeProject/src/presenter"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
	"time"
)

const RUNNER_SERVER = 2

func main() {
	switch RUNNER_SERVER {
	case 0:
		serverHttp()
	case 1:
		serverHttpRouter()
	case 2:
		serverMux()
	default:
		log.Fatal("Option not exist, please use [0,1 or 2]")
	}
	// serverHttp()
	// serverHttpRouter()
}

func serverMux() {
	router := mux.NewRouter()
	router.UseEncodedPath()
	custom := router.PathPrefix("/api").Subrouter()
	custom.HandleFunc("/articles/{category}/{id}", presenter.GetArticle)
	//custom.Queries("name")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
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
