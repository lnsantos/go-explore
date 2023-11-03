package presenter

import (
	"awesomeProject/src/domain"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func GetGoVersion(
	res http.ResponseWriter,
	req *http.Request,
	p httprouter.Params,
) {
	if _, err := fmt.Fprint(
		res,
		domain.GetCommandOutput("go", "version"),
	); err != nil {
		res.WriteHeader(http.StatusBadGateway)
		_, _ = res.Write([]byte("problem with service"))
		log.Fatal(err)
		//return
	}
}
