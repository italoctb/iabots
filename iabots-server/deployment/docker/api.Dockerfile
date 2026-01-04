FROM golang:1.25.5-alpine3.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM alpine:3.23

WORKDIR /app
COPY --from=build /bin/api /app/api

EXPOSE 5001
CMD ["/app/api"]
