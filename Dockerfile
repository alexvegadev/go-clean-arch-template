FROM golang:1.18.1-alpine

WORKDIR /app-go

COPY . .

RUN GOFLAGS=-buildvcs=false CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd/api/

EXPOSE 8080

ENV GIN_MODE=release

CMD ["./app"]