# Etapa 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gocrud-auth-api main.go init_dependencies.go

# Etapa 2: Imagem final enxuta
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/gocrud-auth-api .
ENV GIN_MODE=release
EXPOSE 8080
CMD ["/app/gocrud-auth-api"] 