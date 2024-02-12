FROM golang:1.21 as backend

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN go build -o pastevault ./cmd/server

FROM oven/bun:1 as frontend

WORKDIR /app

COPY ui/package.json ./

RUN bun install

COPY ui .

RUN bun run build

FROM debian:12-slim

RUN apt update

RUN apt install -y nodejs

WORKDIR /app

COPY .docker/entrypoint.sh ./

COPY --from=backend /app/pastevault ./

COPY --from=frontend /app/build /app/package.json ./ui/build/

EXPOSE 8080

EXPOSE 3000

ENV ORIGIN=http://localhost:3000

CMD ["./entrypoint.sh"]