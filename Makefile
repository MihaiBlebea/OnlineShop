setup: cluster deploy-stage deploy-production

cluster: 
	namespaces
	roles

namespaces:
	kubectl create -f ./cluster/namespace-definition.yaml

roles:
	kubectl create -f ./cluster/roles

deploy-stage:
	# kubectl apply -f ./deploy/stage

remove-stage:
	kubectl delete namespace stage

deploy-production:
	kubectl apply -f ./deploy/production

remove-production:
	kubectl delete namespace production

remove: remove-stage remove-production

