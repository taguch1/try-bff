# Option
#===============================================================
ENV           := local
OS            := $(shell uname | tr A-Z a-z )
SHELL         := /bin/bash
BUILD_OPTIONS := -tags netgo -installsuffix netgo
PREFIX        := /usr/local
INSTALL_BIN   :=
TAG_OPTION    := local-latest

# Const
#===============================================================
name                 := rdbms
bin_dir              := bin

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm

k8s_namespace        := default
k8s_deployment       := $(name)

# Include
#===============================================================
include .make/helm.mk

# Task
#===============================================================
migrate: .change-cluster
	kubectl exec -it $$(kubectl get po -l app=rdbms --no-headers  -o "custom-columns=NAME:.metadata.name") -- mysql -uroot -prdbms-password -e "$$(cat ./scripts/rdbms.sql)"

load-testdata: .change-cluster
	kubectl exec -it $$(kubectl get po -l app=rdbms --no-headers  -o "custom-columns=NAME:.metadata.name") -- mysql -uroot -prdbms-password -Drdbms -e "$$(cat scripts/testdata.sql)"

deploy: helm-apply migrate load-testdata

forward:
	kubectl port-forward service/rdbms 3306:3306

.PHONY: migrate load-testdata deploy forward
.DEFAULT_GOAL := deploy
