.PHONY: all
all: run

.PHONY: run
run:
	export DEBUG_ADDR=:5081 HTTP_ADDR=:5082 GRPC_ADDR=:5083 && go run resource-service/cmd/resource/main.go

.PHONY: build
build:
	@docker build -t jasonheshuai/jwt-demo-resource .
	@docker push jasonheshuai/jwt-demo-resource