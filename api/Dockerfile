FROM golang:1.11.0
RUN apt-get update


# Add Project Source Code
RUN mkdir -p /go/src/github.com/user/api
COPY . /go/src/github.com/user/api
WORKDIR /go/src/github.com/user/api

CMD ["go", "run",  "/go/src/github.com/user/api"]