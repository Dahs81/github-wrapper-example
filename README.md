### Github Example Wrapper Project

This is just a simple example wrapper program written in Golang.

Includes:
1. Docker

#### How to Run

```
docker run -it -p 8080:8080 -e GITHUB_TOKEN="<your-github-token>" dahs81/github-wrapper-example
```

OR

```
GITHUB_TOKEN="<your-github-token>" go run main.go
```

Then, in another terminal

```
curl http://localhost:8080/info
```

#### Things I Learned
1. Embedding
1. Github API (Basics)
1. Returning http.HandlerFunc from my function

#### Things To DO
1. Add Prometheus monitoring to app and docker-compose.yml file
1. Testing
1. Some CI build tool (Spinnaker, CircleCI, etc)
