FROM golang:1.18

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o whatsapp-api-pv .

FROM alpine:latest AS production
WORKDIR /app

COPY --from=0 /src/whatsapp-api-pv /app
COPY --from=0 /src/server/ssr /app/server/ssr

EXPOSE 5000
CMD ["./whatsapp-api-pv"]