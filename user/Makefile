.PHONY: all
all: run

.PHONY: run
run:
	export DB_HOST=localhost AUTH_SVC_HOST=:5073 DEBUG_ADDR=:5051 HTTP_ADDR=:5052 GRPC_ADDR=:5053 && \
	export AMQP_DSN=amqp://guest:guest@localhost:5672/ && \
	go run user-service/cmd/user/main.go

.PHONY: build
build:
	@docker build -t jasonheshuai/jwt-demo-user .
	@docker push jasonheshuai/jwt-demo-user
