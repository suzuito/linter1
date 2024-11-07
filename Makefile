BIN_SKELETON = go run github.com/gostaticanalysis/skeleton/v2@v2.2.2

NAME ?= linter_$(shell date +%s)

.PHONY: skeleton
skeleton:
	$(BIN_SKELETON) -gomod=false $(NAME)
