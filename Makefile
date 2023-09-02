.PHONY: database

database:
	kubectl apply -f deploy/database/postgres-config.yaml
	kubectl apply -f deploy/database/postgres-volume-claim.yaml
	kubectl apply -f deploy/database/postgres-deployment.yaml