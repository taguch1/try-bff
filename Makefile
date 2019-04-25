# Task
#===============================================================
all-deploy:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make'

all-delete:
	find ./apps/*/Makefile -exec dirname {} \; | xargs -I % -n 1 bash -c 'cd  "%" && make helm-delete'

all-info:
	kubectl get svc && kubectl get deploy && kubectl get po && kubectl get sts

.PHONY: all-deploy all-delete
.DEFAULT_GOAL := all-info

