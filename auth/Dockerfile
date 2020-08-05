FROM golang:1.12 AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main auth-service/cmd/auth/main.go
# or FROM golang:alpine or some other base depending on need
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
