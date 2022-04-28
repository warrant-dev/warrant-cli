NAME    = warrant
GOENV   = GOARCH=amd64 GOOS=linux
GOCMD   = go
GOBUILD = $(GOCMD) build -o

.PHONY: clean
clean:
	rm -f $(NAME)

.PHONY: dev
dev: clean
	$(GOCMD) get
	$(GOBUILD) $(NAME) main.go

.PHONY: build
build: clean
	$(GOCMD) get
	$(GOENV) $(GOBUILD) $(NAME) main.go
