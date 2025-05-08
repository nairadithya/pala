FROM golang:1.23.8 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN GIN_MODE=release CGO_ENABLED=0 go build -o server -ldflags='-s -w' cmd/server/main.go

FROM golang:1.23.8

WORKDIR /app 

COPY --from=builder /app/server /app/server

COPY --from=builder /app/pkg/repository/migrations /app/pkg/repository/migrations

ENTRYPOINT ["/app/server"]
