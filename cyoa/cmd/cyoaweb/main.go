package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/luciormoraes/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3001, "the port to start the CYOA web application on")
	filename := flag.String("filename", "gopher.json", "the JSON filename with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", story)
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", *port), h))
}
