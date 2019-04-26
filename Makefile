# Const
#===============================================================
k8s_namespace                   := default
k8s_context                     := docker-for-desktop

k8s_namespace_monitoring        := monitoring

# Internal Task
#===============================================================
_forward-api-web-api:
	cd apps/web-api/ && $(MAKE) forward
_forward-frontend-vue:
	cd apps/frontend-vue/ && $(MAKE) forward
_forward-prometheus:
	cd monitoring/prometheus/ && $(MAKE) forward
_forward-grafana:
	cd monitoring/grafana/ && $(MAKE) forward
_open-js-frontend-vue:
	cd apps/frontend-vue/ && $(MAKE) open
_open-prometheus:
	cd monitoring/prometheus && $(MAKE) open
_open-grafana:
	cd monitoring/grafana/ && $(MAKE) open

# Task
#===============================================================
all-cat:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE) helm-cat'
	find ./monitoring/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE) helm-cat'
all-deploy:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE)'
	find ./monitoring/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE)'
all-delete:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE) helm-delete'
	find ./monitoring/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && $(MAKE) helm-delete'
all-info:
	@echo "\n[nodes]------------------------------------------------------------------------------------------------"
	@kubectl get nodes
	@echo "\n[$(k8s_namespace)]----------------------------------------------------------------------------------------------"
	@kubectl get all -n $(k8s_namespace)
	@echo "\n[$(k8s_namespace_monitoring)]-------------------------------------------------------------------------------------------"
	@kubectl get all -n $(k8s_namespace_monitoring)
all-forward:
	$(MAKE) -j 4 _forward-api-web-api _forward-frontend-vue _forward-prometheus _forward-grafana
all-open:
	$(MAKE) -j 3 _open-js-frontend-vue _open-prometheus _open-grafana

.PHONY: all-cat all-deploy all-delete all-info all-foward all-open
.DEFAULT_GOAL := all-info
