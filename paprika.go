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

	// This just kills the server.
	mux.Get("/kill", http.HandlerFunc(kill))

	// Apply typical REST actions to Mux.
	// ie: Product - to /products
	mux.Get("/"+resourceName+"s", http.HandlerFunc(index))

	// ie: Product ID: 1 - to /products/1
	mux.Get("/"+resourceName+"/:id", http.HandlerFunc(show))
}

func Start(port string) {
	// Hand off our configured Mux to the http server.
	http.Handle("/", mux)
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	resourceType := r.URL.String()[1:]
	x := Product{"Abc", 20}
	w.Write([]byte(fmt.Sprintf("%v", params)))
}

func show(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get(":id")
	w.Write([]byte("Hello " + id))
}

func kill(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

type ResourceManager interface {
	add() error
	delete() error
	update(id int) error
	show(id int) error
	list() error
}
