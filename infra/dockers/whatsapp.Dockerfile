FROM golang

WORKDIR /src
COPY /whatsapp-client .

RUN CGO_ENABLED=0 GOOS=linux go build -o whatsapp-client .

FROM alpine:latest AS production
WORKDIR /app

COPY --from=0 /src/whatsapp-client /app

EXPOSE 5000
CMD ["./whatsapp-client"]