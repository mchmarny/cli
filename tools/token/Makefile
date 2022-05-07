all: help

token: ## Prints new token 
	tools/token/make-token $(USER)
.PHONY: token

key: ## Prints new encryption key 
	tools/crypto/make-crypto-key
.PHONY: key

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help