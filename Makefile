
# Const
#===============================================================
k8s_namespace        := default
k8s_context          := docker-for-desktop

# Task
#===============================================================
all-deploy:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make'

all-delete:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make helm-delete'

all-info:
	kubectl get svc && kubectl get deploy && kubectl get po && kubectl get sts

forward-api:
	cd apps/web-api/ && $(MAKE) forward

forward-js:
	cd apps/frontend-vue/ && $(MAKE) forward

forward-prometheus:
	cd monitoring/prometheus/ && $(MAKE) forward

forward-grafana:
	cd monitoring/grafana/ && $(MAKE) forward

forward:
	$(MAKE) -j 4 forward-api forward-js forward-grafana forward-prometheus

open-js:
	cd apps/frontend-vue/ && $(MAKE) open

open-prometheus:
	cd monitoring/prometheus && $(MAKE) open

open-grafana:
	cd monitoring/grafana/ && $(MAKE) open

open:
	$(MAKE) -j 3 open-js open-prometheus open-grafana


.PHONY: all-deploy all-delete foward
.DEFAULT_GOAL := all-info



