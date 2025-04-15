# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum para aproveitar o cache
COPY go.mod go.sum ./
RUN go mod download

# Instalar o Wire se necessário
RUN go install github.com/google/wire/cmd/wire@latest

# Copiar o código fonte
COPY . .

# Navegar para o diretório correto com os arquivos main.go e wire_gen.go
WORKDIR /app/cmd/ordersystem

# Compilar a aplicação incluindo o main.go e wire_gen.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ordersystem main.go wire_gen.go

# Stage final
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/cmd/ordersystem/ordersystem .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations/ ./migrations
# Expor a porta que a aplicação usa
EXPOSE 8000
EXPOSE 8080
EXPOSE 3000

# Executar a aplicação
CMD ["./ordersystem"]
# CMD ["tail","-f","/dev/null"]
