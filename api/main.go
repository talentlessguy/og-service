package main

import (
	"log"
	"net/http"

	"os"

	service "github.com/talentlessguy/og-service"
)

func main() {

	port := os.Getenv("PORT")

	// Those are RelayChat origins, replace them with your own
	origins := []string{"*"}

	s := service.NewReactLinkPreview(origins)

	log.Println("Started a service on http://localhost:" + port)

	log.Fatal(http.ListenAndServe(":"+port, s))
}
