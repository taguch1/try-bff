# Task
#===============================================================
docker-build:
	docker build -t $(image_repo):$(IMAGE_TAG) .
	docker tag $(image_repo):$(tag) $(image_repo):latest

docker-push:
	docker push $(image_repo):$(IMAGE_TAG)
	docker push $(image_repo):latest

docker-clean:
	docker images --format "{{.Repository}}\t{{.CreatedAt}}\t{{.Tag}}" $(image_repo) | \
	  sort -r -k2,3 | \
	  awk -F"\t" '$$3 != "latest" && NR > 3 {print $$1":"$$3}' | \
	  xargs -n 1 docker rmi -f

.PHONY: docker-build docker-push docker-clean