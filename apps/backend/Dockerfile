FROM golang:latest

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o app ./cmd/app

EXPOSE 8080
ENTRYPOINT ["./app"]
