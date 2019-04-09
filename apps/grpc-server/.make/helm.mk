# Task
#===============================================================
helm-cat: .change-cluster
	helm get $(helm_name)

helm-apply: .change-cluster
	$(eval helm_apply_options := --set image.tag=$(IMAGE_TAG))
	helm upgrade --install \
		--namespace $(k8s_namespace) \
		--values $(helm_dir)/env/$(ENV).yaml \
		$(helm_apply_options) \
		$(helm_name) $(helm_dir)

helm-delete: .change-cluster
	helm delete --purge $(helm_name)

helm-diff: .change-cluster
	$(eval helm_apply_options := --set image.tag=$(IMAGE_TAG))
	helm diff \
		--values $(helm_dir)/env/$(ENV).yaml \
		$(helm_apply_options) \
		$(helm_name) $(helm_dir)

helm-restart: .change-cluster
	$(eval hash := $(shell date +'%Y%m%d%H%M%S'))
	kubectl set env deployment/$(k8s_deployment) TMP_RECREATE_=$(hash)

# Internal Task
#===============================================================
.change-cluster:
	kubectx docker-for-desktop