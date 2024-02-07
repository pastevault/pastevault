FROM golang:1.21 as backend

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

FROM oven/bun:1 as frontend

WORKDIR /app

COPY ui/package.json ./

RUN bun install

COPY ui .

RUN bun run build

FROM alpine:3.19

RUN apk add --no-cache nodejs

WORKDIR /app

COPY .docker/entrypoint.sh ./

COPY --from=backend /app/main ./

COPY --from=frontend /app/build /app/package.json ./ui/build/

EXPOSE 8080

EXPOSE 8081

CMD ["./entrypoint.sh"]