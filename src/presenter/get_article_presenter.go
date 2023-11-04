package presenter

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetArticle(
	res http.ResponseWriter,
	req *http.Request,
) {
	params := mux.Vars(req)
	querys := req.URL.Query()

	category := params["category"]
	id := params["id"]
	qName := querys["name"]

	res.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprintf(res, "Name is %v\n", qName); err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fprintf(res, "Category is %v\n", category); err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fprintf(res, "Id is %v", id); err != nil {
		log.Fatal(err)
	}
}
