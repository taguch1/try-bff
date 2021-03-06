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
k8s_context          := docker-for-desktop

local_port           := 3306

# Include
#===============================================================
include .make/helm.mk

# Task
#===============================================================
migrate: .change-cluster
	kubectl exec -it $$(kubectl get po -l app=rdbms --no-headers  -o "custom-columns=NAME:.metadata.name") -- mysql -uroot -prdbms-password -e "$$(cat ./scripts/rdbms.sql)"

load-testdata: .change-cluster
	kubectl exec -it $$(kubectl get po -l app=rdbms --no-headers  -o "custom-columns=NAME:.metadata.name") -- mysql -uroot -prdbms-password -Drdbms -e "$$(cat scripts/testdata.sql)"

setup:migrate load-testdata

deploy:
	$(MAKE) helm-apply
	sleep 30 && $(MAKE) setup &

forward:
	kubectl port-forward -n $(k8s_namespace) service/$(name) $(local_port):3306

.PHONY: migrate load-testdata deploy forward
.DEFAULT_GOAL := deploy
