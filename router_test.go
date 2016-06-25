package router

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func handler(http.ResponseWriter, *http.Request, httprouter.Params) {

}

func TestAverage(t *testing.T) {
	router := New()
	router.GET("/data/:id", handler).Name("dataID")
	path := router.MustReverse("dataID", "2")
	if path != "/data/2" {
		t.Errorf("the path is %s when should be /data/2", path)
		t.Fail()
	}
}
