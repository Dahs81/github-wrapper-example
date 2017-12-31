package main

// Simple Server with /info endpoint that returns all public github info

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	gh "github.com/Dahs81/github-wrapper-example/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := gh.NewMyGithubClient(tc)

	// Example of Embedding/Composition (LEARNING)
	// c.Client.Apps AND c.Apps are the same thing

	r := mux.NewRouter()
	r.Handle("/info", gh.Info(ctx, &client))

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    8 * time.Second,
		WriteTimeout:   8 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}
	log.Fatal(s.ListenAndServe())

}
Hello world....
