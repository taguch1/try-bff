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
name                 := bff-server
bin_dir              := bin
image_repo           := taguch1/$(name)
proto_dir            := ../../protos

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm

k8s_namespace        := default
k8s_deployment       := $(name)
k8s_context          := docker-for-desktop

# Include
#===============================================================
include .make/build.mk
include .make/docker.mk
include .make/helm.mk

# Task
#===============================================================
protoc:
	protoc -I $(proto_dir) --go_out=plugins=grpc:./proto $(proto_dir)/*.proto

run-local:
	docker run -it --rm -p1323:1323 $(image_repo):latest

deploy: docker-build helm-apply

forward:
	kubectl port-forward service/$(name) 1323:1323

.PHONY: protoc run-local deploy forward
.DEFAULT_GOAL := deploy
