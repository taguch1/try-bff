# Option
#===============================================================
ENV           := local
OS            := $(shell uname | tr A-Z a-z )
SHELL         := /bin/bash
BUILD_OPTIONS := -tags netgo -installsuffix netgo
PREFIX        := /usr/local
INSTALL_BIN   :=
IMAGE_TAG     := $(shell git rev-parse --short HEAD)

# Const;
#===============================================================
name                 := grpc-server
bin_dir              := bin
srcs                 := $(shell find . -type f -name '*.go')
image_repo           := taguch1/$(name)
proto_dir            := ../../protos

helm_name            := $(name)
helm_apply_options   :=
helm_dir             := helm
k8s_namespace        := default
k8s_deployment       := $(name)

include .make/build.mk
include .make/docker.mk
include .make/helm.mk

# Task
#===============================================================
protoc:
	protoc -I $(proto_dir) --go_out=plugins=grpc:./proto $(proto_dir)/*.proto

run-local:
	docker run -it --rm -p50051:50051 $(image_repo):$(IMAGE_TAG)

.PHONY: protoc run-local