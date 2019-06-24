FROM golang:alpine

WORKDIR /go/src/app
COPY . .
RUN apk add git
RUN go get -d -v ./...
RUN go build -o hello main.go

CMD ["./hello"]
