FROM golang:1.12 AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main scope-service/cmd/scope/main.go

FROM alpine:latest
COPY --from=builder /app/main /bin/app
ENTRYPOINT ["/bin/app"]