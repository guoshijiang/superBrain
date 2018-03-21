package handle


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/begainPlay", httprouter.Handle(BegainPlay))
	r.GET("/api/endPlay", httprouter.Handle(EndPlay))
	return r
}
