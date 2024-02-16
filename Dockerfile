FROM golang:1.21 as backend

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o pastevault ./cmd/server

FROM debian:12-slim

WORKDIR /app

COPY --from=backend /app/pastevault .

ENV GIN_MODE=release

EXPOSE 8080
EXPOSE 8081

CMD ["./pastevault"]