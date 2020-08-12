.PHONY: all
all: run

.PHONY: run
run:
	export USER_SVC_HOST=:5053 SCOPE_SVC_HOST=:5063 DEBUG_ADDR=:5071 HTTP_ADDR=:5072 GRPC_ADDR=:5073 && \
	export DB_HOST=localhost REDIS_HOST=localhost:6379 && \
	export AMQP_DSN=amqp://guest:guest@localhost:5672/ && \
	go run auth-service/cmd/auth/main.go

.PHONY: build
build:
	docker build -t jasonheshuai/jwt-demo-auth .
	docker push jasonheshuai/jwt-demo-auth