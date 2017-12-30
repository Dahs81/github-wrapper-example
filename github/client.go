package github

// MyGithubClient - Example of using embedding (see lines 26-28 as well)
import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
)

// MyGithubClient -
type MyGithubClient struct {
	*github.Client
}

// NewMyGithubClient - return a MyGithubClient
func NewMyGithubClient(o *http.Client) MyGithubClient {
	c := MyGithubClient{}
	c.Client = github.NewClient(o)

	return c
}

// Info - endpoint for all github info
func Info(ctx context.Context, c *MyGithubClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repos, _, err := c.Repositories.List(ctx, "", nil)
		if err != nil {
			panic(err)
		}

		out, e := json.Marshal(repos)
		if e != nil {
			panic(err)
		}

		w.Write(out)
	}
}
