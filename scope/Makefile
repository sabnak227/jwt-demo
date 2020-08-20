.PHONY: all
all: run

.PHONY: run
run:
	export DEBUG_ADDR=:5061 HTTP_ADDR=:5062 GRPC_ADDR=:5063 && \
	export DB_HOST=localhost SEED=true && \
	export AMQP_DSN=amqp://guest:guest@localhost:5672/ && \
	go run scope-service/cmd/scope/main.go

.PHONY: build
build:
	@docker build -t jasonheshuai/jwt-demo-scope .
	@docker push jasonheshuai/jwt-demo-scope