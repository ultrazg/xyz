FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go clean -modcache && \
    go mod download -x && \
    go mod verify

COPY . .

RUN go build -o xyz .

CMD ["./xyz"]