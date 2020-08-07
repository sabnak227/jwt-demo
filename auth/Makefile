.PHONY: all
all: run

.PHONY: run
run:
	export DB_HOST=localhost USER_HOST=:5053 SCOPE_HOST=:5063 DEBUG_ADDR=:5071 HTTP_ADDR=:5072 GRPC_ADDR=:5073 && go run auth-service/cmd/auth/main.go

.PHONY: build
build:
	docker build -t jasonheshuai/jwt-demo-auth .
	docker push jasonheshuai/jwt-demo-auth
