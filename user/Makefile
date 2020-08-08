.PHONY: all
all: run

.PHONY: run
run:
	export DB_HOST=localhost DEBUG_ADDR=:5051 HTTP_ADDR=:5052 GRPC_ADDR=:5053 && go run user-service/cmd/user/main.go

.PHONY: build
build:
	@docker build -t jasonheshuai/jwt-demo-user .
	@docker push jasonheshuai/jwt-demo-user
