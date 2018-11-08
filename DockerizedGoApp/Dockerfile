FROM golang:latest
MAINTAINER Eky Pradhana eky.pradhana@eaciit.com
ADD . /go/src/DockerizedGoApp
WORKDIR /go/src/DockerizedGoApp
RUN go build main.go
EXPOSE 8080
CMD ["./main"]