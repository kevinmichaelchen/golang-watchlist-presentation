.PHONY: all
all:
	present
	open http://127.0.0.1:3999

.PHONY: install-present
install-present:
	go get -v -u golang.org/x/tools/cmd/present