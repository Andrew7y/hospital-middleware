FROM golang:1.25.12-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hospital-middleware ./cmd/main

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/hospital-middleware .

EXPOSE 8080

CMD ["./hospital-middleware"]