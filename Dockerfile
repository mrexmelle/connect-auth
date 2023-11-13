FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
COPY config /etc/conf
RUN go build -o ./connect-idp ./cmd/main.go
RUN rm -rf ./cmd ./internal go

EXPOSE 8080
CMD ["/app/connect-idp", "serve"]

LABEL org.opencontainers.image.source https://github.com/OWNER/REPO