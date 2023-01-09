FROM golang:latest as builder

COPY . /go/src/app
WORKDIR /go/src/app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM scratch

COPY --from=builder /go/src/app/main /main

CMD ["/main"]
