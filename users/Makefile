.PHONY: all
all: run

.PHONY: run
run:
	export DEBUG_ADDR=:5051 HTTP_ADDR=:5052 GRPC_ADDR=:5053 && go run user-service/cmd/user/main.go

