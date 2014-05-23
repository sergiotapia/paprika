package paprika

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/bmizerany/pat"
)

var mux = pat.New()

func Attach(route string, resource interface{}) {
	// Grab the resource Type name to use in our Mux.
	resourceName := string(reflect.TypeOf(resource).String())
	resourceName = resourceName[strings.Index(resourceName, ".")+1:]

	mux.Get("/kill", http.HandlerFunc(kill))
	mux.Get("/settings", http.HandlerFunc(settings))
}

func Start(port string) {
	loadDatabaseConfiguration()
	// Hand off our configured Mux to the http server.
	http.Handle("/", mux)
	http.ListenAndServe(":"+port, nil)
}

func kill(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", database_setting)))
}

type ResourceManager interface {
	add() error
	delete() error
	update(id int) error
	show(id int) error
	list() error
}
