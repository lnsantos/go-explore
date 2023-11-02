package main

import (
	"awesomeProject/src/presenter"
	"fmt"
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
	http.HandleFunc("/", handlerRoot)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Print("Server not online")
		return
	}
}
