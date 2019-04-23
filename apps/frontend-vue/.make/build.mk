# Task
#===============================================================

setup:
	npm install
lint:
	npm run lint
test:
	npm test

build:
	npm run build

release: lint test build


.PHONY: setup lint test build release
