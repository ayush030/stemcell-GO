LINT=$(GOPATH)/bin/golangci-lint

build:
	docker build . -f Dockerfile

up:
	docker-compose up -d

db:
	docker-compose up db -d

run:
	go run cmd/main.go

lint:
	$(LINT) run


.PHONY: build up db run lint