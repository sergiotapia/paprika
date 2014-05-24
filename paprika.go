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

// This holds the schema of the database tables.
var schema = make(map[string]map[string]string)

func Attach(route string, resource interface{}) {
	loadTypeToSchema(resource)

	mux.Get("/kill", http.HandlerFunc(kill))
	mux.Get("/settings", http.HandlerFunc(settings))

	mux.Get(route, http.HandlerFunc(list(resource)))
}

func Start(port string) {
	loadDatabaseConfiguration()
	// Hand off our configured Mux to the http server.
	http.Handle("/", mux)
	http.ListenAndServe(":"+port, nil)
}

func list(resource interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		db := initDb()
		query := createListQuery(resource)
		rows, err := db.Query(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)

			for rows.Next() {

				// Next up, for every row returned we need to
				// get the values and build it out in the JSON
				// // response.

				// Example from the internet:
				// var name string
				// if err := rows.Scan(&name); err != nil {
				// 	log.Fatal(err)
				// }
				// fmt.Printf("%s is %d\n", name, age)
			}

			w.Write([]byte(fmt.Sprintf("%v", rows)))
		}
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", schema)))
}

func kill(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

// Extracts resource information using reflection and
// saves field names and types.
func loadTypeToSchema(resource interface{}) {
	resourceType := reflect.TypeOf(resource)

	// Grab the resource struct Name.
	fullName := resourceType.String()
	name := strings.ToLower(fullName[strings.Index(fullName, ".")+1:])

	// Grabs the resource struct fields and types.
	fields := make(map[string]string)
	for i := 0; i <= resourceType.NumField()-1; i++ {
		field := resourceType.Field(i)
		fields[string(field.Tag)] = field.Type.Name()
	}

	// Add resource information to schema map.
	schema[name] = fields
}
