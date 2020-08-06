FROM golang:1.12 AS builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main auth-service/cmd/auth/main.go

# Build the jwks.json file
FROM node:14.7 AS keygen
COPY . /app
WORKDIR /app/auth-service/keys
RUN ./keygen.sh

FROM alpine:latest
COPY --from=builder /app/main /bin/app
COPY --from=keygen /app/auth-service/keys/app.rsa /
COPY --from=keygen /app/auth-service/keys/app.rsa.pub /
COPY --from=keygen /app/auth-service/keys/jwks.json /
ENTRYPOINT ["/bin/app"]
