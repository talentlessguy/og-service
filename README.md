# og-service

Golang service for getting OpenGraph data from a URL.

[Open ↗](https://og-service.herokuapp.com/)

## Usage

The URL scheme looks like this:

```
https://https://og-service.herokuapp.com/?link=<link>
```

You can use this both on frontend and backend.

## Installation

The service itself is a go package which you can install and deploy on your own.

```shell
go get -u -v https://github.com/talentlessguy/og-service
```

And then use it in your http server:

```go
package main

import (
	"log"
	"net/http"

	preview "github.com/talentlessguy/og-service"
)

func main() {

  origins := []string{"*"}

	s := preview.NewReactLinkPreviewService(origins)

	log.Println("Started a service on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", s))
}
```
