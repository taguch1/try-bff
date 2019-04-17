

# Task
#===============================================================
docker-build: .build-image-tag
	docker build -t $(image_repo):$(tag) .
	docker tag $(image_repo):$(tag) $(image_repo):latest

docker-push: .image-tag
	docker push $(image_repo):$(tag)
	docker push $(image_repo):latest

docker-clean:
	docker images --format "{{.Repository}}\t{{.CreatedAt}}\t{{.Tag}}" $(image_repo) | \
	  sort -r -k2,3 | \
	  awk -F"\t" '$$3 != "latest" && NR > 3 {print $$1":"$$3}' | \
	  xargs -n 1 docker rmi -f

# Internal Task
#===============================================================
.build-image-tag:
ifeq ($(BUILD_TAG),)
	$(eval tag := $(shell date +'%Y-%m-%dT%H%M%S'))
else
	$(eval tag := $(BUILD_TAG))
endif

.image-tag:
ifeq ($(TAG_OPTION),)
	$(error TAG_OPTION not set correctly.)
endif

ifeq ($(image_repo),)
	$(error image_repo not set correctly.)
endif

	$(eval tag := $(TAG_OPTION))
	$(eval helm_apply_options := --set image.tag=$(TAG_OPTION))

ifeq ($(TAG_OPTION),local-latest)
	$(eval tag := $(shell docker images --format "{{.ID}}\t{{.CreatedAt}}\t{{.Tag}}" $(image_repo) | \
	  sort -r -k2,3 | \
	  awk -F"\t" '$$3 != "latest" {print $$3}' | \
	  head -n 1 ))
	$(eval tag := $(tag))
	$(eval helm_apply_options := --set image.tag=$(tag))
endif

.PHONY: docker-build docker-push docker-clean