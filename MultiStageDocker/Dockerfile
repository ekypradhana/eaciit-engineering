#STAGE 1
FROM golang:latest as builder
MAINTAINER Eky Pradhana eky.pradhana@eaciit.com
ADD . /go/src/MultiStageDocker
WORKDIR /go/src/MultiStageDocker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#STAGE 2
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/MultiStageDocker/main .
CMD ["./main"]