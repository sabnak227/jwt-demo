FROM golang:1.12 AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main scope-service/cmd/scope/main.go

FROM alpine:latest
COPY --from=builder /app/main /bin/app
ENTRYPOINT ["/bin/app"]