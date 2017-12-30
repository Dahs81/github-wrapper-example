### Github Example Wrapper Project

This is just a simple example wrapper program written in Golang.

Includes:
1. Docker

#### How to Run

```
docker run -it -e GITHUB_TOKEN="<your-github-token>" dahs81/github-wrapper-example
```

OR

```
GITHUB_TOKEN="<your-github-token>" go run main.go
```