# Task
#===============================================================
all-deploy:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make'

all-delete:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make helm-delete'

all-info:
	kubectl get svc && kubectl get deploy && kubectl get po && kubectl get sts

forward-api:
	cd apps/web-api/ && make forward

forward-js:
	cd apps/frontend-vue/ && make forward

forward:
	$(MAKE) -j 2 forward-api forward-js

.PHONY: all-deploy all-delete foward
.DEFAULT_GOAL := all-info

