.PHONY: all
all: run

.PHONY: run
run:
	export USER_HOST=:5053 SCOPE_HOST=:5063 DEBUG_ADDR=:5071 HTTP_ADDR=:5072 GRPC_ADDR=:5073 && go run auth-service/cmd/auth/main.go

