FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
COPY config /etc/conf
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency -g ./cmd/main.go
RUN go build -o ./connect-authx ./cmd/main.go
RUN rm -rf ./cmd ./internal go

EXPOSE 8080
CMD ["/app/connect-authx", "serve"]

LABEL org.opencontainers.image.source https://github.com/mrexmelle/connect-authx