package presenter

import (
	"awesomeProject/src/domain"
	"net/http"
	"strconv"
	"strings"
)

func GetTopFilms(
	resource string,
	writer http.ResponseWriter,
) {
	resourceNoneSpace := strings.TrimSpace(resource)
	position, _ := strconv.Atoi(resourceNoneSpace)

	if data, err := domain.SearchTopFilm(position); err == nil {
		writer.WriteHeader(http.StatusOK)
		response := []byte(data)
		_, err := writer.Write(response)
		if err != nil {
			return
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		response := []byte("Resource not found")
		_, err := writer.Write(response)

		if err != nil {
			return
		}
	}
}
