FROM golang:1.18

COPY . /src
WORKDIR /src

RUN CGO_ENABLE=0 GOOS=linux go build -mod=vendor -o whatsapp-api-pv .

FROM heroku/heroku:18
WORKDIR /app

COPY --from=0 /src/whatsapp-api-pv /app
CMD ["./whatsapp-api-pv"]