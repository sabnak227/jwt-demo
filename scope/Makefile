.PHONY: all
all: run

.PHONY: run
run:
	export DEBUG_ADDR=:5061 HTTP_ADDR=:5062 GRPC_ADDR=:5063 && go run scope-service/cmd/scope/main.go

