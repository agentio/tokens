all:
	go install ./...

image:
	docker buildx build --sbom=true --provenance=true --push -t docker.io/agentio/tokens:latest .
