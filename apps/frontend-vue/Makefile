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
name                 := frontend-vue
bin_dir              := bin
image_repo           := taguch1/$(name)
proto_dir            := ../../protos

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm

k8s_namespace        := default
k8s_deployment       := $(name)
k8s_context          := docker-for-desktop

local_port           := 8080
# Include
#===============================================================
include .make/*.mk

# Task
#===============================================================
run-local:
	npm run serve

deploy: docker-build helm-apply

forward:
	kubectl port-forward -n $(k8s_namespace) service/$(name) $(local_port):80

open:
	open http://localhost:$(local_port)


.PHONY: run-local deploy forward
.DEFAULT_GOAL := deploy
