.PHONY: init lint

init:
	@echo "Initializing the project..."
	@echo "Installing git hooks..."
	@rm -f .git/hooks/pre-commit
	@ln -s ../../githooks/pre-commit.sh .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit

lint:
	golangci-lint run ./...