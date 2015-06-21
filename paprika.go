package paprika

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
)

var mux = pat.New()

// This holds the schema of the database tables.
var schema = make(map[string]map[string]string)

func Attach(route string, resource interface{}) {
	mux.Get("/kill", http.HandlerFunc(kill))
	mux.Get("/settings", http.HandlerFunc(settings))
}

func Start(port string) {
	// Hand off our configured Mux to the http server.
	http.Handle("/", mux)
	http.ListenAndServe(":"+port, nil)
}

func settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", schema)))
}

func kill(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
