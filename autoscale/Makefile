.PHONY: all

project_id := ${PROJECT_ID}
version := 'autoscale' # ${GAE_VERSION}

help: ## Print this help
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

serve: ## ローカルサーバ実行
	goapp serve ./

deploy: ## デプロイ
	goapp deploy -application ${project_id} -version ${version} ./
