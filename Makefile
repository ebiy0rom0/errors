GOCMD:=go
GOTEST:=$(GOCMD) test
GOTOOL:=$(GOCMD) tool

.PHONY: test
test:
	$(GOTEST) -v -p 12 -cover ./... -coverprofile='cover.txt'

.PHONY: cover
cover: test
	$(GOTOOL) cover -html 'cover.txt' -o 'cover.html'