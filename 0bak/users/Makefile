
.PHONY: run
run: 
	export DB_HOST=localhost DB_PORT=3306 && go run user-service/cmd/user/main.go

.PHONY: regen
regen: 
	truss user.proto

.PHONY: lint
lint: 
	prototool lint .