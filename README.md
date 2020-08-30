# og-service

Golang service for getting OpenGraph data from a URL. Might be useful for link previews.

## Instances

- [og-service.herokuapp.com](https://og-service.herokuapp.com)

## Usage

The URL scheme looks like this:

```
https://https://og-service.herokuapp.com/?link=<link>
```

## Example

```sh
curl 'https://https://og-service.herokuapp.com/?link=https://v1rtl.site'
```

The endpoint will return data of the following structure:

```json
{
  "title": "v 1 r t l ✨",
  "type": "",
  "url": "v1rtl.site",
  "website": "",
  "description": "16 yo nullstack (TS/Go) developer, open sourcerer. Creator of go-web-app, react-postprocessing and tinyhttp. Author of t.me/we_use_js Telegram channel",
  "images": []
}
```

## Installation

The service itself is a go package which you can install and deploy on your own.

```shell
go get -u -v https://github.com/talentlessguy/og-service
```

## Using with Go

And then use it in your http server.

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
