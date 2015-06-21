package paprika

import (
  "reflect"
  "strings"
)

func createListQuery(resource interface{}) string {
  resourceType := reflect.TypeOf(resource)

  // Grab the resource struct Name.
  fullName := resourceType.String()
  name := strings.ToLower(fullName[strings.Index(fullName, ".")+1:])

  query := "SELECT "

  for fieldName, _ := range schema[name] {
    query += fieldName + ", "
  }

  query = query[:len(query)-2] + " "
  query += "FROM "
  query += name

  return query
}
