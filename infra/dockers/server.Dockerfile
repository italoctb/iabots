FROM golang

COPY /faq-server /src
WORKDIR /src
COPY ../../faq-server .
RUN CGO_ENABLED=0 GOOS=linux go build -o whatsapp-api-pv .

FROM alpine:latest AS production
WORKDIR /app

COPY --from=0 /src/whatsapp-api-pv /app
COPY --from=0 /src/server/ssr /app/server/ssr

EXPOSE 5000
CMD ["./whatsapp-api-pv"]