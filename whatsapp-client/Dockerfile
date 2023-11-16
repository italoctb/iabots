FROM golang

WORKDIR /src
COPY . .

ARG CLIENT_ID
ENV CLIENT_ID=$CLIENT_ID

RUN CGO_ENABLED=0 GOOS=linux go build -o whatsapp-client .

FROM alpine:latest AS production
WORKDIR /app

COPY --from=0 /src/whatsapp-client /app

EXPOSE 5000
CMD ["./whatsapp-client"]