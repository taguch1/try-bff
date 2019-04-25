# Task
#===============================================================
helm-cat: .change-cluster
	helm get $(helm_name)

helm-apply: .change-cluster .image-tag
	helm upgrade --install \
		--namespace $(k8s_namespace) \
		--values $(helm_dir)/env/$(ENV).yaml \
		$(helm_apply_options) \
		$(helm_name) $(helm_dir)

helm-delete: .change-cluster
	helm delete --purge $(helm_name)

helm-rollback: .change-cluster
	helm rollback $(helm_name) 0

helm-diff: .change-cluster .image-tag
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
	kubectl config set-context $(k8s_context) --namespace=$(k8s_namespace)
