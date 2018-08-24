HAS_GLIDE := $(shell command -v glide;)
DIST := $(CURDIR)/_dist
BUILD := $(CURDIR)/_build
LDFLAGS := "-X main.version=${VERSION}"
BINARY := deploy

.PHONY: hookInstall
hookInstall: bootstrap test build

.PHONY: test
test:
	go test -v

.PHONY: build
build:
	go build -o $(BUILD)/$(BINARY) -ldflags $(LDFLAGS)

.PHONY: dist
dist:
	mkdir -p $(BUILD)
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD)/$(BINARY) -ldflags $(LDFLAGS) -a -tags netgo
	tar -C $(BUILD) -zcvf $(DIST)/tomcat-deployer-linux-$(VERSION).tgz $(BINARY)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD)/$(BINARY) -ldflags $(LDFLAGS) -a -tags netgo
	tar -C $(BUILD) -zcvf $(DIST)/tomcat-deployer-macos-$(VERSION).tgz $(BINARY)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD)/$(BINARY).exe -ldflags $(LDFLAGS) -a -tags netgo
	tar -C $(BUILD) -llzcvf $(DIST)/tomcat-deployer-windows-$(VERSION).tgz $(BINARY).exe

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
	glide install --strip-vendor

.PHONY: clean
clean:
	rm -rf _*