setup: cluster deploy-production


# Cluster
cluster: namespaces roles configs

namespaces:
	kubectl create -f ./cluster/namespace-definition.yaml

configs:
	kubectl create -f ./cluster/configs

roles:
	kubectl create -f ./cluster/roles

remove: remove-stage remove-production


# Stage namespace
deploy-stage:
	# kubectl apply -f ./deploy/stage

remove-stage:
	kubectl delete namespace stage


# Production namespace
deploy-production:
	kubectl apply -f ./deploy/production

remove-production:
	kubectl delete namespace production


