FROM golang:1.9
MAINTAINER Shad Beard

RUN go get github.com/gorilla/mux
RUN go get golang.org/x/oauth2
RUN go get github.com/google/go-github/github

ADD . /app/
WORKDIR /app
RUN go build -o github .

CMD ["/app/github"]