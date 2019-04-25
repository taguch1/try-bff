# Option
#===============================================================
ENV           := local
OS            := $(shell uname | tr A-Z a-z )
SHELL         := /bin/bash
PREFIX        := /usr/local
INSTALL_BIN   :=
TAG_OPTION    := local-latest

# Const
#===============================================================
name                 := web-api
bin_dir              := bin
image_repo           := taguch1/$(name)

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm

k8s_namespace        := default
k8s_deployment       := $(name)
k8s_context          := docker-for-desktop

# Include
#===============================================================
include .make/docker.mk
include .make/helm.mk

deploy: docker-build helm-apply

forward:
	kubectl port-forward service/$(name) 8081:80

.PHONY: deploy forward
.DEFAULT_GOAL := deploy
