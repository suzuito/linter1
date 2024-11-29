BIN_SKELETON = go run github.com/gostaticanalysis/skeleton/v2@v2.2.2

NAME ?= linter_$(shell date +%s)

.PHONY: install-tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0

.PHONY: skeleton
skeleton:
	$(BIN_SKELETON) -gomod=false $(NAME)
