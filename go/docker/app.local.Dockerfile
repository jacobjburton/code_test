FROM golang:alpine as build

RUN apk add --no-cache git

RUN go install github.com/githubnemo/CompileDaemon@latest \
    && chmod +x /go/bin/CompileDaemon

WORKDIR /go/src/app

ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main