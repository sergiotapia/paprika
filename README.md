Sprinkle some paprika over your models and instantly have it 
served via REST API.

## Getting Started

Paprika is really simple to use. All you need to do is attach
your structs to it and it'll handle the rest.

``` go
package main

import (
    "fmt"
    "log"

    "github.com/sergiotapia/paprika"
)

type Product struct {
    Name     string
    Quantity int
}

func main() {
    paprika.Attach(new(Product))
    paprika.Start("8080")
    log.Print("Paprika is up and running.")
}
```

## How does it work?

Paprika uses reflection to determine what the attached resource Type
and attaches the appropriate REST URL's to the http server.

## Work in Progress

Paprika is still not complete! I've bitten off a bit more than I can chew
but hopefully I can circle back soon and complete it. I think it'll be 
really helpful for getting an API up and running quickly using Go.

Particularly I'm having trouble figuring out how to store many different
structs since Paprika cannot access structs defined in your calling go
program.

Ideally you will be able to set up your API in 5 minutes flat.