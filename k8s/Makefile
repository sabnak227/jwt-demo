.PHONY: all
all: deploy


.PHONY: deploy
deploy:
	kubectl apply -n jwt-test -f deploy.yml
	kubectl apply -f gateway.yml

.PHONY: kiali
kiali:
	istioctl dashboard kiali

.PHONY: k8s
k8s:
	# kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin | awk '{print $1}')
	# login at http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
	kubectl proxy

.PHONY: init
init:
	kubectl apply -f namespace.yml
    kubectl label namespace jwt-test istio-injection=enabled