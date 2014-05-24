Sprinkle some paprika over your models and instantly have it 
served via REST API.

## Getting Started

Paprika is really simple to use. All you need to do is attach
your structs to it and it'll handle the rest.

``` go
package main

import (
    "fmt"

    "github.com/sergiotapia/paprika"
)

type Users struct {
    Email             string `email`
    EncryptedPassword string `encrypted_password`
    CreatedAt         string `created_at`
    PermissionLevel   string `permission_level`
}

func main() {
    // Have Paprika serve your struct model by attaching
    // a route to respond to and your struct Type.
    paprika.Attach("/users", Users{})

    // Start the Paprika service using a port.
    paprika.Start("9090")

    fmt.Println("Paprika is up and running.")
}

```

## How does it work?

Paprika uses reflection to build out a map of your database using
the structs you attach to it. Using this map, Paprika generates typical
API endpoints and queries.

## Work in Progress

Paprika is still not complete!

The first release will only support PostgreSQL since I don't have much
experience with MySQL or any of the NoSQL databases.

Feature List:

- [x] Build out database schema from user defined structs.
- [x] Load database.toml database settings.
- [ ] User database settings from .toml file to initialize database adapter.
- [ ] List
  - [ ] Show all records in JSON format.
  - [ ] Pagination
  - [ ] Order by field.
- [ ] Detail
  - [ ] Show single record by ID in JSON format.
- [ ] Delete
  - [ ] Delete single record by ID and return result in JSON format.
- [ ] Create
  - [ ] Receive JSON values in POST and create new record and return result in JSON format.
- [ ] Update
  - [ ] Receive JSON values in POST and update existing record and return result in JSON format. 

  