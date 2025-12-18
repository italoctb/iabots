# Usa imagem oficial mínima do Go
FROM golang:1.21-alpine

# Define diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum primeiro para cache de dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN go build -o main ./cmd/api

# Expondo porta da API
EXPOSE 8080

# Comando que será executado ao iniciar o container
CMD [\"/app/main\"]
