.PHONY: all
all: run

.PHONY: run
run:
	docker-compose up -d --build user auth resource scope
	docker-compose up -d mysql

.PHONY: build
build:
	docker-compose build user auth resource scope
	docker push jasonheshuai/jwt-demo-user
	docker push jasonheshuai/jwt-demo-auth
	docker push jasonheshuai/jwt-demo-resource
	docker push jasonheshuai/jwt-demo-scope

.PHONY: mysql
mysql:
	kubectl port-forward $(shell kubectl get pods -n jwt-test -l app=mysql -o jsonpath="{.items[0].metadata.name}") -n jwt-test 33060:3306
