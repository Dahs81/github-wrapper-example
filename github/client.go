package github

// MyGithubClient - Example of using embedding (see lines 26-28 as well)
import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/prometheus/common/log"
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
			log.Error("Unable to list the repositories: ", err)
		}

		// c.Client.Apps AND c.Apps are the same

		out, e := json.Marshal(repos)
		if e != nil {
			log.Error("Unable to marshal repos to json: ", err)
		}

		w.Write(out)
	}
}
