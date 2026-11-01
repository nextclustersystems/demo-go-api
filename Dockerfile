# Build: baixa as dependências pelo proxy de módulos Go e compila um binário estático.
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api .

# Imagem final pequena, sem compilador, rodando como usuário sem privilégios.
FROM alpine:3.20
RUN adduser -D -u 10001 app
COPY --from=build /out/api /api
USER app
EXPOSE 8080
# PORT vem da plataforma (Configurações); 8080 quando rodar local.
ENTRYPOINT ["/api"]
