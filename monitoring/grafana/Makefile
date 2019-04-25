# Option
#===============================================================
ENV           := local
OS            := $(shell uname | tr A-Z a-z )
SHELL         := /bin/bash

# Const
#===============================================================
name                 := grafana
bin_dir              := bin

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm

k8s_namespace        := monitoring
k8s_deployment       := $(name)
k8s_context          := docker-for-desktop

local_port           := 8083

# Include
#===============================================================
include .make/*.mk

# Task
#===============================================================
deploy: helm-apply

forward:
	kubectl port-forward -n $(k8s_namespace) service/$(name) $(local_port):80

open:
	open http://localhost:$(local_port)

password:
	@kubectl get secret -n $(k8s_namespace) grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo


.PHONY: deploy forward
.DEFAULT_GOAL := deploy

